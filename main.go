package main

import (
	"net/http"

	"github.com/HsiaoCz/twitter-go/handlers"
	"github.com/gofiber/fiber/v2"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(handlers.APIError); ok {
			return c.Status(apiError.Status).JSON(&apiError)
		}
		aErr := handlers.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(aErr.Status).JSON(&aErr)
	},
}

func main() {
	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "all is well",
		})
	})

	app.Listen(":3001")
}
