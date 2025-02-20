package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Email     string            `json:"email" bson:"email"`
    Password  string            `json:"password" bson:"password"`
    Name      string            `json:"name" bson:"name"`
    CreatedAt time.Time         `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}

type Task struct {
    ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
    Title       string              `json:"title" bson:"title"`
    Description string              `json:"description" bson:"description"`
    Status      string              `json:"status" bson:"status"`
    DueDate     time.Time           `json:"due_date" bson:"due_date"`
    AssignedTo  primitive.ObjectID  `json:"assigned_to" bson:"assigned_to"`
    Priority    string              `json:"priority" bson:"priority"`
    AITags      []string            `json:"ai_tags" bson:"ai_tags"`
    CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time           `json:"updated_at" bson:"updated_at"`
}
