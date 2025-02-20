package middleware

import (
	"os"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			// Get the JWT token claims with proper error handling
			user := c.Locals("user")
			if user == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Missing token data",
				})
			}

			token, ok := user.(*jwt.Token)
			if !ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Invalid token format",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Invalid claims format",
				})
			}
			
			// Set the user_id in context with nil check
			if userId, exists := claims["user_id"]; exists {
				c.Locals("user_id", userId)
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "User ID not found in token",
				})
			}
			
			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT",
	})
}
