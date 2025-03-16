package grpc_client

import (
	"context"
	"fmt"
	"schedulr-backend/services/authentication/config"
	"schedulr-backend/services/authentication/internal/domain/dto"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	clientPort "schedulr-backend/services/authentication/internal/port/grpc_client"

	userV1 "github.com/virgiawanly/schedulr-backend/protogen/go/user/v1"
)

type userClient struct {
	client userV1.UserServiceClient
	cb     *gobreaker.CircuitBreaker
}

func NewUserClient(conn grpc.ClientConnInterface, conf *config.CircuitBreaker) clientPort.UserClientPort {
	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "schedulr-token",
		MaxRequests: conf.MaxRequests,
		Interval:    conf.ResetInterval,
		Timeout:     conf.OpenStateTimeout,
	})

	return &userClient{
		client: userV1.NewUserServiceClient(conn),
		cb:     breaker,
	}
}

func (u *userClient) FindUserById(ctx context.Context, businessId *string, reqDto *dto.FindUserByIdRequestDTO) (*dto.FindUserByIdResponseDTO, error) {
	if businessId != nil {
		ctx = metadata.AppendToOutgoingContext(ctx, "business_id", *businessId)
	}

	req := &userV1.FindUserByIdRequest{Id: reqDto.ID}

	result, err := u.cb.Execute(func() (interface{}, error) {
		return u.client.FindUserById(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*userV1.FindUserByIdResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	return &dto.FindUserByIdResponseDTO{User: resp.User}, nil
}

func (u *userClient) CreateUserFromRegistration(ctx context.Context, businessId *string, reqDto *dto.CreateUserFromRegistrationRequestDTO) (*dto.CreateUserFromRegistrationResponseDTO, error) {
	if businessId != nil {
		ctx = metadata.AppendToOutgoingContext(ctx, "business_id", *businessId)
	}

	req := &userV1.CreateUserFromRegistrationRequest{
		AccountId: reqDto.AccountID,
		FirstName: reqDto.FirstName,
		LastName:  reqDto.LastName,
		Phone:     reqDto.Phone,
		Address_1: reqDto.Address1,
		Address_2: reqDto.Address2,
		City:      reqDto.City,
		State:     reqDto.State,
		Zipcode:   reqDto.Zipcode,
		Country:   reqDto.Country,
	}

	result, err := u.cb.Execute(func() (interface{}, error) {
		return u.client.CreateUserFromRegistration(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*userV1.CreateUserFromRegistrationResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	return &dto.CreateUserFromRegistrationResponseDTO{User: resp.User}, nil
}

func (u *userClient) UpdateUserByAccountId(ctx context.Context, businessId *string, reqDto *dto.UpdateUserByAccountIdRequestDTO) (*dto.UpdateUserByAccountIdResponseDTO, error) {
	if businessId != nil {
		ctx = metadata.AppendToOutgoingContext(ctx, "business_id", *businessId)
	}

	req := &userV1.UpdateUserByAccountIdRequest{
		AccountId: reqDto.AccountID,
		FirstName: reqDto.FirstName,
		LastName:  reqDto.LastName,
		Phone:     reqDto.Phone,
		Address_1: reqDto.Address1,
		Address_2: reqDto.Address2,
		City:      reqDto.City,
		State:     reqDto.State,
		Zipcode:   reqDto.Zipcode,
		Country:   reqDto.Country,
	}

	result, err := u.cb.Execute(func() (interface{}, error) {
		return u.client.UpdateUserByAccountId(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*userV1.UpdateUserByAccountIdResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	return &dto.UpdateUserByAccountIdResponseDTO{User: resp.User}, nil
}
