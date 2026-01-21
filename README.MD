# Missing Person Service

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)]()

RESTful API service untuk mengelola data orang hilang. Service ini menyediakan endpoints untuk registrasi, pencarian, dan tracking status kasus orang hilang.

## 📚 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Architecture](#-architecture)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Running the Service](#-running-the-service)
- [API Documentation](#-api-documentation)
- [Database Schema](#-database-schema)
- [Testing](#-testing)
- [Project Structure](#-project-structure)
- [Development](#-development)
- [Deployment](#-deployment)
- [Contributing](#-contributing)
- [License](#-license)

## 🔍 Overview

Missing Person Service adalah aplikasi backend yang dibangun dengan Go untuk membantu dalam pengelolaan dan tracking kasus orang hilang. Service ini dirancang dengan clean architecture principles untuk memudahkan maintenance dan scalability.

### Use Cases

- **Law Enforcement**: Tracking dan managing missing person cases
- **NGOs**: Coordinating search efforts untuk orang hilang
- **Public**: Melaporkan dan mencari informasi tentang orang hilang
- **Family Members**: Monitoring status pencarian keluarga yang hilang

## ✨ Features

- ✅ **CRUD Operations** - Create, Read, Update, Delete missing person records
- ✅ **Search & Filter** - Advanced search dengan multiple criteria
- ✅ **Case Status Tracking** - Track status dari setiap kasus (reported, investigating, found, closed)
- ✅ **Geographic Information** - Location-based search dan tracking
- ✅ **RESTful API** - Well-designed REST endpoints
- ✅ **Database Migrations** - Automated database schema management
- ✅ **Validation** - Input validation untuk data integrity
- ✅ **Error Handling** - Comprehensive error handling dan logging
- ✅ **Pagination** - Efficient data pagination untuk large datasets
- ✅ **Authentication** (Coming Soon) - JWT-based authentication
- ✅ **File Upload** (Coming Soon) - Photo upload untuk missing persons

## 🏗 Architecture

Project ini mengikuti **Clean Architecture** pattern:

```
┌─────────────────────────────────────┐
│         API Layer (HTTP)            │
│     (Handlers/Controllers)          │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│        Business Logic Layer         │
│          (Use Cases)                │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│        Data Access Layer            │
│         (Repository)                │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│           Database                  │
│        (PostgreSQL)                 │
└─────────────────────────────────────┘
```

### Design Principles

- **Separation of Concerns**: Setiap layer memiliki tanggung jawab yang jelas
- **Dependency Inversion**: High-level modules tidak depend pada low-level modules
- **Single Responsibility**: Setiap komponen memiliki satu tanggung jawab
- **Testability**: Easy to test dengan mocking dependencies

## 🔧 Prerequisites

Pastikan sistem Anda sudah terinstall:

- **Go** v1.21 atau lebih baru - [Download](https://go.dev/dl/)
- **PostgreSQL** v14 atau lebih baru - [Download](https://www.postgresql.org/download/)
- **Git** - [Download](https://git-scm.com/downloads)
- **Make** (optional, untuk development shortcuts)

## 📥 Installation

### 1. Clone Repository

```bash
git clone https://github.com/Mhabib34/missing-person-service.git
cd missing-person-service
```

### 2. Install Dependencies

```bash
go mod download
go mod verify
```

### 3. Setup Database

#### Windows

```bash
# Buka psql sebagai postgres user
psql -U postgres

# Buat database
CREATE DATABASE missing_person_db;

# Buat user (optional)
CREATE USER missing_person_user WITH ENCRYPTED PASSWORD 'your_secure_password';
GRANT ALL PRIVILEGES ON DATABASE missing_person_db TO missing_person_user;

\q
```

#### macOS

```bash
# Install PostgreSQL (jika belum)
brew install postgresql@14
brew services start postgresql@14

# Buat database
createdb missing_person_db

# Atau via psql
psql postgres
CREATE DATABASE missing_person_db;
\q
```

#### Linux (Ubuntu/Debian)

```bash
# Install PostgreSQL
sudo apt update
sudo apt install postgresql postgresql-contrib

# Start service
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Buat database
sudo -u postgres createdb missing_person_db

# Atau via psql
sudo -u postgres psql
CREATE DATABASE missing_person_db;
\q
```

### 4. Run Database Migrations

```bash
# Migrations akan otomatis run saat aplikasi start
# Atau jalankan manual:
go run cmd/migrate/main.go
```

## ⚙️ Configuration

### Environment Variables

Buat file `.env` di root directory:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_MODE=development  # development atau production

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=missing_person_db
DB_USER=postgres
DB_PASSWORD=your_postgres_password
DB_SSLMODE=disable       # disable untuk development, require untuk production

# Application Configuration
APP_NAME=Missing Person Service
APP_VERSION=1.0.0
APP_ENV=development

# Logging
LOG_LEVEL=debug          # debug, info, warn, error
LOG_FORMAT=json          # json atau text

# Pagination
DEFAULT_PAGE_SIZE=20
MAX_PAGE_SIZE=100
```

### Configuration File (Optional)

Alternatif untuk environment variables, buat `config.yaml`:

```yaml
server:
  host: localhost
  port: 8080
  mode: development

database:
  host: localhost
  port: 5432
  name: missing_person_db
  user: postgres
  password: your_password
  sslmode: disable
  max_connections: 25
  max_idle_connections: 5

logging:
  level: debug
  format: json
```

## 🚀 Running the Service

### Development Mode

```bash
# Run dengan hot reload (jika menggunakan air atau similar)
air

# Atau run langsung
go run cmd/api/main.go

# Dengan environment variables
PORT=8080 DB_NAME=missing_person_db go run cmd/api/main.go
```

### Production Mode

```bash
# Build binary
go build -o bin/missing-person-service cmd/api/main.go

# Run binary
./bin/missing-person-service

# Dengan environment file
export $(cat .env | xargs) && ./bin/missing-person-service
```

### Using Make (Optional)

Buat `Makefile` untuk shortcuts:

```makefile
.PHONY: run build test clean migrate

run:
	go run cmd/api/main.go

build:
	go build -o bin/missing-person-service cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

migrate:
	go run cmd/migrate/main.go

lint:
	golangci-lint run
```

```bash
# Kemudian jalankan:
make run      # Run aplikasi
make build    # Build binary
make test     # Run tests
```

## 📖 API Documentation

Base URL: `http://localhost:8080/api/v1`

### Missing Persons Endpoints

#### 1. Create Missing Person Report

```http
POST /api/v1/missing-persons
Content-Type: application/json

{
  "full_name": "John Doe",
  "age": 25,
  "gender": "male",
  "last_seen_location": "Central Park, New York",
  "last_seen_date": "2024-01-15T14:30:00Z",
  "description": "Wearing blue jeans and red jacket",
  "reporter_name": "Jane Doe",
  "reporter_contact": "+1234567890",
  "reporter_relation": "Sister"
}
```

**Response: 201 Created**
```json
{
  "success": true,
  "message": "Missing person report created successfully",
  "data": {
    "id": "uuid-here",
    "full_name": "John Doe",
    "age": 25,
    "gender": "male",
    "status": "reported",
    "last_seen_location": "Central Park, New York",
    "last_seen_date": "2024-01-15T14:30:00Z",
    "created_at": "2024-01-21T10:00:00Z",
    "updated_at": "2024-01-21T10:00:00Z"
  }
}
```

#### 2. Get All Missing Persons

```http
GET /api/v1/missing-persons?page=1&limit=20&status=reported&gender=male
```

**Query Parameters:**
- `page` (optional): Page number, default: 1
- `limit` (optional): Items per page, default: 20, max: 100
- `status` (optional): Filter by status (reported, investigating, found, closed)
- `gender` (optional): Filter by gender
- `search` (optional): Search by name or description

**Response: 200 OK**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "uuid",
        "full_name": "John Doe",
        "age": 25,
        "gender": "male",
        "status": "reported",
        "last_seen_location": "Central Park, New York",
        "last_seen_date": "2024-01-15T14:30:00Z",
        "created_at": "2024-01-21T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "page_size": 20,
      "total_items": 45,
      "total_pages": 3
    }
  }
}
```

#### 3. Get Missing Person by ID

```http
GET /api/v1/missing-persons/{id}
```

**Response: 200 OK**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "full_name": "John Doe",
    "age": 25,
    "gender": "male",
    "status": "reported",
    "last_seen_location": "Central Park, New York",
    "last_seen_date": "2024-01-15T14:30:00Z",
    "description": "Wearing blue jeans and red jacket",
    "reporter_name": "Jane Doe",
    "reporter_contact": "+1234567890",
    "reporter_relation": "Sister",
    "created_at": "2024-01-21T10:00:00Z",
    "updated_at": "2024-01-21T10:00:00Z"
  }
}
```

#### 4. Update Missing Person

```http
PUT /api/v1/missing-persons/{id}
Content-Type: application/json

{
  "status": "found",
  "description": "Found safe at grandmother's house"
}
```

**Response: 200 OK**
```json
{
  "success": true,
  "message": "Missing person updated successfully",
  "data": {
    "id": "uuid",
    "status": "found",
    "updated_at": "2024-01-21T15:00:00Z"
  }
}
```

#### 5. Delete Missing Person

```http
DELETE /api/v1/missing-persons/{id}
```

**Response: 200 OK**
```json
{
  "success": true,
  "message": "Missing person record deleted successfully"
}
```

#### 6. Update Case Status

```http
PATCH /api/v1/missing-persons/{id}/status
Content-Type: application/json

{
  "status": "investigating",
  "notes": "Case has been assigned to detective Smith"
}
```

**Response: 200 OK**

### Status Codes

| Status Code | Description |
|-------------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 400 | Bad Request - Invalid input |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

### Error Response Format

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": {
      "field": "age",
      "issue": "must be greater than 0"
    }
  }
}
```

## 🗄️ Database Schema

### Missing Persons Table

```sql
CREATE TABLE missing_persons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    image_status VARCHAR(20),
    last_seen TEXT NOT NULL,
    description TEXT,
    photo_id VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

```

### Case Status Types

- `reported` - Awal laporan diterima
- `investigating` - Sedang dalam investigasi
- `found` - Orang telah ditemukan
- `closed` - Kasus ditutup

## 🧪 Testing

### Run All Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with verbose output
go test -v ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run Specific Tests

```bash
# Test specific package
go test ./internal/repository

# Test specific function
go test -run TestCreateMissingPerson ./internal/repository
```

### Integration Tests

```bash
# Run integration tests (requires database)
go test -tags=integration ./test/integration/...
```

### Test Structure

```
test/
├── unit/              # Unit tests
│   ├── handler_test.go
│   ├── usecase_test.go
│   └── repository_test.go
├── integration/       # Integration tests
│   └── api_test.go
└── fixtures/          # Test fixtures dan mock data
    └── missing_persons.json
```

## 📁 Project Structure

```
missing-person-service/
├── api/                    # API specifications (OpenAPI/Swagger)
│   └── openapi.yaml
├── cmd/                    # Application entry points
│   ├── api/
│   │   └── main.go        # Main API server
│   └── migrate/
│       └── main.go        # Database migration tool
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   │   └── config.go
│   ├── handler/          # HTTP handlers (controllers)
│   │   └── missing_person_handler.go
│   ├── middleware/       # HTTP middlewares
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   ├── model/           # Data models/entities
│   │   └── missing_person.go
│   ├── repository/      # Data access layer
│   │   ├── postgres/
│   │   │   └── missing_person_repository.go
│   │   └── interface.go
│   ├── usecase/         # Business logic layer
│   │   └── missing_person_usecase.go
│   ├── util/            # Utility functions
│   │   ├── validator.go
│   │   └── response.go
│   └── database/        # Database connection
│       └── postgres.go
├── migrations/          # Database migrations
│   ├── 001_create_missing_persons_table.up.sql
│   └── 001_create_missing_persons_table.down.sql
├── test/               # Tests
│   ├── unit/
│   ├── integration/
│   └── fixtures/
├── .env.example        # Environment variables example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile           # Development shortcuts
└── README.md
```

### Architecture Layers

1. **Handler Layer** (`internal/handler/`)
   - HTTP request handling
   - Input validation
   - Response formatting

2. **Use Case Layer** (`internal/usecase/`)
   - Business logic
   - Orchestration
   - Transaction management

3. **Repository Layer** (`internal/repository/`)
   - Database operations
   - Query building
   - Data mapping

4. **Model Layer** (`internal/model/`)
   - Domain entities
   - DTOs (Data Transfer Objects)
   - Request/Response structures

## 💻 Development

### Code Style

Project ini mengikuti [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Vet code
go vet ./...
```

### Adding New Features

1. **Create Model** in `internal/model/`
2. **Create Repository Interface** in `internal/repository/`
3. **Implement Repository** in `internal/repository/postgres/`
4. **Create Use Case** in `internal/usecase/`
5. **Create Handler** in `internal/handler/`
6. **Add Routes** in `cmd/api/main.go`
7. **Write Tests** in `test/`
8. **Update Documentation**

### Database Migrations

```bash
# Create new migration
migrate create -ext sql -dir migrations -seq add_new_table

# Run migrations
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

# Rollback
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down 1
```

### Debugging

```bash
# Run with delve debugger
dlv debug cmd/api/main.go

# Or with VS Code, add to launch.json:
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch API",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/cmd/api/main.go"
    }
  ]
}
```

## 🚢 Deployment

### Docker

**Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /missing-person-service cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /missing-person-service .
EXPOSE 8080
CMD ["./missing-person-service"]
```

**docker-compose.yml:**
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_NAME=missing_person_db
      - DB_USER=postgres
      - DB_PASSWORD=postgres
    depends_on:
      - postgres

  postgres:
    image: postgres:14-alpine
    environment:
      - POSTGRES_DB=missing_person_db
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

**Run with Docker:**
```bash
# Build dan run
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop
docker-compose down
```

### Cloud Deployment

#### Heroku
```bash
# Login
heroku login

# Create app
heroku create missing-person-service

# Add PostgreSQL
heroku addons:create heroku-postgresql:hobby-dev

# Deploy
git push heroku main
```

#### AWS (EC2 + RDS)
1. Setup RDS PostgreSQL instance
2. Launch EC2 instance
3. Build and deploy binary
4. Configure environment variables
5. Setup reverse proxy (Nginx)

## 🤝 Contributing

Contributions are welcome! Berikut langkah-langkahnya:

1. **Fork** repository ini
2. **Create** feature branch (`git checkout -b feature/AmazingFeature`)
3. **Commit** changes (`git commit -m 'Add some AmazingFeature'`)
4. **Push** to branch (`git push origin feature/AmazingFeature`)
5. **Open** Pull Request

### Contribution Guidelines

- Ikuti Go best practices dan code style
- Tambahkan tests untuk fitur baru
- Update documentation jika diperlukan
- Pastikan semua tests pass sebelum submit PR
- Tulis commit messages yang jelas dan deskriptif

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Mhabib34**
- GitHub: [@Mhabib34](https://github.com/Mhabib34)

## 🙏 Acknowledgments

- Inspired by real-world missing person tracking systems
- Built with Go best practices and clean architecture principles
- Community feedback and contributions

## 📞 Support

Jika ada pertanyaan atau issues:

1. **Create Issue** di [GitHub Issues](https://github.com/Mhabib34/missing-person-service/issues)
2. **Check Documentation** terlebih dahulu
3. **Contact** maintainer via GitHub

---

**Made with ❤️ for helping find missing persons**

*If you find this project useful, please give it a ⭐*
