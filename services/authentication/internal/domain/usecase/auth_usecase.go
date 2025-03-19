package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

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
	businessClient    clientPort.BusinessClientPort
	jwtConfig         *jwt.JWTConfig
}

func NewAuthUsecase(accountRepo repository.AccountRepository, userClient clientPort.UserClientPort, businessClient clientPort.BusinessClientPort, jwtConfig *jwt.JWTConfig) *AuthUsecase {
	return &AuthUsecase{
		accountRepository: accountRepo,
		userClient:        userClient,
		businessClient:    businessClient,
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

	tx := u.accountRepository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	createBusinessRequest := mapper.ToCreateBusinessFromRegistrationRequestDTO(req)
	businessResponse, err := u.businessClient.CreateBusinessFromRegistration(ctx, &createBusinessRequest)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	now := time.Now()
	account := entity.Account{
		ID:         uuid.New().String(),
		BusinessID: businessResponse.Business.ID,
		Email:      req.Email,
		Password:   string(hashedPassword),
		CreatedAt:  &now,
		UpdatedAt:  &now,
	}

	err = u.accountRepository.Create(ctx, &account, tx)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	// Create User in User Service
	createUserRequestDto := mapper.ToCreateUserFromRegistrationRequestDTO(req, account.ID)
	_, err = u.userClient.CreateUserFromRegistration(ctx, &account.BusinessID, &createUserRequestDto)
	if err != nil {
		// _ = u.businessClient.DeleteBusiness(ctx, businessResponse.Business.ID)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharedError.NewAppError(codes.Unauthenticated, "account not found")
		}
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
