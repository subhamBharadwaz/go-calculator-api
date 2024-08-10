# Simple Calculator API

A simple calculator API built with Go. This API performs basic arithmetic operations like addition and subtraction. It also includes middleware for logging and request handling.

## Features

- **Basic Arithmetic Operations**: Supports addition and subtraction.
- **Logging Middleware**: Uses `slog` for logging API requests and responses.
- **Handler and Routing**: Custom routing and request handling.
- **CORS**: CORS implementation.
- **Rate Limiter** : Rate limiter middleware to prevent misuse of the API.
- **No Database**: This project does not include a database.

## Getting Started

### Prerequisites

- Go 1.22.4 or later
- `github.com/rs/cors` package for implement CORS
- `golang.org/x/time` package for rate limiting

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/yourusername/calculator-api.git

```

2. Navigate to the project directory:

```bash
cd calculator-api
```

3. Install dependencies:

```bash
go mod tidy
```

### Usage

1. Run the API server:

```bash
go run cmd/calculator-api/main.go
```

2. The API will be available at `http://localhost:8080`
