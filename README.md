# AI-Powered Task Management System

## Live Demo
- Frontend: [https://task-management-system-ecru.vercel.app/](https://task-management-system-ecru.vercel.app/)
- Backend API: [https://taskmanagementsystem-production-8e99.up.railway.app](https://taskmanagementsystem-production-8e99.up.railway.app)

## Deployment
The application is deployed using:
- Frontend: Vercel (Next.js optimized hosting)
- Backend: Railway (Go server hosting)
- Database: MongoDB Atlas

## Core Features Implemented

### 1. Backend (Golang)

- **Authentication System** using JWT tokens for secure user sessions
- **RESTful API endpoints** for task management operations
- **Database Integration** with structured models for users and tasks
- **Middleware** for authentication and request validation

### 2. Frontend (Next.js + TypeScript)

- **Modern UI** built with Next.js 13 App Router and Tailwind CSS
- **Authentication Flow** with login and registration pages
- **Dashboard Interface** for task management
- **Task Creation Modal** with intuitive user experience

## Current Implementation Status

### ✅ What Works:

- **Project Structure** following clean architecture principles
- **Authentication System** with JWT implementation
- **Basic Task Management** functionality
- **Type-Safe Frontend** with TypeScript integration

### 🚧 Work in Progress:

- AI integration for smart task suggestions
- WebSocket implementation for real-time updates
- Cloud deployment setup
- Task assignment and priority system

## Technical Architecture

### Backend (Golang)

- **Framework**: Gin for REST API development
- **Authentication**: JWT-based session management
- **Database**: Structured models ready for PostgreSQL integration
- **Middleware**: Custom auth middleware for route protection

### Frontend (Next.js)

- **Framework**: Next.js 13 with App Router
- **Styling**: Tailwind CSS for responsive design
- **State Management**: Server actions for data fetching
- **Components**: Modular design with reusable components

## Planned Enhancements

### 1. AI Integration

- OpenAI/Gemini API integration for smart task suggestions
- AI-powered task prioritization
- Intelligent task breakdown assistance

### 2. Real-time Features

- WebSocket implementation for live updates
- Task status synchronization
- Real-time collaboration features

### 3. DevOps

- Docker containerization
- Cloud platform deployment
- CI/CD pipeline setup

## Development Approach

I focused on establishing a solid foundation with:

1. Clean architecture principles
2. Type-safe implementations
3. Scalable project structure
4. Modern development practices

## Challenges Faced

### 1. Architecture Decisions

- Balancing between feature completeness and code quality
- Structuring the project for future AI integration
- Setting up type-safe communication between frontend and backend

### 2. Time Management

- Prioritizing core features within the time constraint
- Setting up the development environment efficiently
- Planning for scalable architecture while maintaining rapid development

## Next Steps

1. **AI Integration**

   - Implement OpenAI/Gemini API endpoints
   - Add AI-powered task suggestions
   - Develop smart task prioritization

2. **Real-time Features**

   - Set up WebSocket connections
   - Implement live task updates
   - Add collaborative features

3. **Deployment**
   - Configure cloud services
   - Set up monitoring and logging
   - Implement CI/CD pipeline

## Technical Learnings

1. **Golang Backend Development**

   - JWT implementation best practices
   - Gin framework optimization
   - Clean architecture patterns

2. **Next.js 13 Features**
   - App Router implementation
   - Server actions utilization
   - TypeScript integration

## Environment Setup

1. Copy `example.env` to `.env`
2. Replace the placeholder values with your actual credentials:
   - `DB_USER`: MongoDB username
   - `DB_PASSWORD`: MongoDB password
   - `MONGODB_URI`: Your MongoDB connection string
   - `JWT_SECRET`: A secure random string for JWT signing
   - `PORT`: The port number for the server (default: 8080)

## Additional Background

I had a lot of fun making this, I'm a swift developer and have worked mostly with Node.js backends.This was a pretty new techstack for me and I loved working on this. I did all of this in 5 hours and I'm thankful to Zocket for giving out such an interesting problem statement.

## Final Note

While the current implementation focuses on core functionality, I've designed the system with scalability in mind. The architecture is ready for AI integration and real-time features, which would be my immediate next steps.

I'm excited about the potential of this project and look forward to developing it further with AI capabilities and real-time collaboration features.

## TL;DR

✅ **Core Backend & Frontend** implemented with authentication  
✅ **Deployed & Live** on Vercel and Railway  
🚧 **AI & Real-time features** planned and architected  
🔜 **Next Steps:** AI integration, WebSockets, and cloud deployment

# Project Structure

```
Zocket
├─ README.md
├─ TaskManagementSystem
│  ├─ backend
│  │  ├─ .env
│  │  ├─ config
│  │  │  └─ database.go
│  │  ├─ go.mod
│  │  ├─ go.sum
│  │  ├─ handlers
│  │  │  ├─ auth.go
│  │  │  ├─ tasks.go
│  │  │  └─ users.go
│  │  ├─ main.go
│  │  ├─ middleware
│  │  │  └─ auth.go
│  │  ├─ models
│  │  │  └─ models.go
│  │  ├─ routes
│  │  │  └─ routes.go
│  │  └─ utils
│  │     └─ jwt.go
│  └─ frontend
│     ├─ eslint.config.mjs
│     ├─ next-env.d.ts
│     ├─ next.config.ts
│     ├─ package-lock.json
│     ├─ package.json
│     ├─ postcss.config.mjs
│     ├─ public
│     ├─ README.md
│     ├─ src
│     │  ├─ app
│     │  │  ├─ auth
│     │  │  │  ├─ login
│     │  │  │  │  └─ page.tsx
│     │  │  │  └─ register
│     │  │  │     └─ page.tsx
│     │  │  ├─ dashboard
│     │  │  │  ├─ layout.tsx
│     │  │  │  ├─ page.tsx
│     │  │  │  └─ tasks
│     │  │  │     ├─ CreateTaskModal.tsx
│     │  │  │     ├─ page.tsx
│     │  │  │     └─ TaskFilters.tsx
│     │  │  └─ globals.css
│     │  ├─ components
│     │  └─ lib
│     │     └─ actions.ts
│     ├─ tailwind.config.ts
│     └─ tsconfig.json
└─ test.go

```
