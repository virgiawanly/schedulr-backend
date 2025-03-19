package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/virgiawanly/schedulr-backend/gateway/config"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authenticationV1Gw "github.com/virgiawanly/schedulr-backend/protogen/gateway/go/authentication/v1"
	businessV1Gw "github.com/virgiawanly/schedulr-backend/protogen/gateway/go/business/v1"
	userV1Gw "github.com/virgiawanly/schedulr-backend/protogen/gateway/go/user/v1"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.config.Server.Ip, s.config.Server.Port)
}

func NewServer(config *config.Config) (*Server, error) {
	srv := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	if err := srv.registerServices(); err != nil {
		return nil, err
	}

	return srv, nil
}

func (s *Server) registerServices() error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx := context.Background()

	if err := authenticationV1Gw.RegisterAuthServiceHandlerFromEndpoint(ctx, s.mux, s.config.Server.Service.Authentication, opts); err != nil {
		return fmt.Errorf("failed to register Authentication Service: %v", err)
	}

	if err := businessV1Gw.RegisterBusinessServiceHandlerFromEndpoint(ctx, s.mux, s.config.Server.Service.Business, opts); err != nil {
		return fmt.Errorf("failed to register Business Service: %v", err)
	}

	if err := userV1Gw.RegisterUserServiceHandlerFromEndpoint(ctx, s.mux, s.config.Server.Service.User, opts); err != nil {
		return fmt.Errorf("failed to register User Service: %v", err)
	}

	return nil
}

func (s *Server) Run() error {
	logger.Infof("Starting API Gateway on %s", s.Address())

	server := &http.Server{
		Addr:    s.Address(),
		Handler: s.mux,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start API Gateway: %v", err)
	}

	return nil
}
