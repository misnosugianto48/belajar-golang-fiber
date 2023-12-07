package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("i'am middleware before processing request")
		err := c.Next()
		fmt.Println("i'am middleware after processing request")
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	if fiber.IsChild() {
		fmt.Println("i a child process")
	} else {
		fmt.Println("im a parent process")
	}

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	// log.Panic(app.Listen(":3000"))

}
