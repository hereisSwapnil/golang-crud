# Go CRUD API

A simple REST API for managing student records. Built with Go standard library and SQLite.

## What It Does

This API lets you:
- Create a new student
- Get a student by ID
- Get all students
- Update a student
- Delete a student

## Prerequisites

- Go 1.25.1 or higher
- SQLite3 (usually comes with your OS)

## Project Structure

```
golang-crud/
├── cmd/
│   └── golang-crud/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go           # Configuration loading
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers for student operations
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go       # SQLite implementation
│   ├── types/
│   │   └── types.go            # Student data structure
│   └── utils/
│       └── response.go          # HTTP response helpers
├── storage/
│   └── storage.db              # SQLite database file (created on first run)
├── go.mod
└── go.sum
```

## How It Works

1. **Startup**: The application loads configuration from `config/local.yaml`, connects to SQLite, and creates the students table if it doesn't exist.

2. **Architecture**: Uses an interface-based design. The `storage.Storage` interface allows swapping database implementations without changing handlers.

3. **HTTP Server**: Uses Go's standard `net/http` package with `http.ServeMux` for routing.

4. **Request Flow**:
   - Request comes in → Handler validates input → Storage layer executes database operation → Response sent back

5. **Shutdown**: Handles graceful shutdown on SIGINT or SIGTERM signals with a 5-second timeout.

## API Endpoints

### Create Student
- **POST** `/api/v1/student`
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "age": 20,
    "email": "john@example.com"
  }
  ```
- **Success Response** (200):
  ```json
  {
    "status": 200,
    "data": {
      "id": 1
    }
  }
  ```

### Get Student by ID
- **GET** `/api/v1/student/{id}`
- **Success Response** (200):
  ```json
  {
    "status": 200,
    "data": {
      "id": 1,
      "name": "John Doe",
      "age": 20,
      "email": "john@example.com"
    }
  }
  ```

### Get All Students
- **GET** `/api/v1/students`
- **Success Response** (200):
  ```json
  {
    "status": 200,
    "data": [
      {
        "id": 1,
        "name": "John Doe",
        "age": 20,
        "email": "john@example.com"
      }
    ]
  }
  ```

### Update Student
- **PUT** `/api/v1/student/{id}`
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
    "age": 21,
    "email": "jane@example.com"
  }
  ```
- **Success Response** (200):
  ```json
  {
    "status": 200,
    "data": {
      "id": 1
    }
  }
  ```

### Delete Student
- **DELETE** `/api/v1/student/{id}`
- **Success Response** (200):
  ```json
  {
    "status": 200,
    "data": {
      "id": 1
    }
  }
  ```

## Response Formats

### Success Response
```json
{
  "status": 200,
  "data": { ... }
}
```

### Error Response
```json
{
  "status": 400,
  "message": "Failed to decode student data"
}
```

### Validation Error Response
```json
{
  "status": 400,
  "errors": [
    "Field name is required",
    "Field email is not valid"
  ]
}
```

## Testing the API

Example curl commands:

**Create a student:**
```bash
curl -X POST http://localhost:8082/api/v1/student \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 20, "email": "john@example.com"}'
```

**Get a student by ID:**
```bash
curl http://localhost:8082/api/v1/student/1
```

**Get all students:**
```bash
curl http://localhost:8082/api/v1/students
```

**Update a student:**
```bash
curl -X PUT http://localhost:8082/api/v1/student/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "age": 21, "email": "jane@example.com"}'
```

**Delete a student:**
```bash
curl -X DELETE http://localhost:8082/api/v1/student/1
```

## Packages Used

- `github.com/mattn/go-sqlite3` - SQLite database driver
- `github.com/ilyakaznacheev/cleanenv` - Loads configuration from YAML files
- `github.com/go-playground/validator/v10` - Validates request data
- Standard library: `net/http`, `database/sql`, `encoding/json`, `context`, `os/signal`

## Configuration

Edit `config/local.yaml`:

```yaml
env: "development"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

## Running the Application

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/golang-crud/main.go
   ```

3. The server will start on `localhost:8082` (or the address in your config).