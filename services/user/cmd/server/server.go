package main

import (
	"fmt"
	"log"
	"net"
	"schedulr-backend/services/user/config"
	"schedulr-backend/services/user/internal/adapter/grpc/handler"
	"schedulr-backend/services/user/internal/adapter/persistence/pg"
	"schedulr-backend/services/user/internal/domain/usecase"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/virgiawanly/schedulr-backend/shared/grpc/interceptor"
	"github.com/virgiawanly/schedulr-backend/shared/logger"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	userV1 "github.com/virgiawanly/schedulr-backend/protogen/go/user/v1"
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

	userRepository := pg.NewUserRepository(s.db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	userV1.RegisterUserServiceServer(s.grpcServer, userHandler)

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
