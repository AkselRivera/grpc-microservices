package router

import (
	"github.com/akselrivera/grpc-apigateway/internal/handlers"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/auth"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/validation"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {

	router.Get("/me", auth.PrivateRoute, handlers.ValidateToken)

	router.Post("/login", validation.ValidateLogin, handlers.Login)

	router.Post("/register", validation.ValidateRegister, handlers.Register)

	router.Post("/logout", auth.PrivateRoute, handlers.Logout)

}
