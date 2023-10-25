package main

import (
	"bytes"
	_ "embed"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//go:embed source/contoh.txt
var contohFile []byte

func TestMultipartForm(t *testing.T) {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	file, err := w.CreateFormFile("file", "contoh.txt")
	assert.Nil(t, err)
	file.Write(contohFile)
	w.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}
