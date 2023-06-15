package authentication

import "github.com/gofiber/fiber/v2"

// Login - Login is a function that handles the login of users.
//
//	@return fiber.Handler
//	@return error
func Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
