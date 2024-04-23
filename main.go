package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"fmt"
	"net/http"
)
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
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
	
	app.Listen(":"+getEnv("PORT", "8080"))
	{

    http.HandleFunc("/health", healthCheckHandler)

    fmt.Println("Starting server on port 8080...")
    http.ListenAndServe(":8080", nil)
}
}
