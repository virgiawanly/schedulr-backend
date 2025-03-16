package handler

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/dto"
	"schedulr-backend/services/authentication/internal/domain/mapper"
	"schedulr-backend/services/authentication/internal/domain/usecase"

	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/durationpb"

	authenticationV1 "github.com/virgiawanly/schedulr-backend/protogen/go/authentication/v1"
	errorutil "github.com/virgiawanly/schedulr-backend/shared/util/errors"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
	validator   *validator.Validate
	authenticationV1.UnimplementedAuthServiceServer
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: usecase,
		validator:   validator.New(),
	}
}

func (h *AuthHandler) CheckRegisteredEmail(ctx context.Context, req *authenticationV1.CheckRegisteredEmailRequest) (*authenticationV1.CheckRegisteredEmailResponse, error) {
	checkRegisteredEmailDto := dto.CheckRegisteredEmailRequestDTO{Email: req.Email}

	if err := h.validator.Struct(checkRegisteredEmailDto); err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	exists, err := h.authUsecase.CheckRegisteredEmail(ctx, checkRegisteredEmailDto)
	if err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	return &authenticationV1.CheckRegisteredEmailResponse{IsExist: exists}, nil
}

func (h *AuthHandler) Register(ctx context.Context, req *authenticationV1.RegisterRequest) (*authenticationV1.RegisterResponse, error) {
	registerDto := mapper.ToRegisterRequestDTO(req)

	if err := h.validator.Struct(registerDto); err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	result, err := h.authUsecase.Register(ctx, registerDto)

	if err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	return &authenticationV1.RegisterResponse{AccountId: result.AccountID}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *authenticationV1.LoginRequest) (*authenticationV1.LoginResponse, error) {
	loginDto := dto.LoginRequestDTO{Email: req.Email, Password: req.Password}

	if err := h.validator.Struct(loginDto); err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	res, err := h.authUsecase.Login(ctx, loginDto)
	if err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	return &authenticationV1.LoginResponse{
		AccessToken:  res.AccessToken,
		TokenType:    res.TokenType,
		ExpiryTime:   durationpb.New(res.ExpiryTime),
		RefreshToken: res.RefreshToken,
	}, nil
}
