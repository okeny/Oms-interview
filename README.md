# 🏢 Building Management System API

A RESTful API built with **Go**, **PostgreSQL**, **Fiber**, and **SQLBoiler**, designed to manage buildings and apartments. It provides endpoints for creating, updating, retrieving, and deleting building and apartment records. The project uses **Cobra** for CLI operations and **Docker** for containerization.


## ✅ Prerequisites

- **Go**: v1.23+
- **Docker** & **Docker Compose**
- **PostgreSQL** (optional if not using Docker)
- **Tools**:
  ```bash
  go install github.com/volatiletech/sqlboiler/v4@latest
  go install github.com/golang/mock/mockgen@latest
  ```

## 🚀 Setup
git clone 
go mod tidy
install docker
configure env variables. sample is in env.example
run docker compose up
get the postman collection from the postman folder


### Configure Environment Variables

Copy the example `.env` file and update as needed:

```bash
cp .env.example .env
```

Example:

```env
DB_USER=postgres
DB_PASSWORD=something
DB_NAME=building_management
DB_PORT=5432
DB_HOST=db
API_PORT=8000
API_VERSION=v1
DB_SSLMODE=disable
```

### Start with Docker Compose

```bash
docker-compose up --build
```

API will be accessible at: [http://localhost:8000](http://localhost:8000)

---

## 🧩 Running the API (Without Docker)

Ensure the database is running locally, then:

```bash
go run main.go api
```

---

## 📡 API Endpoints

Base path: `/api/v1/`

### Apartments

**POST /api/v1/apartments**  
Upserts an apartment

```json
{
  "building_id": 1,
  "number": "A1",
  "floor": 1,
  "sq_meters": 50
}
```

**Success Response (200 OK)**

```json
{
  "message": "Apartment upserted successfully",
  "apartment": {
    "id": 1,
    "building_id": 1,
    "number": "A1",
    "floor": 1,
    "sq_meters": 50,
    "created_at": "...",
    "updated_at": "..."
  }
}
```

**Error Responses**

- `400`: Invalid input or building ID  
- `404`: Building not found

### TODO

- `GET /api/v1/apartments`: List all  
- `GET /api/v1/apartments/:id`: By ID  
- `DELETE /api/v1/apartments/:id`: Delete  
- **Buildings**: [Add similar documentation as features are added]

---

## 🔃 Database Migrations

### Apply (Up)

```bash
go run main.go migrations up
```

### Rollback (Down)

```bash
go run main.go migrations down
```

### Create New Migration

```bash
go run main.go migrations create <migration_name>
```

e.g.

```bash
go run main.go migrations create add_users_table
```

### Check Status

```bash
go run main.go migrations status
```

---

## 🧪 Testing

Run all tests:

formating code
```bash
gofmt -s -w .
```
Running golinter 
```bash
 golangci-lint run
```
```bash
go test ./... -v
```

With coverage:

```bash
go test ./... -cover
```

---

## 🪰 Generating Mocks

```bash
mockgen -source=interfaces/api/building/repository.go \
  -destination=mocks/interfaces/api/building/mock_repository.go \
  -package=building
```

---

## 🔄 SQLBoiler Model Generation

```bash
sqlboiler psql
```

Make sure your `sqlboiler.toml` is configured:

```toml
[psql]
dbname = "building_management"
host   = "localhost"
port   = 5432
user   = "postgres"
pass   = "root"
sslmode = "disable"

improvements
add test cases for the other packages 
add cache like redis to speed up the response