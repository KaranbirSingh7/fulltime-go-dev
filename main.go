package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/karanbirsingh7/fulltime-go-dev/api"
)

var (
	PORT int
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

	// APP: routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": "running",
		})
	})

	appV1 := app.Group("/api/v1")
	appV1.Get("/user", api.HandleGetUser)

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
