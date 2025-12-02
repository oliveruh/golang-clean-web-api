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
- ✅ Ready for adding custom endpoints

## Project Structure

```
src/
├── api/
│   ├── dto/              # Data Transfer Objects
│   ├── handler/          # HTTP handlers
│   ├── helper/           # Helper functions
│   ├── middleware/       # Middleware (CORS, etc.)
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
├── pkg/
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
```

### Environment Variables

- `APP_ENV` - Set to "docker" or "production" to use different configs
- `PORT` - Override the configured port

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

### Countries
```bash
# Get all countries with pagination
curl -X POST http://localhost:8080/api/v1/countries/get-by-filter \
  -H "Content-Type: application/json" \
  -d '{"pageNumber": 1, "pageSize": 10}'

# Get country by ID
curl http://localhost:8080/api/v1/countries/1

# Create country
curl -X POST http://localhost:8080/api/v1/countries \
  -H "Content-Type: application/json" \
  -d '{"name": "Brazil"}'

# Update country
curl -X PUT http://localhost:8080/api/v1/countries/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Name"}'

# Delete country
curl -X DELETE http://localhost:8080/api/v1/countries/1
```

Similar endpoints are available for:
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

## Next Steps

This template provides a solid foundation. You can extend it with:

- Authentication & authorization (JWT)
- API documentation (Swagger)
- More comprehensive logging
- Rate limiting
- Request validation
- Unit and integration tests
- CI/CD pipelines

## License

MIT License - see LICENSE file for details
