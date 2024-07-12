package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/akselrivera/grpc-apigateway/internal/config"
	"github.com/akselrivera/grpc-apigateway/internal/router"
	pb "github.com/akselrivera/grpc-apigateway/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"google.golang.org/grpc"
)

func main() {

	// Configurar servidor gRPC
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &pb.UnimplementedAuthServiceServer{})
	pb.RegisterProductServiceServer(grpcServer, &pb.UnimplementedProductServiceServer{})
	pb.RegisterOrderServiceServer(grpcServer, &pb.UnimplementedOrderServiceServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Configurar servidor REST
	app := fiber.New()
	// Enable logger
	app.Use(logger.New())

	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Enable limiter
	app.Use(limiter.New())

	config.NewSession()

	// Custom storage
	store := config.NewStore()

	// Configuring limiter
	app.Use(limiter.New(limiter.Config{
		Max:        150,
		Expiration: 10 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
		Storage: store,
	}))

	// Grouping routes by version
	apiV1 := app.Group("/api/v1")

	// Grouping routes by resource
	auth := apiV1.Group("/auth")
	products := apiV1.Group("/products")
	orders := apiV1.Group("/orders")

	// Registering routers
	router.AuthRouter(auth)
	router.ProductsRouter(products)
	router.OrdersRouter(orders)

	// Starting server
	go func() {
		log.Println("Starting HTTP server on port 8080")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Starting gRPC server
	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
