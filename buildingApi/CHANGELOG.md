# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2026-01-05

### Refactor

- **Architecture**: Migrated from single-file monolithic structure to a clean, layered architecture.

  - `cmd/api`: Application entry point.
  - `internal/config`: Configuration management.
  - `internal/models`: Domain models.
  - `internal/repository`: Data access layer (Repository Pattern).
  - `internal/handlers`: HTTP handlers (Controllers).
  - `internal/middleware`: HTTP middleware.
  - `internal/routes`: Route definitions.
  - `pkg/response`: Standardized response utilities.

- **Modern Go Features**:

  - **Context**: Implemented `context.Context` for request lifecycle management and cancellation propagation.
  - **Structured Logging**: Replaced standard `log` with `log/slog` for structured, leveled logging (JSON format).
  - **Graceful Shutdown**: Added signal handling to ensure the server shuts down gracefully, finishing active requests before exiting.
  - **Concurrency**: Implemented `sync.RWMutex` in the in-memory repository for thread-safe concurrent access.

- **API Best Practices**:
  - **Dependency Injection**: Handlers and Repositories are injected, improving testability.
  - **Standardized Responses**: Created a `Response` struct and helper functions (`Success`, `Error`, `Created`, etc.) for consistent JSON output.
  - **Input Validation**: Added `Validate()` methods to models to ensure data integrity before processing.
  - **Middleware**:
    - `LoggingMiddleware`: Logs request method, path, status, and duration.
    - `RecoveryMiddleware`: Recovers from panics to prevent server crashes.
    - `CORSMiddleware`: Handles Cross-Origin Resource Sharing.
  - **HTTP Status Codes**: Used appropriate status codes (201 Created, 400 Bad Request, 404 Not Found, 500 Internal Server Error).

### Removed

- `main.go` (root): Removed the old monolithic file.

### Security

- **Panic Recovery**: Middleware ensures the server stays up even if a handler panics.
- **Secure ID Generation**: Using `crypto/rand` for generating secure random IDs instead of `math/rand`.
