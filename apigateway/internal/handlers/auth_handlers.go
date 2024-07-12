package handlers

import (
	"net/http"
	"time"

	"github.com/akselrivera/grpc-apigateway/internal/config"
	"github.com/akselrivera/grpc-apigateway/internal/handlers/errors"
	"github.com/akselrivera/grpc-apigateway/internal/models"
	pb "github.com/akselrivera/grpc-apigateway/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authClient pb.AuthServiceClient

func init() {
	// auth-service
	connection, err := grpc.NewClient("auth-service:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	authClient = pb.NewAuthServiceClient(connection)
}

func Login(c *fiber.Ctx) error {
	var req pb.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Couldn't parse body - microservice",
			"error":   err,
		})
	}

	resp, err := authClient.Login(c.Context(), &pb.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return errors.HandleAuthError(c, err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    resp.Token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	session := config.GetSession(c)

	session.Set("email", req.Email)
	session.Save()

	return c.Status(http.StatusOK).JSON(&resp)
}

func Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("Error parsing body: %v", err)
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Couldn't parse body - microservice",
			"error":   err,
		})
	}

	if req.Password != req.ConfirmPassword {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Passwords don't match",
		})
	}

	resp, err := authClient.Register(c.Context(), &pb.RegisterRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return errors.HandleAuthError(c, err)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    resp.Token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	session := config.GetSession(c)

	session.Set("email", req.Email)
	session.Save()

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})

}

func Logout(c *fiber.Ctx) error {
	session := config.GetSession(c)
	session.Destroy()

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	})

	return c.Status(http.StatusNoContent).Send(nil)
}

func ValidateToken(c *fiber.Ctx) error {
	cookie := c.Cookies("token")

	resp, err := authClient.ValidateToken(c.Context(), &pb.ValidateTokenRequest{Token: cookie})
	if err != nil {
		return errors.HandleAuthError(c, err)
	}

	return c.Status(http.StatusOK).JSON(&resp)
}

func ValidateTokenMiddleware(c *fiber.Ctx) bool {
	cookie := c.Cookies("token")

	resp, err := authClient.ValidateToken(c.Context(), &pb.ValidateTokenRequest{Token: cookie})

	if err != nil {
		return false
	}

	return resp.Valid
}
