# Subscribes

REST API service for managing user subscriptions built with Go, Gin, and PostgreSQL.

## Architecture

The project follows clean architecture principles with three-layer structure:

```
cmd/main.go                 — Application entry point
internal/
  handlers/                 — HTTP handlers (Gin controllers)
  service/                  — Business logic layer
  repository/               — Data access layer (raw SQL)
    models/                 — Domain models
pkg/db/                     — Database connection and configuration
migrations/                 — SQL migration files
docs/                       — Swagger/OpenAPI documentation
```

## Tech Stack

| Component       | Technology                          |
|-----------------|-------------------------------------|
| Language        | Go 1.25                             |
| Web Framework   | Gin                                 |
| Database        | PostgreSQL 15                       |
| Migrations      | golang-migrate                      |
| Logging         | Logrus (JSON format)                |
| API Docs        | Swagger (swaggo/gin-swagger)        |
| Containerization| Docker + Docker Compose             |

## API Endpoints

All endpoints are prefixed with `/api/subscribes`.

| Method   | Endpoint                                              | Description                          |
|----------|-------------------------------------------------------|--------------------------------------|
| `POST`   | `/api/subscribes`                                     | Create a new subscription            |
| `GET`    | `/api/subscribes/:serviceID`                          | Get subscription by ID               |
| `GET`    | `/api/subscribes/service/:serviceName`                | Get all subscriptions by service     |
| `GET`    | `/api/subscribes/user_subscriptions/:userID`          | Get all subscriptions for a user     |
| `PUT`    | `/api/subscribes/update_subscription/:id/:price`      | Update subscription price            |
| `DELETE` | `/api/subscribes/delete_subscription/:subscriptionID` | Delete a subscription                |
| `GET`    | `/api/subscribes/list/:startDate/:endDate`            | Get aggregated price sum by period   |

### Request Examples

**Create subscription**

```bash
curl -X POST http://localhost:8080/api/subscribes \
  -H "Content-Type: application/json" \
  -d '{
    "service_name": "Netflix",
    "price": 322,
    "user_id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "start_date": "03-2024",
    "end_date": "05-2024"
  }'
```

**Get user subscriptions**

```bash
curl http://localhost:8080/api/subscribes/user_subscriptions/f47ac10b-58cc-4372-a567-0e02b2c3d479
```

**Get price sum for a period**

```bash
curl http://localhost:8080/api/subscribes/list/01-2024/06-2024
```

## Database Schema

```sql
CREATE TABLE users (
    id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name     VARCHAR(50)  NOT NULL,
    email    VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE subscribes (
    id           SERIAL PRIMARY KEY,
    service_name VARCHAR(50) NOT NULL,
    price        INTEGER     NOT NULL CHECK (price > 0),
    user_id      UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    start_date   TIMESTAMP   NOT NULL,
    end_date     TIMESTAMP
);
```

## Getting Started

### Prerequisites

- Go 1.25+
- Docker + Docker Compose
- PostgreSQL 15 (or use Docker)

### Running with Docker Compose

```bash
docker-compose up --build
```

The API will be available at `http://localhost:7777`. Swagger UI is at `http://localhost:7777/swagger/index.html`.

### Running Locally

1. Start PostgreSQL (via Docker or locally):

```bash
docker-compose up -d db
```

2. Configure environment variables (create `.env`):

```
DB_HOST=<your_db_host>
DB_PORT=<your_db_port>
DB_USER=<your_db_user>
DB_PASSWORD=<your_db_password>
DB_NAME=<your_db_name>
```

3. Run migrations and start the server:

```bash
go run ./cmd/main.go
```

The server listens on port `8080`.

## Project Structure

- **`cmd/main.go`** — Initializes configuration, database, wires up layers, starts the HTTP server with graceful shutdown.
- **`internal/handlers`** — Parses HTTP requests, validates input, calls service methods, returns JSON responses.
- **`internal/service`** — Contains business logic: input validation, date parsing, price checks, UUID parsing.
- **`internal/repository`** — Executes raw SQL queries against PostgreSQL. Uses transactions for write operations.
- **`pkg/db`** — Reads DB config from environment variables and establishes the database connection.
- **`migrations/`** — Sequential migration files for schema creation and seed data.
- **`docs/`** — Auto-generated Swagger spec and Go wrappers.

## Swagger

Interactive API documentation is available at `/swagger/index.html` when the server is running.

To regenerate docs after code changes:

```bash
swag init -g cmd/main.go -o docs
```

## License

[MIT](LICENSE) © 2026 @AlGrushino
