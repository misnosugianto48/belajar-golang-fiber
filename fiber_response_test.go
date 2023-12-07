package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestResponseJson(t *testing.T) {
	app := fiber.New()
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name": "Misno Sugianto",
		})
	})

	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Set("Accept", "application/json")
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Misno Sugianto"}`, string(bytes))
}
