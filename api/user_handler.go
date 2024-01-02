package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karanbirsingh7/fulltime-go-dev/types"
)

func HandleGetUser(c *fiber.Ctx) error {
	c.JSON(map[string]interface{}{
		"username": "john",
	})
	return nil
}

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "john",
		LastName:  "doe",
	}
	c.JSON(u)
	return nil
}
