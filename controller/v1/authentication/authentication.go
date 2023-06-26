package authentication

import (
	"bytes"
	"context"
	"encoding/base64"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/braswelljr/xerte/firebase"
	"github.com/braswelljr/xerte/models"
)

var (
	basicAuthPrefix = []byte("Basic ")
)

// Login - Login is a function that handles the login of users.
//
//	@return fiber.Handler
//	@return error
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// create a context with a timeout of 10 seconds for the request
		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		// firestore client
		auth, err := firebase.InitAuth(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		// get the basic authorization credentials
		authorization := c.Request().Header.Peek("Authorization")

		data := &models.LoginReq{}

		if bytes.HasPrefix(authorization, basicAuthPrefix) {
			// Decode to get the raw token
			token, err := base64.StdEncoding.DecodeString(string(authorization[len(basicAuthPrefix):]))
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  fiber.StatusUnauthorized,
					"message": err.Error(),
				})
			}

			// Split the token into its parts (username and password)
			credentials := bytes.Split(token, []byte(":"))
			if len(credentials) != 2 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  fiber.StatusUnauthorized,
					"message": "invalid authorization credentials",
				})
			}

			// assign the credentials to the data struct
			data.Email = string(credentials[0])
			data.Password = string(credentials[1])
		}

		// validate user details
		if err := validator.New().Struct(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// login the user
		user, err := auth.Login(ctx, data.Email, data.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "user logged in successfully",
			"data":    user,
		})
	}
}

// Signup - Signup is a function that handles the signup of users.
//
//	@return fiber.Handler
//	@return error
func Signup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// create a context with a timeout of 10 seconds for the request
		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
		defer cancel()

		// fireauth client
		auth, err := firebase.InitAuth(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		// create a new instance of the user model
		data := new(models.SignupRequest)

		// get the request body
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// validate user details
		if err := validator.New().Struct(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// create the user
		user, err := auth.Signup(ctx, data)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  fiber.StatusCreated,
			"message": "user created successfully",
			"data":    user,
		})
	}
}
