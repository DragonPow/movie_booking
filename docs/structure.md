# Codebase Architecture – gRPC + HTTP Gateway (Single Go Module)

## Overview

This codebase is designed to support **gRPC-first communication** while optionally **extending to REST via gRPC-Gateway**. It uses a **single Go module** (`go.mod` at root), with a clean and scalable directory structure suitable for both monolith and microservice-ready architecture.

---

## Directory Structure

```plaintext
movie_booking/
├── go.mod
├── go.sum
│
├── cmd/
│   ├── auth/
│   │   └── main.go                  # Entry point for Auth service
│   ├── movie/
│   └── booking/
│
├── api/
│   ├── auth/auth.proto              # Protobuf definitions (gRPC + HTTP annotations)
│   ├── movie/movie.proto
│   └── booking/booking.proto
│
├── gen/
│   └── proto/                       # Generated gRPC & gRPC-Gateway code (DO NOT EDIT)
│       ├── auth/
│       │   ├── auth.pb.go
│       │   ├── auth_grpc.pb.go
│       │   └── auth.pb.gw.go
│       └── ...
│
├── internal/
│   ├── auth/
│   │   ├── service/                 # gRPC service implementations
│   │   ├── usecase/                 # Business logic layer
│   │   ├── repository/              # Database operations
│   │   ├── model/                   # Data structures, DTOs, Entities
│   │   └── middleware/              # Auth interceptors, logging, etc.
│   └── ...
│
├── pkg/
│   ├── config/                      # Config loader from ENV or YAML
│   ├── logger/                      # Logging utilities (zap, logrus, etc.)
│   └── grpcutil/                    # Helper to setup gRPC server + REST Gateway
│
├── third_party/                    # Custom annotations, Google API definitions
│
├── scripts/
│   └── generate.sh                 # Protobuf + Gateway code generation script
│
├── deployments/
│   ├── docker-compose.yml
│   └── k8s/
│       ├── auth-deployment.yaml
│       └── ...
│
└── README.md
```

---

## 🧭 Responsibilities

| Path / Directory | Description |
|------------------|-------------|
| `cmd/<service>/main.go` | Main application entry point for each service |
| `api/` | Protobuf files defining API contracts with HTTP annotations |
| `gen/` | Auto-generated `.pb.go`, `.pb.gw.go` from `.proto` files |
| `internal/<service>/` | Domain-specific logic (Clean Architecture style) |
| `pkg/` | Shared helper code, reusable across services |
| `third_party/` | Optional: Protobuf options, HTTP bindings |
| `scripts/generate.sh` | Compiles all `.proto` into Go and HTTP gateway code |
| `deployments/` | Docker & Kubernetes deployment configurations |

---

## 🧩 Protobuf Design Conventions

- Use `google.api.http` annotations for REST mappings.
- Naming convention: snake_case for packages, CamelCase for messages and RPCs.
- Set the Go package properly:

```proto
option go_package = "movie_booking/gen/proto/auth";
```

### Sample `auth.proto`

```proto
syntax = "proto3";

package auth;

option go_package = "movie_booking/gen/proto/auth";

import "google/api/annotations.proto";

service AuthService {
  rpc Login(LoginRequest) returns (LoginResponse) {
    option (google.api.http) = {
      post: "/v1/login"
      body: "*"
    };
  }
}

message LoginRequest {
  string email = 1;
  string password = 2;
}

message LoginResponse {
  string token = 1;
}
```

---

## 🛠️ Code Generation (`scripts/generate.sh`)

```bash
#!/bin/bash

protoc -I api/protobuf \
       -I ./third_party \
       --go_out=gen/proto --go_opt=paths=source_relative \
       --go-grpc_out=gen/proto --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=gen/proto --grpc-gateway_opt=paths=source_relative \
       api/protobuf/auth/v1/auth.proto \
       api/protobuf/movie/v1/movie.proto \
       api/protobuf/booking/v1/booking.proto
```

Run it using:

```bash
bash scripts/generate.sh
```

---

## Testing Strategy

- Each service has isolated business logic for easy unit testing.
- Use mock interfaces for `repository/` and `usecase/`.
- Add gRPC and REST integration tests under `internal/<service>/test/`.

---

## Coding Conventions & Rules

| Topic | Convention |
|-------|------------|
| Struct Naming | CamelCase, with JSON & DB tags: `json:"email" db:"email"` |
| Function Naming | Descriptive: `HandleLogin`, `NewToken`, `ToDTO` |
| Error Handling | Use `errors.Wrap`, return descriptive errors |
| API Design | gRPC first, REST is derived via gateway |
| Package Naming | `internal/auth/service`, `internal/auth/repository`, etc. |
| Generated Code | Always placed under `gen/proto/<service>` |
| Imports | Always import proto from `gen/proto/auth`, not from `api/` |
| Dependency Management | Use `go mod tidy` to keep `go.mod` clean |
| Logging | Use `pkg/logger` for structured logging |

Perfect! 🎯 Since you're aiming for a **scalable, developer-friendly Go project** for a **movie ticket booking system**, below is a detailed expansion of the **recommended tech stack**, tools, and utilities that will **boost development speed, maintainability, and productivity**.

All of these are **compatible with your gRPC + HTTP Gateway + single-module Golang repo structure**.

---

# Extended Tech Stack for Fast & Clean Golang Development

## Core Stack Recap

| Component         | Technology         | Purpose |
|------------------|--------------------|---------|
| Language          | Go                 | Backend system |
| Command Args      | `cobra`            | CLI commands |
| API Transport     | gRPC + REST (Gateway) | Main communication layer |
| DB                | PostgreSQL         | Primary relational database |
| Config Management | Viper              | Read from `.yaml`, `.env`, flags |
| ORM / DB Tool     | `sqlc`             | Generate type-safe SQL & models |
| Docker            | Docker / Compose   | Dev and deployment orchestration |
| Codegen           | `protoc`, `grpc-gateway` | Generate gRPC & REST |
| Logging           | `zap`  | High-performance structured logging |
| Migration Tool    | `golang-migrate`   | Manage DB schema versions |
| Auth              | `JWT` + middleware | Secure endpoints |
| Linting & Format  | `golangci-lint`, `goimports`, `revive` | Code quality |
| Testing           | `testify`, `mockgen` | Unit & integration testing |
| Monitoring        | Prometheus + Grafana (optional) | System observability |

---

## Configuration – Viper

**Tool:** [`spf13/viper`](https://github.com/spf13/viper)  
**Use Case:** Load configs from `.yaml`, `.env`, env vars, flags, or remote sources.

### Sample Config Structure

```yaml
# config/config.yaml
server:
  grpc_port: 50051
  http_port: 8080
  read_timeout: 10s
  write_timeout: 10s

database:
  driver: postgres
  dsn: "postgres://user:pass@localhost:5432/moviedb?sslmode=disable"
```

### ✅ Sample Go Loader

```go
v := viper.New()
v.SetConfigName("config")
v.AddConfigPath("./config")
v.SetConfigType("yaml")

if err := v.ReadInConfig(); err != nil {
    log.Fatal("Config load error:", err)
}

config := struct {
    Server struct {
        GRPCPort string `mapstructure:"grpc_port"`
    }
    Database struct {
        DSN string `mapstructure:"dsn"`
    }
}{}

if err := v.Unmarshal(&config); err != nil {
    log.Fatal("Config parse error:", err)
}
```

---

## 🧾 Database Layer – sqlc

**Tool:** [`sqlc`](https://github.com/kyleconroy/sqlc)  
**Use Case:** Auto-generate **type-safe queries and models** from `.sql` files.

### ✅ Project Structure

```
internal/
└── auth/
    └── repository/
        ├── queries.sql        # Your SQL queries
        ├── db.go              # SQLC generated interface
        └── models.go          # Structs
```

### Sample SQL

```sql
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
```

### Sample `sqlc.yaml` config

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "internal/auth/repository/queries.sql"
    schema: "scripts/migrations"
    gen:
      go:
        package: "repository"
        out: "internal/auth/repository"
```

Run:

```bash
sqlc generate
```

---

## 🧪 Migrations – golang-migrate

**Tool:** [`golang-migrate/migrate`](https://github.com/golang-migrate/migrate)  
**Use Case:** Database schema versioning.

Where to Place Migrations (Still Per-Service)

```plaintext
movie_booking/
├── internal/
│   ├── auth/
│   │   ├── repository/
│   │   └── migrations/           # ✅ Migration files only for auth-service
│   ├── movie/
│   │   ├── repository/
│   │   └── migrations/           # ✅ Migration files only for movie-service
│   ├── booking/
│   │   ├── repository/
│   │   └── migrations/
│   └── ...
```

---

## 🪵 Logging – zap

- `zap` – structured, high-performance logging by Uber
### ✅ Example with `zap`

```go
logger, _ := zap.NewProduction()
defer logger.Sync()
logger.Info("server started", zap.String("port", ":50051"))
```

---

## 🧰 Dev Tools

| Tool | Usage |
|------|-------|
| `mockgen` | Generate mock interfaces for testing |
| `golangci-lint` | Linting and formatting |
| `protoc-gen-go` / `grpc-gateway` | gRPC + REST code generation |
| `buf` (optional) | Proto linting and registry |

---

## 🧼 Lint & Format – golangci-lint

Add `.golangci.yml`:

```yaml
run:
  timeout: 3m
linters:
  enable:
    - govet
    - revive
    - staticcheck
    - errcheck
    - gocyclo
    - goimports
```

Run:

```bash
golangci-lint run ./...
```

---

## 🔁 Authentication – JWT Middleware

Use `github.com/golang-jwt/jwt/v5` + `grpc.UnaryInterceptor` or `Echo/Fiber middleware` for REST.

### ✅ gRPC Interceptor

```go
func AuthInterceptor(...) grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{},
        info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        token := extractToken(ctx)
        if !validateJWT(token) {
            return nil, status.Error(codes.Unauthenticated, "unauthorized")
        }
        return handler(ctx, req)
    }
}
```