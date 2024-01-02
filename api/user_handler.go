package api

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	c.JSON(map[string]interface{}{
		"username": "john",
	})
	return nil
}
