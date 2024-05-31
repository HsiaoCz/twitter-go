package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/twitter-go/db"
	"github.com/HsiaoCz/twitter-go/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(handlers.APIError); ok {
			return c.Status(apiError.Status).JSON(&apiError)
		}
		aErr := handlers.NewAPIError(http.StatusInternalServerError, err.Error())
		return c.Status(aErr.Status).JSON(&aErr)
	},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userColl       = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		mongoUserStore = db.NewMongoUserStore(client, userColl)
		store          = &db.Store{User: mongoUserStore}
		userHandlers   = handlers.NewUserHandlers(store)
		app            = fiber.New(config)
		v1             = app.Group("/app/v1")
	)

	// routers
	{
		v1.Post("/user", userHandlers.HandleCreateUser)
	}

	go func() {
		if err := app.Listen(os.Getenv("PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
