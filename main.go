package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"log"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, Template",
		})
	})

	messageHandler := func(c *fiber.Ctx) error {
		return c.SendString("Hello, HTMX")
	}

	app.Get("/message", messageHandler)

	log.Fatal(app.Listen(":9000"))
}
