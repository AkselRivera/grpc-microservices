package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory/v2"
)

var Session *session.Store

func NewSession() {
	Session = session.New(session.Config{
		Storage: memory.New(),
	})

}

func GetSession(c *fiber.Ctx) *session.Session {
	session, err := Session.Get(c)
	if err != nil {
		log.Errorf("Error getting session: %v", err)
	}

	return session
}
