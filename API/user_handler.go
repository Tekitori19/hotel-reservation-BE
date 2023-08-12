package API

import (
	"github.com/Tekitori19/hotel-reservation-BE/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	users := types.User{
		FirstName: "Guma",
		LastName:  "gala",
	}
	return c.JSON(users)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("1 user")
}
