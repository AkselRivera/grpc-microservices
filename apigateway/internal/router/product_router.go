package router

import (
	"github.com/akselrivera/grpc-apigateway/internal/handlers"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/auth"
	"github.com/akselrivera/grpc-apigateway/internal/middlewares/validation"
	"github.com/gofiber/fiber/v2"
)

func ProductsRouter(router fiber.Router) {
	router.Get("/", auth.PrivateRoute, handlers.GetProducts)
	router.Get("/:id", auth.PrivateRoute, handlers.GetProduct)

	router.Post("/", auth.PrivateRoute, validation.ValidateProductWithoutId, handlers.CreateProduct)

	router.Patch("/stock/:id", auth.PrivateRoute, validation.ValidateProductStock, handlers.HandleStock)
	router.Patch("/:id", auth.PrivateRoute, validation.ValidateProduct, handlers.UpdateProduct)
	router.Delete("/:id", auth.PrivateRoute, handlers.DeleteProduct)

}
