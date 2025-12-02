# Adding New Endpoints - Example Guide

This guide shows how to add new endpoints to the minimal template.

## Example: Adding a "Hello" Endpoint

### Step 1: Add Handler Function

In `src/api/handler_simple.go`, add:

```go
// Hello godoc
// @Summary Say Hello
// @Description Returns a personalized greeting
// @Tags example
// @Accept json
// @Produce json
// @Param name query string false "Name to greet"
// @Success 200 {object} map[string]interface{}
// @Router /v1/hello [get]
func Hello(c *gin.Context) {
    name := c.DefaultQuery("name", "World")
    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Hello, %s!", name),
    })
}
```

### Step 2: Register Route

In `src/api/api.go`, update the `RegisterRoutes` function:

```go
func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    v1 := api.Group("/v1")
    {
        v1.GET("/health", Health)
        v1.GET("/hello", Hello)  // Add this line
    }
}
```

### Step 3: Test It

```bash
# Build and run
cd src
go build -o ../bin/server ./cmd/main.go
cd cmd
../../bin/server

# In another terminal, test the endpoint
curl "http://localhost:8080/api/v1/hello"
curl "http://localhost:8080/api/v1/hello?name=Alice"
```

Expected responses:
```json
{"message":"Hello, World!"}
{"message":"Hello, Alice!"}
```

## Example: Adding a POST Endpoint with JSON Body

### Step 1: Define Request/Response Structures

Create a new file `src/api/types.go`:

```go
package api

type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

type CreateUserResponse struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### Step 2: Add Handler

In `src/api/handler_simple.go`:

```go
func CreateUser(c *gin.Context) {
    var req CreateUserRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    // In a real app, you would save to database here
    response := CreateUserResponse{
        ID:    "user-123",
        Name:  req.Name,
        Email: req.Email,
    }
    
    c.JSON(http.StatusCreated, response)
}
```

### Step 3: Register Route

```go
func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    v1 := api.Group("/v1")
    {
        v1.GET("/health", Health)
        v1.POST("/users", CreateUser)  // Add this line
    }
}
```

### Step 4: Test It

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

## Organizing Code for Larger Projects

As your project grows, consider this structure:

```
src/api/
├── handlers/
│   ├── health.go
│   ├── user.go
│   └── product.go
├── models/
│   ├── user.go
│   └── product.go
├── middleware/
│   └── cors.go
└── api.go
```

Then import and use them in `api.go`:

```go
import (
    "github.com/naeemaei/golang-clean-web-api/api/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    v1 := api.Group("/v1")
    {
        v1.GET("/health", handlers.Health)
        
        // User routes
        users := v1.Group("/users")
        {
            users.GET("", handlers.GetUsers)
            users.POST("", handlers.CreateUser)
            users.GET("/:id", handlers.GetUser)
        }
    }
}
```

## Next Steps

- Add database integration (see PostgreSQL, MongoDB examples)
- Add authentication middleware
- Add request validation
- Add error handling middleware
- Add logging

## Testing

The template includes a basic test example in `src/api/api_test.go`.

### Running Tests

```bash
cd src
go test ./... -v
```

### Adding Tests for Your Endpoints

Create a test file alongside your handler file (e.g., `user_handler_test.go` for `user_handler.go`):

```go
package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestYourEndpoint(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.New()
    RegisterRoutes(r)
    
    req, _ := http.NewRequest("GET", "/api/v1/your-endpoint", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("Expected 200, got %d", w.Code)
    }
}
```

### Test Coverage

```bash
# Run tests with coverage
go test ./... -cover

# Generate HTML coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```
