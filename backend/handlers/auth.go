package handlers

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
    "TaskManagementSystem/models"
    "TaskManagementSystem/utils"
    "TaskManagementSystem/config"
)

func Register(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Validate required fields
    if user.Email == "" || user.Password == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "Email and password are required",
        })
    }

    // Check if email already exists
    var existingUser models.User
    err := config.DB.Collection("users").FindOne(c.Context(), bson.M{
        "email": user.Email,
    }).Decode(&existingUser)

    if err == nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Email already exists",
        })
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not hash password",
        })
    }
    user.Password = string(hashedPassword)

    // Add timestamps
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    // Save user to database
    result, err := config.DB.Collection("users").InsertOne(c.Context(), user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not create user",
        })
    }

    return c.JSON(fiber.Map{
        "id": result.InsertedID,
        "message": "User created successfully",
    })
}

func Login(c *fiber.Ctx) error {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&loginData); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Validate required fields
    if loginData.Email == "" || loginData.Password == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "Email and password are required",
        })
    }

    // Find user
    var user models.User
    err := config.DB.Collection("users").FindOne(c.Context(), bson.M{
        "email": loginData.Email,
    }).Decode(&user)

    if err != nil {
        return c.Status(401).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
        return c.Status(401).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }

    // Generate token
    token, err := utils.GenerateToken(user.ID.Hex())
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not generate token",
        })
    }

    return c.JSON(fiber.Map{
        "token": token,
        "user": fiber.Map{
            "id": user.ID,
            "email": user.Email,
            "name": user.Name,
        },
    })
}
