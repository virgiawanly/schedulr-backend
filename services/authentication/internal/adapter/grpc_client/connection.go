package grpc_client

import (
	"schedulr-backend/services/authentication/config"

	"github.com/virgiawanly/schedulr-backend/shared/logger"
	"google.golang.org/grpc"

	clientPort "schedulr-backend/services/authentication/internal/port/grpc_client"
)

type GrpcClient struct {
	UserClient     clientPort.UserClientPort
	BusinessClient clientPort.BusinessClientPort
	connections    []*grpc.ClientConn
}

func (e *GrpcClient) Close() {
	for _, conn := range e.connections {
		if err := conn.Close(); err != nil {
			logger.Warnf("Error on closing external: %s", err.Error())
		}
	}
}

func NewClientWithConnection(userConn *grpc.ClientConn, businessConn *grpc.ClientConn, conf *config.CircuitBreaker) *GrpcClient {
	return &GrpcClient{
		UserClient:     NewUserClient(userConn, conf),
		BusinessClient: NewBusinessClient(businessConn, conf),
		connections:    []*grpc.ClientConn{userConn},
	}
}

func NewClientWithConfig(conf *config.Service, breaker *config.CircuitBreaker, options ...grpc.DialOption) (*GrpcClient, error) {
	userConn, err := grpc.NewClient(conf.User, options...)
	if err != nil {
		return nil, err
	}

	businessConn, err := grpc.NewClient(conf.Business, options...)
	if err != nil {
		return nil, err
	}

	return NewClientWithConnection(userConn, businessConn, breaker), nil
}
