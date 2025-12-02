# Golang Clean Web API - Template with Essential Features

A clean architecture web API template built with Go, Gin, PostgreSQL, and Redis. This template includes essential infrastructure while remaining simple and extendable.

## Features

- ✅ Clean and organized project structure
- ✅ Gin web framework for routing
- ✅ PostgreSQL database with GORM ORM
- ✅ Redis caching support
- ✅ Basic logging with Zap
- ✅ Docker and Docker Compose setup
- ✅ CORS middleware support
- ✅ Configuration management with Viper
- ✅ Health check endpoint
- ✅ Sample CRUD endpoints (Countries, Cities, Colors)
- ✅ Database migrations
- ✅ **JWT Authentication & Authorization**
- ✅ **API Documentation with Swagger/OpenAPI**
- ✅ **Rate Limiting middleware**
- ✅ Ready for adding custom endpoints

## Project Structure

```
src/
├── api/
│   ├── dto/              # Data Transfer Objects
│   ├── handler/          # HTTP handlers (including auth)
│   ├── helper/           # Helper functions
│   ├── middleware/       # Middleware (CORS, Auth, Rate Limiting)
│   ├── router/           # Route definitions
│   └── validation/       # Custom validators
├── cmd/
│   └── main.go          # Application entry point
├── common/              # Common utilities
├── config/              # Configuration management
├── constant/            # Application constants
├── dependency/          # Dependency injection
├── domain/
│   ├── filter/          # Query filters
│   ├── model/           # Database models
│   └── repository/      # Repository interfaces
├── infra/
│   ├── cache/           # Redis cache implementation
│   └── persistence/     # Database & repository implementation
├── docs/                # Swagger documentation (auto-generated)
├── pkg/
│   ├── jwt/             # JWT token service
│   ├── logging/         # Logging utilities
│   ├── metrics/         # Metrics (Prometheus)
│   └── service_errors/  # Error handling
└── usecase/            # Business logic layer
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Redis](https://redis.io/) - Caching
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Zap](https://github.com/uber-go/zap) - Structured logging
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token authentication
- [Swagger](https://github.com/swaggo/swag) - API documentation
- [Tollbooth](https://github.com/didip/tollbooth) - Rate limiting

## Quick Start

### Prerequisites

- Go 1.22 or higher
- Docker and Docker Compose (for running with containers)

### Option 1: Run with Docker Compose (Recommended)

This will start the API, PostgreSQL, and Redis:

```bash
docker-compose up -d
```

The API will be available at `http://localhost:8080`

To stop:
```bash
docker-compose down
```

### Option 2: Run Locally

1. Start PostgreSQL and Redis:
```bash
docker-compose up -d postgres redis
```

2. Build and run the application:
```bash
cd src
go build -o ../bin/server ./cmd/main.go
cd cmd
../../bin/server
```

The server will start on port 8080 by default.

## Configuration

Configuration files are located in `src/config/`:
- `config-development.yml` - For local development
- `config-docker.yml` - For Docker environment
- `config-production.yml` - For production

Example configuration:
```yaml
server:
  port: "8080"
  runMode: "debug"
logger:
  filePath: "../logs/"
  encoding: "json"
  level: "debug"
  logger: "zap"
cors:
  allowOrigins: "*"
postgres:
  host: "localhost"
  port: "5432"
  user: "postgres"
  password: "admin"
  dbName: "web_api_db"
  sslMode: "disable"
redis:
  host: "localhost"
  port: "6379"
  password: "password"
jwt:
  secret: "your-256-bit-secret-key"
  accessExpireTime: 60  # minutes
  refreshExpireTime: 10080  # minutes (7 days)
rateLimiter:
  enabled: true
  requestsPerMin: 100
```

### Environment Variables

- `APP_ENV` - Set to "docker" or "production" to use different configs
- `PORT` - Override the configured port

## API Documentation

Interactive API documentation is available via Swagger UI when running in debug/development mode:

```
http://localhost:8080/swagger/index.html
```

The Swagger documentation provides:
- Complete API endpoint reference
- Request/response schemas
- Try-it-out functionality for testing endpoints
- Authentication support for protected endpoints

## Authentication

The API uses JWT (JSON Web Token) for authentication. Protected endpoints require a valid JWT token in the Authorization header.

### Register a New User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "testuser@example.com"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

Response:
```json
{
  "result": {
    "access_token": "eyJhbGc...",
    "refresh_token": "eyJhbGc..."
  },
  "success": true,
  "resultCode": 0
}
```

### Refresh Token
```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGc..."
  }'
```

### Using Protected Endpoints

All CRUD endpoints (Countries, Cities, Colors) require authentication. Include the access token in the Authorization header:

```bash
# Get all countries with authentication
curl -X POST http://localhost:8080/api/v1/countries/get-by-filter \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGc..." \
  -d '{"pageNumber": 1, "pageSize": 10}'

# Create country with authentication
curl -X POST http://localhost:8080/api/v1/countries \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGc..." \
  -d '{"name": "Brazil"}'
```

## Rate Limiting

Rate limiting is enabled by default to protect the API from abuse. The default configuration allows:
- 100 requests per minute in development
- 60 requests per minute in production

Rate limiting can be configured in the config files:
```yaml
rateLimiter:
  enabled: true
  requestsPerMin: 100
```

When rate limit is exceeded, the API returns a 429 (Too Many Requests) status code.

## API Endpoints

### Health Check (Public)
```bash
curl http://localhost:8080/api/v1/health
```

### Countries (Protected - requires authentication)
```bash
# Note: All examples below require "Authorization: Bearer <token>" header

# Get all countries with pagination
curl -X POST http://localhost:8080/api/v1/countries/get-by-filter \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"pageNumber": 1, "pageSize": 10}'

# Get country by ID
curl http://localhost:8080/api/v1/countries/1 \
  -H "Authorization: Bearer <token>"

# Create country
curl -X POST http://localhost:8080/api/v1/countries \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"name": "Brazil"}'

# Update country
curl -X PUT http://localhost:8080/api/v1/countries/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"name": "Updated Name"}'

# Delete country
curl -X DELETE http://localhost:8080/api/v1/countries/1 \
  -H "Authorization: Bearer <token>"
```

Similar protected endpoints are available for:
- `/api/v1/cities` - City management
- `/api/v1/colors` - Color management

## Adding New Endpoints

1. **Create Model** in `src/domain/model/`:
```go
type YourModel struct {
    BaseModel
    Name string `gorm:"size:50;not null"`
}
```

2. **Create DTO** in `src/api/dto/`:
```go
type CreateYourModelRequest struct {
    Name string `json:"name" binding:"required"`
}
```

3. **Create Handler** in `src/api/handler/`:
```go
func NewYourModelHandler(cfg *config.Config) *GenericCrud[models.YourModel, dto.CreateYourModelRequest, dto.UpdateYourModelRequest, dto.YourModelResponse] {
    return NewGenericCrud[models.YourModel, dto.CreateYourModelRequest, dto.UpdateYourModelRequest, dto.YourModelResponse](cfg, dependency.GetYourModelRepository(cfg))
}
```

4. **Register Router** in `src/api/router/` and `src/api/api.go`

## Database Migrations

Migrations run automatically on startup. Initial data includes:
- Sample countries with cities
- Sample colors

To modify migrations, edit `src/infra/persistence/migration/1_Init.go`

## Development

### Running Tests
```bash
cd src
go test ./...
```

### Building
```bash
cd src
go build -o ../bin/server ./cmd/main.go
```

### Linting
```bash
cd src
go fmt ./...
go vet ./...
```

### Regenerating Swagger Documentation

After making changes to API endpoints or handlers, regenerate the Swagger documentation:

```bash
cd src
swag init -g cmd/main.go -o ./docs
```

## Security Considerations

1. **JWT Secret**: Change the default JWT secret in production environments. Use a strong, random 256-bit key.
2. **HTTPS**: Always use HTTPS in production to protect JWT tokens in transit.
3. **Password Security**: Passwords are hashed using bcrypt before storage.
4. **Rate Limiting**: Adjust rate limits based on your application's needs and infrastructure.
5. **CORS**: Configure CORS settings appropriately for your frontend domains.

## Next Steps

This template now includes authentication, API documentation, and rate limiting. You can further extend it with:

- Role-based access control (RBAC)
- More comprehensive logging and monitoring
- Advanced rate limiting strategies (per-user, per-endpoint)
- Unit and integration tests
- CI/CD pipelines
- Database backup and recovery strategies
- Distributed tracing
- API versioning strategies

## License

MIT License - see LICENSE file for details
