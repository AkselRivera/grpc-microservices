package handlers

import (
	"context"

	"github.com/akselrivera/grpc-auth-microservice/internal/service"
	pb "github.com/akselrivera/grpc-auth-microservice/proto"
)

type AuthController struct {
	pb.UnimplementedAuthServiceServer
}

var authService service.AuthService

// Login implements proto.AuthServiceServer.
func (s AuthController) Login(c context.Context, loginRequest *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginResponse, err := authService.Login(c, loginRequest)

	if err != nil {
		return &pb.LoginResponse{}, err
	}

	return loginResponse, nil
}

// Register implements proto.AuthServiceServer.
func (s AuthController) Register(c context.Context, registerRequest *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	registerResponse, err := authService.Register(c, registerRequest)

	if err != nil {
		return &pb.RegisterResponse{}, err
	}

	return registerResponse, nil
}

// ValidateToken implements proto.AuthServiceServer.
func (s AuthController) ValidateToken(c context.Context, validateTokenRequest *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	validateTokenResponse, err := authService.ValidateToken(c, validateTokenRequest)

	if err != nil {
		return &pb.ValidateTokenResponse{}, err
	}

	return validateTokenResponse, nil
}
