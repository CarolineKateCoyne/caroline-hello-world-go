package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"fmt"
)

func healthCheckHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", nil)
	})

	app.Get("/health", healthCheckHandler)

	port := getEnv("PORT", "8080")
	fmt.Println("Starting server on port", port)
	app.Listen(":" + port)
}
