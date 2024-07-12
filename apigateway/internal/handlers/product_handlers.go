package handlers

import (
	"net/http"

	"github.com/akselrivera/grpc-apigateway/internal/handlers/errors"
	pb "github.com/akselrivera/grpc-apigateway/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var productClient pb.ProductServiceClient

func init() {
	// product-service
	connection, err := grpc.NewClient("product-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	productClient = pb.NewProductServiceClient(connection)
}

func GetProducts(c *fiber.Ctx) error {
	resp, err := productClient.ListProducts(c.Context(), &pb.ListProductsRequest{})
	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&resp.Products)
}

func GetProduct(c *fiber.Ctx) error {
	resp, err := productClient.GetProduct(c.Context(), &pb.GetProductRequest{
		Id: c.Params("id")})
	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&resp.Product)
}

func CreateProduct(c *fiber.Ctx) error {
	var req pb.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Couldn't parse body - API Gateway",
			"error":   err,
		})
	}

	resp, err := productClient.CreateProduct(c.Context(), &req)
	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(&resp.Product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var req pb.Product
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't parse body - API Gateway",
			"error":   err,
		})
	}

	resp, err := productClient.UpdateProduct(c.Context(), &pb.UpdateProductRequest{
		Product: &pb.Product{
			Id:          c.Params("id"),
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Quantity:    req.Quantity,
			IsActive:    req.IsActive,
		},
	})

	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(&resp.Product)
}

func DeleteProduct(c *fiber.Ctx) error {
	resp, err := productClient.DeleteProduct(c.Context(), &pb.DeleteProductRequest{
		Id: c.Params("id")})

	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product deleted",
		"success": resp.Success,
	})
}

func HandleStock(c *fiber.Ctx) error {
	var req pb.HandleProductStockRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Couldn't parse body - API Gateway",
			"error":   err,
		})
	}

	resp, err := productClient.HandleProductStock(c.Context(), &pb.HandleProductStockRequest{
		Id:       c.Params("id"),
		Quantity: req.Quantity,
		Action:   req.Action,
	})

	if err != nil {
		return errors.HandleProductError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(&resp.Success)
}
