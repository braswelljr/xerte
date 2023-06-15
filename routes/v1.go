package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/braswelljr/xerte/controller/v1/authentication"
)

// Version1 - Version 1 (prefix - v1)
//
//	@param router fiber.Router
//	@return void
func Version1(router fiber.Router) {
	// authentication
	auth := router.Group("/authentication")
	// login
	{
		auth.Post("/login", authentication.Login())
	}
}
