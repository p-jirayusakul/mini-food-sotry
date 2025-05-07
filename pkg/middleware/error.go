package middleware

import (
	"errors"
	"food-story/pkg/exceptions"

	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error) error {
	// Default to 500 Internal Server Error
	code := fiber.StatusInternalServerError
	message := "Unknown error"

	// Ensure error is not nil
	if err != nil {
		// Retrieve the custom status code if it's a fiber.*Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}
		message = getMessage(err, code)
	}
	return c.Status(code).JSON(ErrorResponse{Message: message, Status: "error"})
}

func getMessage(err error, code int) string {
	if code == fiber.StatusInternalServerError {
		return exceptions.ErrInternalServerError.Error()
	}
	return err.Error()
}
