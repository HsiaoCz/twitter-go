package handlers

import (
	"github.com/HsiaoCz/twitter-go/db"
	"github.com/gofiber/fiber/v2"
)

type PostHandlers struct {
	post db.PostStorer
}

func PostHandlersInit(post db.PostStorer) *PostHandlers {
	return &PostHandlers{
		post: post,
	}
}

func (p *PostHandlers) HandleCreatePost(c *fiber.Ctx) error {
	return nil
}
