package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "TaskManagementSystem/config"
    "TaskManagementSystem/models"
	"time"
)

// GetCurrentUser retrieves the authenticated user's profile
func GetCurrentUser(c *fiber.Ctx) error {
    // Get user ID from JWT token
    userId := c.Locals("user_id").(string)
    objID, err := primitive.ObjectIDFromHex(userId)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    var user models.User
    err = config.DB.Collection("users").FindOne(c.Context(), bson.M{
        "_id": objID,
    }).Decode(&user)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    // Don't send password in response
    user.Password = ""
    
    return c.JSON(user)
}

// UpdateUser updates the authenticated user's profile
func UpdateUser(c *fiber.Ctx) error {
    userId := c.Locals("user_id").(string)
    objID, err := primitive.ObjectIDFromHex(userId)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    var updates struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := c.BodyParser(&updates); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Check if email is already taken
    if updates.Email != "" {
        count, err := config.DB.Collection("users").CountDocuments(c.Context(), bson.M{
            "_id":   bson.M{"$ne": objID},
            "email": updates.Email,
        })
        if err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": "Database error",
            })
        }
        if count > 0 {
            return c.Status(400).JSON(fiber.Map{
                "error": "Email already taken",
            })
        }
    }

    result, err := config.DB.Collection("users").UpdateOne(
        c.Context(),
        bson.M{"_id": objID},
        bson.M{
            "$set": bson.M{
                "name":       updates.Name,
                "email":      updates.Email,
                "updated_at": time.Now(),
            },
        },
    )

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not update user",
        })
    }

    if result.MatchedCount == 0 {
        return c.Status(404).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "User updated successfully",
    })
}
