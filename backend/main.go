package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
    "TaskManagementSystem/config"
    "TaskManagementSystem/middleware"
    "TaskManagementSystem/handlers"
    "TaskManagementSystem/routes"

)

func main() {
    // Load env variables
    if err := godotenv.Load(); err != nil {
        log.Printf("Warning: .env file not found")
    }

    // Get PORT from environment with fallback
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // fallback port
    }

    // Connect to MongoDB
    config.ConnectDB()

    // Initialize Fiber
    app := fiber.New(fiber.Config{
        AppName: "Task Management API",
    })

    // Global Middleware
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "https://task-management-system-ecru.vercel.app",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
    }))

    // Setup Routes
    routes.SetupRoutes(app)

    // Public Routes
    auth := app.Group("/auth")
    auth.Post("/register", handlers.Register)
    auth.Post("/login", handlers.Login)

    // Protected Routes
    api := app.Group("/api", middleware.AuthMiddleware())
    
    // Task routes
    tasks := api.Group("/tasks")
    tasks.Get("/", handlers.GetTasks)
    tasks.Post("/", handlers.CreateTask)
    tasks.Get("/:id", handlers.GetTask)
    tasks.Put("/:id", handlers.UpdateTask)
    tasks.Delete("/:id", handlers.DeleteTask)

    // User routes
    users := api.Group("/users")
    users.Get("/me", handlers.GetCurrentUser)
    users.Put("/me", handlers.UpdateUser)

    // Start server with proper port format
    log.Fatal(app.Listen(":" + port))
}
