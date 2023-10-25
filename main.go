package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}
