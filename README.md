# SCS Operator API

Smart City System (SCS) Operator API for managing premises, alarms, incidents, guidance templates, and guards.

## 🚀 Features

- **Premise Management**: Create, update, and manage premises with user assignments
- **Alarm System**: Monitor and manage alarms with status tracking
- **Incident Management**: Handle incidents with guidance assignment and completion tracking
- **Guidance Templates**: Create and manage guidance templates with steps
- **Guard Management**: Manage guard users and their assignments
- **Real-time Processing**: Kafka integration for event streaming
- **API Documentation**: Comprehensive Swagger/OpenAPI documentation
- **Authentication**: JWT-based authentication system
- **Database**: PostgreSQL with GORM ORM
- **Logging**: Structured logging with Zap

## 📋 Prerequisites

Before running this project, make sure you have the following installed:

- **Go 1.23.3** or later
- **PostgreSQL 12+**
- **Apache Kafka** (for event streaming)
- **Git**

## 🛠️ Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd scs-operator
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Copy the example environment file and configure it:

```bash
cp .env .env.local
```

Edit `.env` with your configuration:

```env
# Server Configuration
PORT=1323
MODE=development
READ_TIMEOUT=5s
WRITE_TIMEOUT=5s

# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=smart_city
DB_HOST=127.0.0.1
DB_PORT=5432

# Kafka Configuration
KAFKA_BROKERS=localhost:9093

# Logging Configuration
LOG_LEVEL=debug
```

### 4. Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE smart_city;
CREATE USER your_db_user WITH PASSWORD 'your_db_password';
GRANT ALL PRIVILEGES ON DATABASE smart_city TO your_db_user;
```

The application will automatically run database migrations on startup.

### 5. Kafka Setup

Start Kafka (using Docker):

```bash
# Start Kafka with Docker Compose (if you have docker-compose.yml)
docker-compose up -d kafka

# Or start Kafka manually
# Follow Kafka installation guide for your OS
```

## 🏃‍♂️ Running the Application

### Development Mode

```bash
# Run directly with Go
go run cmd/server/main.go

# Or build and run
go build -o scs-operator cmd/server/main.go
./scs-operator
```

### Production Mode

```bash
# Set environment to production
export ENV=production

# Build optimized binary
go build -ldflags="-w -s" -o scs-operator cmd/server/main.go

# Run the application
./scs-operator
```

### Using Docker

```bash
# Build Docker image
docker build -t scs-operator .

# Run with Docker
docker run -p 1323:1323 --env-file .env scs-operator
```

## 📚 API Documentation (Swagger)

The API documentation is automatically generated using Swagger/OpenAPI and is available at:

**Swagger UI**: `http://localhost:1323/swagger/index.html`

### Swagger Features

- **Interactive API Explorer**: Test endpoints directly from the browser
- **Complete API Reference**: All endpoints, request/response schemas, and examples
- **Authentication Support**: JWT Bearer token authentication
- **Model Definitions**: Detailed schemas for all data models

### Regenerating Swagger Documentation

If you make changes to the API endpoints or models, regenerate the Swagger docs:

```bash
# Install swag CLI tool
go install github.com/swaggo/swag/cmd/swag@latest

# Generate/update Swagger documentation
swag init -g cmd/server/main.go

# The following files will be generated/updated:
# - docs/docs.go
# - docs/swagger.json
# - docs/swagger.yaml
```

## 🔐 Authentication

The API uses JWT Bearer token authentication. Include the token in the Authorization header:

```bash
Authorization: Bearer <your-jwt-token>
```

Most endpoints require authentication except:
- Health check endpoints
- Swagger documentation

## 📖 API Endpoints

### Health Check
- `GET /api/v1/health` - Application health status

### Premises
- `POST /api/v1/premises` - Create a new premise
- `GET /api/v1/premises` - Get paginated list of premises
- `GET /api/v1/premises/{id}` - Get premise by ID
- `PUT /api/v1/premises/{id}` - Update premise
- `POST /api/v1/premises/{id}/assign-users` - Assign users to premise
- `GET /api/v1/premises/{id}/available-users` - Get available users for premise

### Alarms
- `GET /api/v1/alarms` - Get alarms with optional status filtering
- `PATCH /api/v1/alarms/{id}` - Update alarm status

### Incidents
- `POST /api/v1/incidents` - Create a new incident
- `GET /api/v1/incidents` - Get paginated list of incidents
- `GET /api/v1/incidents/{id}` - Get incident by ID
- `PATCH /api/v1/incidents/{id}` - Update incident
- `POST /api/v1/incidents/{id}/assign-guidance` - Assign guidance to incident
- `GET /api/v1/incidents/{id}/guidance` - Get incident guidance
- `PATCH /api/v1/incidents/{id}/complete` - Mark incident as complete

### Guidance Templates
- `POST /api/v1/guidance-templates` - Create guidance template
- `GET /api/v1/guidance-templates` - Get all guidance templates
- `GET /api/v1/guidance-templates/{id}` - Get guidance template by ID
- `PUT /api/v1/guidance-templates/{id}` - Update guidance template

### Guidance Steps
- `POST /api/v1/guidance-steps` - Create guidance step
- `GET /api/v1/guidance-steps` - Get all guidance steps
- `GET /api/v1/guidance-steps/{id}` - Get guidance step by ID

### Guards
- `POST /api/v1/guards` - Create a new guard

## 🏗️ Project Structure

```
scs-operator/
├── cmd/
│   └── server/          # Application entry point
├── config/              # Configuration management
├── docs/                # Swagger documentation (auto-generated)
├── internal/
│   ├── app/            # Application modules (premises, alarms, etc.)
│   ├── container/      # Dependency injection container
│   ├── middlewares/    # HTTP middlewares
│   ├── models/         # Database models
│   ├── server/         # HTTP server setup
│   └── types/          # Custom types and responses
├── pkg/
│   ├── db/             # Database connection
│   ├── errors/         # Error handling
│   ├── kafka/          # Kafka client
│   ├── logger/         # Logging utilities
│   └── validation/     # Input validation
└── test/               # Test files
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/app/premise/...
```

## 🐳 Docker Support

The project includes Docker support for easy deployment:

```bash
# Build image
docker build -t scs-operator .

# Run with environment file
docker run --env-file .env -p 1323:1323 scs-operator
```

## 📝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Ensure PostgreSQL is running
   - Check database credentials in `.env`
   - Verify database exists

2. **Kafka Connection Issues**
   - Ensure Kafka is running on the specified brokers
   - Check `KAFKA_BROKERS` configuration

3. **Swagger Documentation Not Loading**
   - Run `swag init -g cmd/server/main.go` to regenerate docs
   - Ensure the server is running on the correct port

4. **Build Errors**
   - Run `go mod tidy` to clean up dependencies
   - Ensure Go version 1.23.3 or later

For more detailed documentation, check the `docs/` directory.
