package authentication

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Login - Login is a function that handles the login of users.
//
//	@return fiber.Handler
//	@return error
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// create a context with a timeout of 10 seconds for the request
		_, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		// get the basic auth credentials

		return nil
	}
}
