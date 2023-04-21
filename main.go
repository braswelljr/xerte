package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/braswelljr/xte/routes"
)

func main() {
	// load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// get port from env
	PORT := os.Getenv("PORT")

	// Initialize app
	app := fiber.New()

	// Logger
	app.Use(logger.New())

	// add CORS
	app.Use(cors.New())

	// add static files
	app.Static("/", "./static")

	// index route
	app.All("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Welcome to the axxxe api",
		})
	})

	// handle routes
	routes.Routes(app)

	// Listen on port
	if err := app.Listen(fmt.Sprintf(":%v", PORT)); err != nil {
		log.Fatal("Something went wrong  ", err)
	}
}
