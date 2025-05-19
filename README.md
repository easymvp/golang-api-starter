# EasyMVP API

A Go-based RESTful API service built with Gin framework, GORM, and JWT authentication. This project provides a solid foundation for building scalable and maintainable backend services.

## Features

- RESTful API with Gin framework
- JWT authentication
- PostgreSQL and SQLite database support with GORM
- Dependency injection with Uber FX
- Structured logging with Zap
- API documentation with Swagger
- Environment variable configuration
- CORS support
- Health check endpoint

## Requirements

- Go 1.23.0 or higher
- PostgreSQL or SQLite
- Docker (optional, for containerized deployment)

## Installation

### Local Development

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd easymvp_api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:
   ```
   DB_URL=postgres://username:password@localhost:5432/dbname
   DB_DRIVER=postgres
   DB_MAX_OPEN_CONNS=10
   DB_MAX_IDLE_CONNS=5
   DB_CONN_MAX_LIFE=10
   DB_AUTO_MIGRATION=true
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Docker Deployment

1. Build the Docker image:
   ```bash
   docker build -t easymvp_api .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 --env-file .env easymvp_api
   ```

## Project Structure

```
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── app/                    # Application core
│   │   ├── app.go              # Main application setup
│   │   ├── context.go          # Context management
│   │   ├── env.go              # Environment configuration
│   │   ├── error_response.go   # Error handling
│   │   ├── handlers/           # HTTP handlers
│   │   └── ...
│   ├── database/               # Database configuration and models
│   │   ├── db_config.go        # Database configuration
│   │   ├── gorm.go             # GORM setup
│   │   └── ...
│   ├── log/                    # Logging configuration
│   ├── swagger/                # API documentation
│   ├── tests/                  # Test utilities
│   ├── users/                  # User management
│   └── utils/                  # Utility functions
├── cicd/                       # CI/CD configuration
│   └── db-migration/           # Database migration scripts
├── Dockerfile                  # Docker configuration
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksums
```

## API Endpoints

### Public Endpoints

- `GET /health-check` - Health check endpoint

### Protected Endpoints (Require JWT Authentication)

- Protected endpoints require a valid JWT token in the Authorization header

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_URL | Database connection URL | - |
| DB_DRIVER | Database driver (postgres, sqlite) | - |
| DB_MAX_OPEN_CONNS | Maximum number of open connections | 10 |
| DB_MAX_IDLE_CONNS | Maximum number of idle connections | 10 |
| DB_CONN_MAX_LIFE | Connection maximum lifetime (seconds) | 10 |
| DB_AUTO_MIGRATION | Enable automatic database migrations | true |

## Key Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Uber FX](https://github.com/uber-go/fx) - Dependency injection
- [Zap](https://github.com/uber-go/zap) - Structured logging
- [Gin-JWT](https://github.com/appleboy/gin-jwt) - JWT middleware for Gin
- [Gin-Swagger](https://github.com/swaggo/gin-swagger) - Swagger documentation for Gin
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading

## Database Support

The application supports both PostgreSQL and SQLite databases through GORM. Configure the database connection using the `DB_URL` and `DB_DRIVER` environment variables.

### PostgreSQL Configuration Example

```
DB_URL=postgres://username:password@localhost:5432/dbname
DB_DRIVER=postgres
```

### SQLite Configuration Example

```
DB_URL=file:test.db
DB_DRIVER=sqlite
```

## Testing

Run tests with:

```bash
go test ./...
```

## License

[MIT License](LICENSE)
