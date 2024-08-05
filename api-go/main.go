package main

import (
	"go-api/handlers"
	"go-api/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Post("/auth", handlers.AuthHandler)

	// Middleware para validar JWT
	app.Use(middleware.ValidateJWT)

	app.Post("/matrix", handlers.MatrixHandler)

	log.Fatal(app.Listen(":8080"))
}
