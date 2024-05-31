package handlers

import (
	"github.com/HsiaoCz/twitter-go/db"
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
	return nil
}
