package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestRoutingHelloWorld(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	req := httptest.NewRequest("GET", "/", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestContext(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello " + name)
	})

	req := httptest.NewRequest("GET", "/hello?name=Misno", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Misno", string(bytes))

	req = httptest.NewRequest("GET", "/hello", nil)
	res, err = app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err = io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest", string(bytes))
}

func TestHttpReq(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("Hello " + first + " " + last)
	})

	req := httptest.NewRequest("GET", "/request", nil)
	req.Header.Set("firstname", "Misno")
	req.AddCookie(&http.Cookie{Name: "lastname", Value: "Sugianto"})
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Misno Sugianto", string(bytes))
}
