package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"

	"schedulr-backend/services/authentication/internal/domain/dto"
	"schedulr-backend/services/authentication/internal/domain/entity"
	"schedulr-backend/services/authentication/internal/domain/mapper"
	"schedulr-backend/services/authentication/internal/port/repository"

	"github.com/virgiawanly/schedulr-backend/shared/jwt"

	clientPort "schedulr-backend/services/authentication/internal/port/grpc_client"

	sharedError "github.com/virgiawanly/schedulr-backend/shared/errors"
)

type AuthUsecase struct {
	accountRepository repository.AccountRepository
	userClient        clientPort.UserClientPort
	jwtConfig         *jwt.JWTConfig
}

func NewAuthUsecase(accountRepo repository.AccountRepository, userClient clientPort.UserClientPort, jwtConfig *jwt.JWTConfig) *AuthUsecase {
	return &AuthUsecase{
		accountRepository: accountRepo,
		userClient:        userClient,
		jwtConfig:         jwtConfig,
	}
}

func (u *AuthUsecase) CheckRegisteredEmail(ctx context.Context, req dto.CheckRegisteredEmailRequestDTO) (bool, *sharedError.AppError) {
	result, err := u.accountRepository.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return false, sharedError.WrapError(codes.Internal, err)
	}
	return result, nil
}

func (u *AuthUsecase) Register(ctx context.Context, req dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, *sharedError.AppError) {
	emailExists, err := u.accountRepository.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}
	if emailExists {
		return nil, sharedError.WrapError(codes.AlreadyExists, sharedError.ErrEmailAlreadyRegistered)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	now := time.Now()
	businessId := uuid.New().String() // @todo: get business id from newly created account

	account := entity.Account{
		ID:         uuid.New().String(),
		BusinessID: businessId,
		Email:      req.Email,
		Password:   string(hashedPassword),
		CreatedAt:  &now,
		UpdatedAt:  &now,
	}

	tx := u.accountRepository.BeginTransaction()
	defer tx.Rollback()

	err = u.accountRepository.Create(ctx, &account, tx)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	_, err = u.userClient.CreateUserFromRegistration(ctx, &account.BusinessID, &dto.CreateUserFromRegistrationRequestDTO{
		AccountID: account.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address1:  req.Address1,
		Address2:  req.Address2,
		City:      req.City,
		State:     req.State,
		Zipcode:   req.Zipcode,
		Country:   req.Country,
	})
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	response := mapper.ToRegisterResponseDTO(account)
	return &response, nil
}

func (u *AuthUsecase) Login(ctx context.Context, req dto.LoginRequestDTO) (*dto.LoginResponseDTO, *sharedError.AppError) {
	account, err := u.accountRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}
	if account == nil {
		return nil, sharedError.NewAppError(codes.Unauthenticated, "account not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return nil, sharedError.NewAppError(codes.Unauthenticated, "invalid credentials")
	}

	accessToken, err := u.jwtConfig.GenerateToken(account.ID, account.Role, account.Email, false)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	refreshToken, err := u.jwtConfig.GenerateToken(account.ID, account.Role, account.Email, true)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	return &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiryTime:   time.Duration(u.jwtConfig.AccessTokenTTL.Seconds()),
		RefreshToken: refreshToken,
	}, nil
}
