package middleware

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
    "os"
)

func AuthMiddleware() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
        ErrorHandler: jwtError,
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
