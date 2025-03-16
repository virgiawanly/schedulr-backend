package main

import (
	"fmt"
	"log"
	"net"
	"schedulr-backend/services/authentication/config"
	"schedulr-backend/services/authentication/internal/adapter/grpc/handler"
	"schedulr-backend/services/authentication/internal/adapter/grpc_client"
	"schedulr-backend/services/authentication/internal/adapter/persistence/pg"
	"schedulr-backend/services/authentication/internal/domain/usecase"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/virgiawanly/schedulr-backend/shared/grpc/interceptor"
	"github.com/virgiawanly/schedulr-backend/shared/jwt"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	authenticationV1 "github.com/virgiawanly/schedulr-backend/protogen/go/authentication/v1"
)

type Server struct {
	config     *config.Config
	db         *gorm.DB
	logger     logger.Logger
	grpcServer *grpc.Server
}

func NewServer(config *config.Config) (*Server, error) {
	srv := &Server{
		config: config,
	}

	err := srv.setup()
	return srv, err
}

func (s *Server) setup() error {
	if err := s.setupDatabase(); err != nil {
		return err
	}

	if err := s.setupGrpcServer(); err != nil {
		return err
	}

	return nil
}

func (s *Server) setupDatabase() error {
	logger.Info("Connecting to the database...")

	dsn := s.config.DB.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	s.db = db
	logger.Info("Connected to the database")

	return nil
}

func (s *Server) setupGrpcServer() error {
	s.logger = logger.GetGlobal()

	globalZapLogger, ok := s.logger.(*logger.ZapLogger)
	if !ok {
		return fmt.Errorf("logger is not a ZapLogger")
	}

	zapLogger := interceptor.ZapLogger(globalZapLogger.Logger)

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(zapLogger),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(zapLogger),
		),
	)

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	grpcClient, err := grpc_client.NewClientWithConfig(&s.config.Server.Service, &s.config.Server.CircuitBreaker, creds)
	if err != nil {
		return fmt.Errorf("failed to create grpc client: %w", err)
	}

	jwtConfig := jwt.NewJWTConfig(
		s.config.JWTConfig.SecretKey,
		s.config.JWTConfig.AccessTokenTTL,
		s.config.JWTConfig.RefreshTokenTTL,
	)

	accountRepository := pg.NewAccountRepository(s.db)
	authUsecase := usecase.NewAuthUsecase(accountRepository, grpcClient.UserClient, jwtConfig)
	authHandler := handler.NewAuthHandler(authUsecase)
	authenticationV1.RegisterAuthServiceServer(s.grpcServer, authHandler)

	return nil
}

func (s *Server) Run() {
	logger.Info("Starting gRPC server...")

	address := s.config.Server.GetAddress()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %v: %v", address, err)
	}

	if err := s.grpcServer.Serve(listener); err != nil {
		logger.Fatalf("Failed to serve gRPC: %v", err)
	}

	logger.Infof("gRPC Server is running on %v", address)
}
