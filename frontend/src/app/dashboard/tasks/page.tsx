"use client";

import React, { useState, useEffect } from "react";
import CreateTaskModal from "@/app/dashboard/tasks/CreateTaskModal";

// Define the Task interface
interface Task {
  id: string;
  title: string;
  description: string;
  status: 'pending' | 'in_progress' | 'completed';
  priority: 'low' | 'medium' | 'high';
  due_date: string;
  assigned_to: string;
}

export default function TasksPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  // Fetch tasks when component mounts
  useEffect(() => {
    const fetchTasks = async () => {
      setLoading(true);
      try {
        const token = localStorage.getItem("token");
        const response = await fetch("http://localhost:8080/api/tasks", {
          headers: {
            "Authorization": `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch tasks");
        }

        const data = await response.json();
        setTasks(data);
      } catch (error) {
        setError("Failed to fetch tasks");
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchTasks();
  }, []);

  const handleCreateTask = async (taskData: Partial<Task>) => {
    setLoading(true);
    setError("");
    
    try {
      const token = localStorage.getItem("token");
      const response = await fetch("http://localhost:8080/api/tasks", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(taskData),
      });

      if (!response.ok) {
        throw new Error("Failed to create task");
      }

      const newTask = await response.json();
      setTasks((prevTasks) => [...prevTasks, newTask]);
      setIsModalOpen(false);
      
    } catch (error) {
      setError("Failed to create task. Please try again.");
      console.error("Error creating task:", error);
    } finally {
      setLoading(false);
    }
  };

  if (loading && tasks.length === 0) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <div className="text-gray-500">Loading tasks...</div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="bg-white shadow rounded-lg p-6">
        <div className="flex justify-between items-center">
          <h2 className="text-2xl font-bold">Your Tasks</h2>
          <button 
            onClick={() => setIsModalOpen(true)}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            disabled={loading}
          >
            {loading ? 'Creating...' : 'Add New Task'}
          </button>
        </div>

        {error && (
          <div className="mt-4 p-4 text-red-700 bg-red-100 rounded-md">
            {error}
          </div>
        )}
      </div>

      {/* Task List */}
      <div className="bg-white shadow rounded-lg p-6">
        <div className="space-y-4">
          {tasks.length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-500">No tasks found. Create a new task to get started!</p>
            </div>
          ) : (
            <div className="space-y-4">
              {tasks.map((task: Task) => (
                <div 
                  key={task.id} 
                  className="border rounded-lg p-4 hover:shadow-md transition-shadow"
                >
                  <div className="flex justify-between items-start">
                    <div>
                      <h3 className="font-semibold">{task.title}</h3>
                      <p className="text-gray-600">{task.description}</p>
                    </div>
                    <span className={`px-2 py-1 rounded text-sm ${
                      task.priority === 'high' ? 'bg-red-100 text-red-800' :
                      task.priority === 'medium' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-green-100 text-green-800'
                    }`}>
                      {task.priority}
                    </span>
                  </div>
                  <div className="mt-2 flex justify-between items-center text-sm text-gray-500">
                    <span>Due: {new Date(task.due_date).toLocaleDateString()}</span>
                    <span className={`px-2 py-1 rounded ${
                      task.status === 'completed' ? 'bg-green-100 text-green-800' :
                      task.status === 'in_progress' ? 'bg-blue-100 text-blue-800' :
                      'bg-gray-100 text-gray-800'
                    }`}>
                      {task.status}
                    </span>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Create Task Modal */}
      <CreateTaskModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleCreateTask}
      />
    </div>
  );
}
