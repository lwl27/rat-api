# AGENTS.md - Guidelines for Agentic Coding

This document provides guidance for agents working in this repository.

## Project Overview

- **Project Name**: rat-api
- **Type**: Simple API service in Go
- **Purpose**: API that uses lograt for logging

## Build/Lint/Test Commands

### Building
```bash
go build ./...        # Build all packages
go build -o bin/      # Build to bin directory
```

### Running
```bash
go run main.go        # Run the main application
go run .             # Run in module mode (Go 1.17+)
```

### Testing
```bash
go test ./...                    # Run all tests
go test -v ./...                 # Run with verbose output
go test -run TestName ./...      # Run single test by name
go test -v -run TestName ./...   # Run single test with verbose
go test -cover ./...            # Run with coverage
go test -race ./...              # Run with race detector
```

### Linting
```bash
go vet ./...                     # Run go vet
golangci-lint run                # Run golangci-lint (if installed)
staticcheck ./...                # Run staticcheck (if installed)
```

### Code Generation
```bash
go generate ./...                # Run generate directives
```

### Dependencies
```bash
go mod download                  # Download dependencies
go mod tidy                      # Clean up go.mod/go.sum
```

## Code Style Guidelines

### General Principles
- Write clean, readable, and idiomatic Go code
- Follow the official Go code review comments: https://github.com/golang/go/wiki/CodeReviewComments
- Keep functions small and focused (single responsibility)
- Early returns are preferred over deeply nested conditionals

### Naming Conventions
- **Files**: Use snake_case (e.g., `user_service.go`, `api_handler.go`)
- **Types/Interfaces**: Use PascalCase (e.g., `UserService`, `APIHandler`)
- **Functions/Variables**: Use camelCase (e.g., `getUser`, `handleRequest`)
- **Constants**: Use PascalCase for exported, camelCase for unexported
- **Packages**: Use short, lowercase names (e.g., `handler`, `service`, `repo`)
- Avoid underscores in names unless necessary for test files (`_test.go`)

### Imports
- Group imports in this order:
  1. Standard library
  2. Third-party packages
  3. Internal packages
- Use blank line between groups
- Example:
  ```go
  import (
      "context"
      "encoding/json"
      "time"

      "github.com/gin-gonic/gin"
      "github.com/go-playground/validator/v10"

      "rat-api/internal/handler"
      "rat-api/internal/service"
  )
  ```
- Use import aliases only when necessary (e.g., `log "github.com/sirupsen/logrus"`)

### Formatting
- Run `go fmt ./...` before committing
- Use goimports for automatic import management
- Maximum line length: 100 characters (soft limit)
- Leave a blank line between top-level declarations

### Types
- Use primitive types where possible (`string`, `int`, `bool`, etc.)
- Use custom types for domain-specific concepts
- Define errors as sentinel errors or custom error types
- Use `context.Context` as first parameter for functions that may timeout/cancel

### Error Handling
- Always handle errors explicitly - never ignore with `_`
- Return meaningful error messages (lowercase, no punctuation)
- Wrap errors with `fmt.Errorf("context: %w", err)` for stack traces
- Use custom error types for domain-specific errors
- Check errors early and return immediately

### Logging
- Use lograt for logging (as per project requirement)
- Log levels: Debug, Info, Warn, Error
- Include contextual information in log fields
- Avoid logging sensitive information (passwords, tokens, etc.)

### API Design
- Follow RESTful conventions
- Request/Response structs should be clearly named (`CreateUserRequest`, `UserResponse`)
- Use HTTP status codes correctly (200, 201, 400, 404, 500, etc.)
- Return consistent error response format

### Testing
- Test files should be named `*_test.go` in the same package
- Use table-driven tests when testing multiple cases
- Use descriptive test names: `TestFunctionName_ShouldDoSomething_WhenCondition`
- Follow AAA pattern: Arrange, Act, Assert
- Mock external dependencies

### Configuration
- Use environment variables or config files
- Never hardcode sensitive values
- Use `.env` files for local development (already gitignored)

### Git Conventions
- Commit messages: Present tense, imperative mood ("Add user service" not "Added user service")
- Branch naming: `feature/description`, `fix/description`, `refactor/description`
- PR titles should match commit message style

## Project Structure (Recommended)

```
rat-api/
├── cmd/
│   └── api/
│       └── main.go           # Entry point
├── internal/
│   ├── handler/               # HTTP handlers
│   ├── service/               # Business logic
│   ├── repository/            # Data access
│   ├── model/                 # Data models
│   └── middleware/           # HTTP middleware
├── pkg/                       # Reusable packages
├── configs/                   # Configuration files
├── go.mod
├── go.sum
└── README.md
```

## Common Tasks

### Adding a New Endpoint
1. Define request/response types in `internal/model/`
2. Add handler function in `internal/handler/`
3. Register route in `cmd/api/main.go`
4. Add tests in `internal/handler/*_test.go`

### Running a Specific Test
```bash
go test -v -run TestHandler_GetUser ./internal/handler/
```

### Running All Tests with Coverage
```bash
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```
