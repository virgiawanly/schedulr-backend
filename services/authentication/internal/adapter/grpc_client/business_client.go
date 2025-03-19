package grpc_client

import (
	"context"
	"fmt"
	"schedulr-backend/services/authentication/config"
	"schedulr-backend/services/authentication/internal/domain/dto"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"

	clientPort "schedulr-backend/services/authentication/internal/port/grpc_client"

	businessV1 "github.com/virgiawanly/schedulr-backend/protogen/go/business/v1"
)

type businessClient struct {
	client businessV1.BusinessServiceClient
	cb     *gobreaker.CircuitBreaker
}

func NewBusinessClient(conn grpc.ClientConnInterface, conf *config.CircuitBreaker) clientPort.BusinessClientPort {
	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "schedulr-token",
		MaxRequests: conf.MaxRequests,
		Interval:    conf.ResetInterval,
		Timeout:     conf.OpenStateTimeout,
	})

	return &businessClient{
		client: businessV1.NewBusinessServiceClient(conn),
		cb:     breaker,
	}
}

func (u *businessClient) CreateBusinessFromRegistration(ctx context.Context, reqDto *dto.CreateBusinessFromRegistrationRequestDTO) (*dto.CreateBusinessFromRegistrationResponseDTO, error) {
	req := &businessV1.CreateBusinessFromRegistrationRequest{
		Name:        reqDto.Name,
		Industry:    reqDto.Industry,
		CompanySize: reqDto.CompanySize,
		Address_1:   reqDto.Address1,
		Address_2:   reqDto.Address2,
		City:        reqDto.City,
		State:       reqDto.State,
		Zipcode:     reqDto.Zipcode,
		Country:     reqDto.Country,
	}

	result, err := u.cb.Execute(func() (interface{}, error) {
		return u.client.CreateBusinessFromRegistration(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*businessV1.CreateBusinessFromRegistrationResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}

	return &dto.CreateBusinessFromRegistrationResponseDTO{Business: dto.BusinessResponse{
		ID:       resp.Business.Id,
		Name:     resp.Business.Name,
		Industry: resp.Business.Industry,
	}}, nil
}
