package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Routes handles application routes.
//
//	APIs are prefixed with `api`.
//	Versions are prefixed with `v(number)`. Example `v1`, `v2`.
//	@param app fiber.App
//	@return void
func Routes(app *fiber.App) {
	// API Routes with /api
	API := app.Group("/api")
	// Versioning
	// Version 1 (prefix - v1)
	Version1(API.Group("/v1"))
	// Version 2 (prefix - v2)
	Version2(API.Group("/v2"))
}
