# Demo Go API

A simple REST API for managing users and products.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /users | List all users (supports `?is_active=true/false`) |
| POST | /users | Create a new user |
| GET | /users/:id | Get user by ID |
| GET | /products | List all products |

## Running

```bash
go run main.go
```
