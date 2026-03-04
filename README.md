# rat-api

A simple Go API service that provides a logging endpoint using [lograt](https://github.com/lwl27/lograt).

## Overview

rat-api exposes an HTTP endpoint that accepts log messages and forwards them to the lograt logging library. It supports multiple log levels and flexible input methods via query parameters or JSON body.

## Endpoints

### POST /log

Log a message with optional level and custom fields.

**Query Parameters:**

| Parameter | Type   | Required | Description              |
|-----------|--------|----------|--------------------------|
| message   | string | Yes      | The log message         |
| level     | string | No       | Log level (debug, info, warn, error). Default: info |
| *         | any    | No       | Any additional fields to include in the log |

**Request Body (JSON):**

```json
{
  "message": "User logged in",
  "level": "info",
  "user_id": 123,
  "action": "login"
}
```

**Response:**

```json
{
  "status": "ok",
  "message": "User logged in"
}
```

## Examples

### Using query parameters

```bash
curl "http://localhost:8080/log?message=Hello%20World&level=info"
```

### Using JSON body

```bash
curl -X POST "http://localhost:8080/log?message=Test" \
  -H "Content-Type: application/json" \
  -d '{"user": "john", "action": "login"}'
```

### Log levels

```bash
curl "http://localhost:8080/log?message=Debug%20message&level=debug"
curl "http://localhost:8080/log?message=Warning%20message&level=warn"
curl "http://localhost:8080/log?message=Error%20message&level=error"
```

## Getting Started

### Prerequisites

- Go 1.25+

### Build

```bash
go build -o bin/rat-api
```

### Run

```bash
go run main.go
```

The server will start on port 8080.

### Test

```bash
go test -v ./...
```

## Project Structure

```
rat-api/
├── main.go          # Application entry point and handlers
├── main_test.go     # Unit tests
├── go.mod           # Go module dependencies
└── README.md        # This file
```
