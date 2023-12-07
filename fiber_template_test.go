package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("./template", ".mustache")

var app3 = fiber.New(fiber.Config{
	Views: engine,
})

func TestViews(t *testing.T) {
	app3.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello title",
			"header":  "Hello header",
			"content": "Hello content",
		})
	})

	req := httptest.NewRequest("GET", "/view", nil)
	res, err := app3.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello title")
	assert.Contains(t, string(bytes), "Hello header")
	assert.Contains(t, string(bytes), "Hello content")

}
