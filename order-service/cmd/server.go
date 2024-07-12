package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/akselrivera/grpc-order-microservice/internal/handlers"
	pb "github.com/akselrivera/grpc-order-microservice/proto"
)

func main() {
	s := grpc.NewServer()

	var orderController handlers.OrderController
	pb.RegisterOrderServiceServer(s, orderController)

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	log.Println("Servidor gRPC escuchando en :50055")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
