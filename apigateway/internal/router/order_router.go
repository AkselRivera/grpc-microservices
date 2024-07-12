package router

import (
	"github.com/akselrivera/grpc-apigateway/internal/handlers"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/auth"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/validation"
	"github.com/gofiber/fiber/v2"
)

func OrdersRouter(router fiber.Router) {
	// Implementar manejo de solicitudes HTTP
	router.Get("/", auth.PrivateRoute, handlers.ListOrders)
	router.Get("/:id", auth.PrivateRoute, handlers.GetOrder)

	router.Post("/", auth.PrivateRoute, validation.ValidateOrder, handlers.CreateOrder)

	router.Patch("/:id", auth.PrivateRoute, validation.ValidateChangeStatus, handlers.ChangeOrderStatus)
	router.Delete("/:id", auth.PrivateRoute, handlers.CancelOrder)

}
