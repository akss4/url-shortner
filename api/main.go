package main

import (
	"log"
	"os"

	"github.com/akss4/url_shortner/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {

	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenUrl)

}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load env")

	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"

	}
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Printf("Server listening on %s", port)
	log.Fatal(app.Listen(port))

}
