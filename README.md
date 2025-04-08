# Movie Ticket Booking System

A microservices-based movie ticket booking system built with Go, gRPC, and Clean Architecture.

## Services

1. **Auth Service** (`:50051`) - Handles user authentication and authorization
2. **Movie Service** (`:50052`) - Manages movie listings and details
3. **Booking Service** (`:50053`) - Handles ticket bookings and seat management

## Tech Stack

- Go 1.21+
- PostgreSQL 15
- Redis 7
- gRPC + gRPC-Gateway
- Docker & Docker Compose

## Project Structure


## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Protocol Buffers compiler
- Migration tool (golang-migrate)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/DragonPow/movie_booking.git
cd movie_booking
```

2. Start dependencies (PostgreSQL, Redis):
```bash
make docker-up
```

3. Install protocol buffer plugins:
```bash
go mod tidy
go mod vendor
```

4. Generate protocol buffer code:
```bash
make proto
```

5. Run database migrations:
```bash
make migrate-up
```

6. Build and run services:
```bash
make build
make run-services
```

## Development

- Build all services: `make build`
- Run tests: `make test`
- Generate proto files: `make proto`
- Clean build artifacts: `make clean`

## Database Migrations

- Create new migration: `migrate create -ext sql -dir db/migrations/{service} -seq {name}`
- Run migrations up: `make migrate-up`
- Run migrations down: `make migrate-down`

## License

MIT License