# Backend — Go (Gin) + MongoDB + JWT

A clean REST API implementing onboarding requirements:
- User and Event models
- Roles: Event Creators and Attendees
- JWT Authentication (login/register)
- Example protected routes

## Quick Start

1) Copy env and set values
```
cp .env.example .env
```
Set `MONGO_URI`, `DB_NAME`, and `JWT_SECRET`.

2) Run (Go 1.21+)
```
go run .
```
Server runs on `http://localhost:${PORT:-8080}`.

## API (Summary)

- POST /api/auth/register {name,email,password,role}
- POST /api/auth/login {email,password} → {token}
- GET  /api/events (public)
- POST /api/events (role: creator)
- GET  /api/events/:id (auth)
- PATCH /api/events/:id (role: creator)
- DELETE /api/events/:id (role: creator)

Auth: set `Authorization: Bearer <token>` header for protected routes.

## Structure
```
backend/
├─ main.go
├─ config/db.go
├─ controllers/
│  ├─ auth_controller.go
│  └─ event_controller.go
├─ middleware/auth.go
├─ models/
│  ├─ user.go
│  └─ event.go
├─ routes/routes.go
└─ utils/
   ├─ jwt.go
   └─ password.go
```

## Notes
- Passwords are hashed with bcrypt
- Minimal validation; extend as needed
- Add rate limiting and CORS per deployment needs
