const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export interface Task {
    _id: string;
    title: string;
    description: string;
    status: 'pending' | 'in_progress' | 'completed';
    priority: 'low' | 'medium' | 'high';
    due_date: string;
    assigned_to: string;
    created_at: string;
    updated_at: string;
}

export const api = {
    // Auth endpoints
    register: async (userData: { email: string; password: string }) => {
        const response = await fetch(`${API_URL}/auth/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(userData),
        });
        return response.json();
    },

    login: async (credentials: { email: string; password: string }) => {
        const response = await fetch(`${API_URL}/auth/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(credentials),
        });
        return response.json();
    },

    // Tasks endpoints
    getTasks: async (token: string) => {
        const response = await fetch(`${API_URL}/api/tasks`, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        });
        return response.json();
    },

    createTask: async (task: Partial<Task>, token: string) => {
        const response = await fetch(`${API_URL}/api/tasks`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(task),
        });
        return response.json();
    },

    updateTask: async (taskId: string, updates: Partial<Task>, token: string) => {
        const response = await fetch(`${API_URL}/api/tasks/${taskId}`, {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updates),
        });
        return response.json();
    },

    deleteTask: async (taskId: string, token: string) => {
        const response = await fetch(`${API_URL}/api/tasks/${taskId}`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        });
        return response.json();
    },
};