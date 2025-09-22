# Swagger Documentation Guide

This guide explains how to use, maintain, and extend the Swagger documentation for the SCS Operator API.

## 📖 Overview

The SCS Operator API uses Swagger/OpenAPI 2.0 for comprehensive API documentation. The documentation is automatically generated from code annotations and provides an interactive interface for testing endpoints.

## 🌐 Accessing Swagger UI

Once the server is running, access the Swagger documentation at:

**URL**: `http://localhost:1323/swagger/index.html`

### Features Available in Swagger UI

- **Interactive API Testing**: Execute API calls directly from the browser
- **Request/Response Examples**: See sample data for all endpoints
- **Authentication Support**: Test authenticated endpoints with JWT tokens
- **Model Schemas**: View detailed data structures and validation rules
- **Error Response Documentation**: Understand error formats and status codes

## 🔐 Authentication in Swagger UI

Most endpoints require JWT authentication. To test authenticated endpoints:

1. **Obtain a JWT Token**: Use your authentication system to get a valid JWT token
2. **Authorize in Swagger UI**:
   - Click the "Authorize" button (🔒) at the top of the Swagger UI
   - Enter: `Bearer <your-jwt-token>`
   - Click "Authorize"
3. **Test Endpoints**: All subsequent requests will include the authorization header

## 📝 Swagger Annotations Reference

### Main API Information (in `cmd/server/main.go`)

```go
// @title SCS Operator API
// @version 1.0
// @description Smart City System (SCS) Operator API for managing premises, alarms, incidents, guidance templates, and guards
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
```

### Endpoint Annotations

```go
// @Summary Brief description of the endpoint
// @Description Detailed description of what the endpoint does
// @Tags tag-name
// @Accept json
// @Produce json
// @Param paramName paramType dataType required "Parameter description"
// @Success 200 {object} ResponseType
// @Failure 400 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /endpoint-path [method]
```

### Parameter Types

- **path**: URL path parameter (`/users/{id}`)
- **query**: Query string parameter (`?page=1&limit=10`)
- **body**: Request body parameter
- **header**: HTTP header parameter
- **formData**: Form data parameter

### Common Annotations Examples

#### GET Endpoint with Query Parameters
```go
// @Summary Get paginated list of premises
// @Description Get a paginated list of all premises with optional filtering
// @Tags premises
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} types.PremiseListResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /premises [get]
```

#### POST Endpoint with Request Body
```go
// @Summary Create a new premise
// @Description Create a new premise with the provided information
// @Tags premises
// @Accept json
// @Produce json
// @Param premise body dto.CreatePremiseDto true "Premise creation data"
// @Success 201 {object} models.Premise
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /premises [post]
```

#### Path Parameter Endpoint
```go
// @Summary Get premise by ID
// @Description Get a specific premise by its ID
// @Tags premises
// @Accept json
// @Produce json
// @Param id path string true "Premise ID"
// @Success 200 {object} models.Premise
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /premises/{id} [get]
```

## 🔄 Regenerating Documentation

### When to Regenerate

Regenerate Swagger documentation when you:
- Add new endpoints
- Modify existing endpoints
- Change request/response structures
- Update API metadata

### How to Regenerate

1. **Install Swag CLI** (if not already installed):
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. **Generate Documentation**:
   ```bash
   swag init -g cmd/server/main.go
   ```

3. **Files Generated**:
   - `docs/docs.go` - Go code for embedding docs
   - `docs/swagger.json` - JSON format documentation
   - `docs/swagger.yaml` - YAML format documentation

### Automation

Consider adding documentation generation to your CI/CD pipeline:

```bash
# In your build script or CI/CD
swag init -g cmd/server/main.go
go build -o scs-operator cmd/server/main.go
```

## 🏗️ Response Type Definitions

### Custom Response Types

For better Swagger documentation, we've created specific response types in `internal/types/swagger_responses.go`:

```go
// PremiseListResponse represents a paginated response for premises
type PremiseListResponse struct {
    Data       []models.Premise `json:"data"`
    Pagination Pagination       `json:"pagination"`
}

// IncidentListResponse represents a paginated response for incidents
type IncidentListResponse struct {
    Data       []models.Incident `json:"data"`
    Pagination Pagination        `json:"pagination"`
}
```

### Why Custom Types?

- **Generic Type Support**: Swagger doesn't handle Go generics well
- **Clear Documentation**: Specific types provide better API documentation
- **Type Safety**: Ensures consistent response structures

## 🎯 Best Practices

### 1. Consistent Tagging
Group related endpoints with consistent tags:
- `premises` - All premise-related endpoints
- `incidents` - All incident-related endpoints
- `alarms` - All alarm-related endpoints
- `guidance-templates` - Guidance template endpoints
- `guidance-steps` - Guidance step endpoints
- `guards` - Guard management endpoints

### 2. Comprehensive Error Documentation
Always document possible error responses:
```go
// @Failure 400 {object} errors.ErrorResponse "Bad Request"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Not Found"
// @Failure 500 {object} errors.ErrorResponse "Internal Server Error"
```

### 3. Security Annotations
Add security annotations to protected endpoints:
```go
// @Security BearerAuth
```

### 4. Parameter Documentation
Provide clear parameter descriptions:
```go
// @Param id path string true "Unique identifier for the resource"
// @Param page query int false "Page number for pagination" default(1)
```

### 5. Model Documentation
Document your models with struct tags:
```go
type CreatePremiseDto struct {
    Name        string `json:"name" validate:"required" example:"Main Office"`
    Address     string `json:"address" validate:"required" example:"123 Main St"`
    Description string `json:"description" example:"Primary office location"`
}
```

## 🐛 Troubleshooting

### Common Issues

1. **Documentation Not Updating**
   - Run `swag init -g cmd/server/main.go` to regenerate
   - Restart the server
   - Clear browser cache

2. **Generic Type Errors**
   - Use specific response types instead of generics
   - Check `internal/types/swagger_responses.go` for examples

3. **Authentication Not Working**
   - Ensure JWT token is valid
   - Check token format: `Bearer <token>`
   - Verify security definitions in main.go

4. **Missing Endpoints**
   - Check if annotations are properly formatted
   - Ensure the handler is registered in routes
   - Verify the file is included in the scan path

### Validation

After regenerating documentation:
1. Check that `docs/swagger.json` is valid JSON
2. Verify all endpoints appear in Swagger UI
3. Test authentication flow
4. Validate request/response examples

## 📚 Additional Resources

- [Swaggo Documentation](https://github.com/swaggo/swag)
- [OpenAPI 2.0 Specification](https://swagger.io/specification/v2/)
- [Swagger UI Documentation](https://swagger.io/tools/swagger-ui/)

## 🔧 Configuration

### Swagger Configuration

The Swagger configuration is defined in `cmd/server/main.go`. To modify:

1. **Change Host/Port**:
   ```go
   // @host your-domain.com:8080
   ```

2. **Update Base Path**:
   ```go
   // @BasePath /api/v2
   ```

3. **Modify API Information**:
   ```go
   // @title Your API Title
   // @version 2.0
   // @description Your API description
   ```

### Environment-Specific Configuration

For different environments, you can modify the generated `docs/docs.go`:

```go
// Development
SwaggerInfo.Host = "localhost:1323"

// Production
SwaggerInfo.Host = "api.yourdomain.com"
```

This completes the comprehensive Swagger documentation guide for the SCS Operator API.
