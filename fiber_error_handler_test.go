package main

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app2 = fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error : " + err.Error())
	},
	Views: engine,
})

func TestError(t *testing.T) {
	app2.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("internal server error")
	})

	req := httptest.NewRequest("GET", "/error", nil)
	res, err := app2.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 500, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error : internal server error", string(bytes))

}
