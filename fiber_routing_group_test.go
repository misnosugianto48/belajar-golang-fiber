package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	helloWorld := func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")
	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)

	req := httptest.NewRequest("GET", "/api/hello", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello world", string(bytes))

}
