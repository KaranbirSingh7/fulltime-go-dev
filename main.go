package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/karanbirsingh7/fulltime-go-dev/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	PORT int
)

const (
	DB_URI          = "mongodb://localhost:27017"
	DB_NAME         = "hotel-reservation"
	USER_COLLECTION = "users"
)

func main() {
	println("starting...")

	// flags for configuration
	flag.IntVar(&PORT, "port", 8080, "port to use for running server")
	flag.Parse()

	// server basic configurations
	app := fiber.New(
		fiber.Config{},
	)

	// DATABASE configuration
	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Panic(err)
	}
	dbClient.Database(DB_NAME).Collection(USER_COLLECTION)

	// APP: routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": "running",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		var filter any
		databases, err := dbClient.ListDatabaseNames(context.TODO(), filter, options.ListDatabases())
		if err != nil {
			c.JSON(map[string]interface{}{
				"error": err,
			})
		}
		return c.JSON(databases)
	})

	appV1 := app.Group("/api/v1")
	appV1.Get("/user", api.HandleGetUsers)
	appV1.Get("/user/:id", api.HandleGetUser)

	// start server in background
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", PORT)); err != nil {
			log.Panic(err)
		}
	}()

	// graceful shutdown logic
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// block program until we receive a failure or termination signal
	<-signalChan

	if err := app.ShutdownWithTimeout(3 * time.Second); err != nil {
		log.Panic(err)
	}

	fmt.Println("Received termination signal. Initiating graceful shutdown...")
	fmt.Println("Server gracefully shutdown")
}
