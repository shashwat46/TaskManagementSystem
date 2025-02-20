package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "TaskManagementSystem/config"
    "TaskManagementSystem/models"
    "time"
)

// GetTasks retrieves all tasks for the authenticated user
func GetTasks(c *fiber.Ctx) error {
    // Get user ID from JWT token
    userId := c.Locals("user_id").(string)
    objID, _ := primitive.ObjectIDFromHex(userId)

    // Query tasks
    cursor, err := config.DB.Collection("tasks").Find(c.Context(), bson.M{
        "assigned_to": objID,
    })
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not fetch tasks",
        })
    }

    var tasks []models.Task
    if err := cursor.All(c.Context(), &tasks); err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Error processing tasks",
        })
    }

    return c.JSON(tasks)
}

// CreateTask creates a new task
func CreateTask(c *fiber.Ctx) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Set task metadata
    userId := c.Locals("user_id").(string)
    objID, _ := primitive.ObjectIDFromHex(userId)
    task.AssignedTo = objID
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()
    task.Status = "pending"

    // Insert task
    result, err := config.DB.Collection("tasks").InsertOne(c.Context(), task)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not create task",
        })
    }

    task.ID = result.InsertedID.(primitive.ObjectID)
    return c.Status(201).JSON(task)
}

// GetTask retrieves a single task by ID
func GetTask(c *fiber.Ctx) error {
    taskId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid task ID",
        })
    }

    userId := c.Locals("user_id").(string)
    objID, _ := primitive.ObjectIDFromHex(userId)

    var task models.Task
    err = config.DB.Collection("tasks").FindOne(c.Context(), bson.M{
        "_id": taskId,
        "assigned_to": objID,
    }).Decode(&task)

    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(task)
}

// UpdateTask updates an existing task
func UpdateTask(c *fiber.Ctx) error {
    taskId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid task ID",
        })
    }

    var updates struct {
        Title       string    `json:"title"`
        Description string    `json:"description"`
        Status      string    `json:"status"`
        DueDate     time.Time `json:"due_date"`
        Priority    string    `json:"priority"`
    }

    if err := c.BodyParser(&updates); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    userId := c.Locals("user_id").(string)
    objID, _ := primitive.ObjectIDFromHex(userId)

    result, err := config.DB.Collection("tasks").UpdateOne(
        c.Context(),
        bson.M{
            "_id": taskId,
            "assigned_to": objID,
        },
        bson.M{
            "$set": bson.M{
                "title":       updates.Title,
                "description": updates.Description,
                "status":      updates.Status,
                "due_date":    updates.DueDate,
                "priority":    updates.Priority,
                "updated_at":  time.Now(),
            },
        },
    )

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not update task",
        })
    }

    if result.MatchedCount == 0 {
        return c.Status(404).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Task updated successfully",
    })
}

// DeleteTask deletes a task
func DeleteTask(c *fiber.Ctx) error {
    taskId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Invalid task ID",
        })
    }

    userId := c.Locals("user_id").(string)
    objID, _ := primitive.ObjectIDFromHex(userId)

    result, err := config.DB.Collection("tasks").DeleteOne(c.Context(), bson.M{
        "_id": taskId,
        "assigned_to": objID,
    })

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Could not delete task",
        })
    }

    if result.DeletedCount == 0 {
        return c.Status(404).JSON(fiber.Map{
            "error": "Task not found",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Task deleted successfully",
    })
}
