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
		log.Fatalf("Connection error: %v", err)
	}

	log.Println("gRPC server listening on :50055")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Connection error: %v", err)
	}
}
