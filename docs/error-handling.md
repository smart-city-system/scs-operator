# Error Handling System

This document describes the comprehensive error handling system implemented in the smart-city application.

## Overview

The error handling system provides:
- **Standardized error responses** across all API endpoints
- **Proper HTTP status codes** based on error types
- **Detailed validation errors** with field-level information
- **Centralized error logging** and monitoring
- **Consistent error structure** for frontend consumption

## Architecture

### 1. Custom Error Types (`pkg/errors/errors.go`)

The system defines several error types:

```go
// Client errors (4xx)
ErrorTypeValidation    = "VALIDATION_ERROR"      // 400
ErrorTypeNotFound      = "NOT_FOUND"             // 404
ErrorTypeUnauthorized  = "UNAUTHORIZED"          // 401
ErrorTypeForbidden     = "FORBIDDEN"             // 403
ErrorTypeBadRequest    = "BAD_REQUEST"           // 400
ErrorTypeConflict      = "CONFLICT"              // 409

// Server errors (5xx)
ErrorTypeInternal      = "INTERNAL_ERROR"        // 500
ErrorTypeDatabase      = "DATABASE_ERROR"        // 500
ErrorTypeExternal      = "EXTERNAL_SERVICE_ERROR" // 500
ErrorTypeTimeout       = "TIMEOUT_ERROR"         // 500
```

### 2. Error Response Structure

All errors return a consistent JSON structure:

```json
{
  "error": {
    "type": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "email",
        "message": "email must be a valid email address",
        "value": "invalid-email"
      }
    ]
  },
  "request_id": "req-123456",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

### 3. Validation System (`pkg/validation/validator.go`)

Automatic validation using struct tags:

```go
type CreateUserDto struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email,max=255"`
    Password string `json:"password" validate:"required,min=6,max=100"`
    Role     string `json:"role" validate:"required,role"`
}
```

## Usage Examples

### In Controllers

```go
func (h *Handler) CreateUser() echo.HandlerFunc {
    return func(c echo.Context) error {
        createUserDto := &dto.CreateUserDto{}
        if err := c.Bind(createUserDto); err != nil {
            return errors.NewBadRequestError("Invalid request body")
        }
        
        // Validate the DTO
        if err := validation.ValidateStruct(createUserDto); err != nil {
            return err // Returns validation error with details
        }
        
        user, err := h.svc.CreateUser(c.Request().Context(), createUserDto)
        if err != nil {
            return err // Service errors are handled by middleware
        }
        
        return c.JSON(201, user)
    }
}
```

### In Services

```go
func (s *Service) CreateUser(ctx context.Context, dto *dto.CreateUserDto) (*models.User, error) {
    // Business logic validation
    if dto.Email == "" {
        return nil, errors.NewValidationError("Email is required", nil)
    }
    
    user, err := s.userRepo.CreateUser(ctx, user)
    if err != nil {
        if isDuplicateEmailError(err) {
            return nil, errors.NewConflictError("User with this email already exists")
        }
        return nil, errors.NewDatabaseError("create user", err)
    }
    
    return user, nil
}
```

### In Repositories

```go
func (r *Repository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.NewNotFoundError("user")
        }
        return nil, errors.NewDatabaseError("get user by id", err)
    }
    return &user, nil
}
```

## Error Types and When to Use Them

### Validation Errors (400)
- Invalid input format
- Missing required fields
- Field length violations
- Invalid enum values

```go
errors.NewValidationError("Invalid input", validationDetails)
```

### Not Found Errors (404)
- Resource doesn't exist
- Invalid ID provided

```go
errors.NewNotFoundError("user")
```

### Conflict Errors (409)
- Duplicate resources
- Business rule violations

```go
errors.NewConflictError("Email already exists")
```

### Database Errors (500)
- Database connection issues
- Query execution failures
- Transaction failures

```go
errors.NewDatabaseError("operation name", err)
```

### Internal Errors (500)
- Unexpected application errors
- Third-party service failures

```go
errors.NewInternalError("Unexpected error occurred", err)
```

## Middleware

The error handling middleware automatically:
- Catches all errors from handlers
- Converts them to appropriate HTTP responses
- Logs errors with request context
- Handles GORM-specific errors
- Provides consistent error formatting

## Testing

Run error handling tests:

```bash
go test ./pkg/errors/...
```

## Best Practices

1. **Always use custom error types** instead of returning raw errors
2. **Provide meaningful error messages** that help users understand the issue
3. **Include relevant details** in validation errors
4. **Log errors at the appropriate level** (validation = info, database = error)
5. **Don't expose internal details** in error messages sent to clients
6. **Use specific error types** rather than generic ones when possible
