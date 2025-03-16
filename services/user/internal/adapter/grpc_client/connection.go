package grpc_client

import (
	"schedulr-backend/services/user/config"

	"github.com/virgiawanly/schedulr-backend/shared/logger"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	connections []*grpc.ClientConn
}

func (e *GrpcClient) Close() {
	for _, conn := range e.connections {
		if err := conn.Close(); err != nil {
			logger.Warnf("Error on closing external: %s", err.Error())
		}
	}
}

func NewClientWithConnection(userConn *grpc.ClientConn, conf *config.CircuitBreaker) *GrpcClient {
	return &GrpcClient{
		connections: []*grpc.ClientConn{userConn},
	}
}

func NewClientWithConfig(conf *config.Service, breaker *config.CircuitBreaker, options ...grpc.DialOption) (*GrpcClient, error) {
	userConn, err := grpc.NewClient(conf.User, options...)
	if err != nil {
		return nil, err
	}

	return NewClientWithConnection(userConn, breaker), nil
}
