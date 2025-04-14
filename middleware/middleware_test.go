package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleCorsMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(HandleCorsMiddleware())

	// Add dummy route
	app.Options("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	// Simulate a preflight request
	req := httptest.NewRequest(http.MethodOptions, "/", http.NoBody)
	req.Header.Set("Origin", "http://example.com")
	req.Header.Set("Access-Control-Request-Method", "GET")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Now CORS headers will exist
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "GET")
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Headers"), "Origin")
}
