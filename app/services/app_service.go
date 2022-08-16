package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wily13/go-whatsapp-api/structs"
)

type AppService interface {
	Login(c *fiber.Ctx) (response structs.LoginResponse, err error)
	Logout(c *fiber.Ctx) (err error)
	Reconnect(c *fiber.Ctx) (err error)
}
