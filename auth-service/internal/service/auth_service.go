package service

import (
	"context"

	"github.com/akselrivera/grpc-auth-microservice/internal/utils"
	pb "github.com/akselrivera/grpc-auth-microservice/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

type User struct {
	Id       string
	Email    string
	Password string
}

var users []User

func (s *AuthService) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	for _, user := range users {
		if user.Email == req.Email {

			if ok := utils.CheckPasswordHash(req.Password, user.Password); ok {
				token := utils.NewToken(user.Id)

				return &pb.LoginResponse{Token: token}, nil
			}
			return &pb.LoginResponse{}, status.Errorf(codes.PermissionDenied, "invalid credentials")
		}
	}

	return &pb.LoginResponse{}, status.Errorf(codes.PermissionDenied, "invalid credentials")
}

func (s *AuthService) Register(c context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	for _, user := range users {
		if user.Email == req.Email {
			return &pb.RegisterResponse{}, status.Errorf(codes.AlreadyExists, "user already exists")
		}
	}
	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return &pb.RegisterResponse{}, status.Errorf(codes.Internal, "error hashing password")
	}

	id := uuid.New().String()

	users = append(users, User{Id: id, Email: req.Email, Password: hashedPassword})
	token := utils.NewToken(id)

	return &pb.RegisterResponse{Token: token}, nil
}

func (s *AuthService) ValidateToken(c context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	if err := utils.VerifyToken(req.Token); err != nil {
		return &pb.ValidateTokenResponse{}, status.Errorf(codes.PermissionDenied, "invalid token")
	}
	return &pb.ValidateTokenResponse{Valid: true}, nil
}
