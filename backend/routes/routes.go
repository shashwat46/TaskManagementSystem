package routes

import (
    "github.com/gofiber/fiber/v2"
    "TaskManagementSystem/handlers"
    "TaskManagementSystem/middleware"
)

func SetupRoutes(app *fiber.App) {
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
}
