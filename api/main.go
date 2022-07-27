package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/manimovassagh/shortenURLFiberRedis/routes"
)

func setupRoutes(app *fiber.app) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Listen("3000")
}
