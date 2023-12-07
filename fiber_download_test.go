package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	app := fiber.New()
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/contoh.txt", "contoh.txt")
	})

	req := httptest.NewRequest("GET", "/download", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, "attachment; filename=\"contoh.txt\"", res.Header.Get("Content-Disposition"))
	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "this is sample file for upload", string(bytes))
}
