package handlers

import (
	"net/http"

	"github.com/HsiaoCz/twitter-go/db"
	"github.com/HsiaoCz/twitter-go/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	store *db.Store
}

func NewUserHandlers(store *db.Store) *UserHandlers {
	return &UserHandlers{
		store: store,
	}
}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	createUserParam := types.CreateUserParams{}
	if err := c.BodyParser(&createUserParam); err != nil {
		return NewAPIError(http.StatusBadRequest, "please check the request params")
	}
	msg := createUserParam.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	return nil
}
