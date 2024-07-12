package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/akselrivera/grpc-auth-microservice/internal/handlers"
	pb "github.com/akselrivera/grpc-auth-microservice/proto"
)

func main() {
	s := grpc.NewServer()

	var authController handlers.AuthController
	pb.RegisterAuthServiceServer(s, authController)

	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	log.Println("Servidor gRPC escuchando en :50056")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
