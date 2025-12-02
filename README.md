# Golang Clean Web API - Minimal Template

A minimal, clean architecture web API template built with Go and Gin framework. This is a basic starting point for building REST APIs without unnecessary dependencies or complexity.

## Features

- ✅ Clean and simple project structure
- ✅ Gin web framework for routing
- ✅ CORS middleware support
- ✅ Configuration management with Viper
- ✅ Health check endpoint
- ✅ Ready for adding custom endpoints

## Project Structure

```
src/
├── api/
│   ├── middleware/
│   │   └── cors.go          # CORS middleware
│   ├── api.go               # Server initialization and routing
│   └── handler_simple.go    # Basic health check handler
├── cmd/
│   └── main.go              # Application entry point
└── config/
    ├── config.go            # Configuration management
    ├── config-development.yml
    └── config-production.yml
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [Viper](https://github.com/spf13/viper) - Configuration management

## Quick Start

### Prerequisites

- Go 1.22 or higher

### Installation & Running

1. Clone the repository:
```bash
git clone https://github.com/oliveruh/golang-clean-web-api
cd golang-clean-web-api
```

2. Build and run:
```bash
cd src
go build -o ../bin/server ./cmd/main.go
cd cmd
../../bin/server
```

Or run directly with Go:
```bash
cd src/cmd
go run main.go
```

The server will start on port 8080 by default.

### Testing the API

Test the health check endpoint:
```bash
curl http://localhost:8080/api/v1/health
```

Response:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

## Configuration

Edit `src/config/config-development.yml` to customize settings:

```yaml
server:
  port: "8080"
  runMode: "debug"  # or "release" for production
cors:
  allowOrigins: "*"
```

### Environment Variables

- `PORT` - Override the configured port
- `APP_ENV` - Set to "production" to use production config

## Adding New Endpoints

To add a new endpoint:

1. Create a handler function in `src/api/handler_simple.go` or a new file:
```go
func YourHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Your response",
    })
}
```

2. Register the route in `src/api/api.go`:
```go
func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    v1 := api.Group("/v1")
    {
        v1.GET("/health", Health)
        v1.GET("/your-endpoint", YourHandler)  // Add your route here
    }
}
```

## Building for Production

```bash
cd src
go build -o ../bin/server ./cmd/main.go
```

Set the environment:
```bash
export APP_ENV=production
./bin/server
```

## Next Steps

This template provides a minimal foundation. You can extend it with:

- Database integration (PostgreSQL, MongoDB, etc.)
- Authentication & authorization (JWT)
- Request validation
- Logging
- API documentation (Swagger)
- Testing infrastructure
- Docker containerization

## License

MIT License - see LICENSE file for details
