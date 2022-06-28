package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		msg := SomeFunction(c.Context())
		return c.SendString(msg)
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}

}

func SomeFunction(_ context.Context) string {
	return "Hello, World!"
}
