package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestReqBody(t *testing.T) {
	app := fiber.New()
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()
		req := new(LoginRequest)
		err := json.Unmarshal(body, req)
		if err != nil {
			return err
		}
		return c.SendString("Login Success " + req.Username)
	})

	body := strings.NewReader(`{
		"username" : "Misno",
		"password" :"Misno"
	}`)
	req := httptest.NewRequest("POST", "/login", body)
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Login Success Misno", string(bytes))
}
