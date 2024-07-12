package handlers

import (
	"fmt"
	"net/http"

	"github.com/akselrivera/grpc-apigateway/internal/handlers/errors"
	"github.com/akselrivera/grpc-apigateway/internal/models"
	pb "github.com/akselrivera/grpc-apigateway/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var orderClient pb.OrderServiceClient

func init() {
	// order-service
	connection, err := grpc.NewClient("order-service:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	orderClient = pb.NewOrderServiceClient(connection)
}

func GetOrder(c *fiber.Ctx) error {
	resp, err := orderClient.GetOrder(c.Context(), &pb.GetOrderRequest{Id: c.Params("id")})
	if err != nil {
		return errors.HandleOrderError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&resp)
}

func ListOrders(c *fiber.Ctx) error {
	resp, err := orderClient.ListOrders(c.Context(), &pb.ListOrdersRequest{})
	if err != nil {
		return errors.HandleOrderError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&resp.Orders)
}

func CreateOrder(c *fiber.Ctx) error {
	var req pb.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Couldn't parse body - microservice",
			"error":   err,
		})
	}

	resp, err := orderClient.CreateOrder(c.Context(), &req)
	if err != nil {
		return errors.HandleOrderError(c, err)
	}
	fmt.Println(resp)

	return c.Status(http.StatusCreated).JSON(&resp.Order)
}

func ChangeOrderStatus(c *fiber.Ctx) error {
	var req models.ChangeOrderStatus
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Couldn't parse body - microservice",
			"error":   err,
		})
	}

	resp, err := orderClient.ChangeOrderStatus(c.Context(), &pb.ChangeOrderStatusRequest{
		Id:     c.Params("id"),
		Status: pb.EnumStatus(req.Status),
	})

	if err != nil {
		return errors.HandleOrderError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(&resp.Order)
}

func CancelOrder(c *fiber.Ctx) error {
	resp, err := orderClient.CancelOrder(c.Context(), &pb.CancelOrderRequest{
		Id: c.Params("id")})

	if err != nil {
		return errors.HandleOrderError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product deleted",
		"success": resp.Success,
	})
}
