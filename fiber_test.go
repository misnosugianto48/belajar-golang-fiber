package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestRouteParams(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " from User " + userId)
	})

	req := httptest.NewRequest("GET", "/users/misno/orders/3", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 3 from User misno", string(bytes))
}

func TestReqForm(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Misno")
	req := httptest.NewRequest("POST", "/hello", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Misno", string(bytes))
}
