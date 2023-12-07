package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app := fiber.New()
	app.Post("/register", func(c *fiber.Ctx) error {
		req := new(RegisterRequest)
		err := c.BodyParser(req)
		if err != nil {
			return nil
		}

		return c.SendString("Register Success " + req.Username)
	})
}

func TestBodyParserJson(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username":"Misno", "password":"rahasia", "name":"Misno" Sugianto}`)
	req := httptest.NewRequest("POST", "/register", body)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Misno", string(bytes))
}
