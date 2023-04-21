package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Routes handles application routes.
// - APIs are prefixed with `api`.
// - Versions are prefixed with `v(number)`. Example `v1`, `v2`.
func Routes(app *fiber.App) {
	// API Routes with /api
	api := app.Group("/api")
	// Versioning
	// Version 1 (prefix - v1)
	_ = api.Group("/v1")

}
