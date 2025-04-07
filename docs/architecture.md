# System Architecture

## Components
- API Gateway
- Auth Service
- Movie Service
- Booking Service
- PostgreSQL
- Redis
- Optional: Kafka, Prometheus, Grafana

## Diagram
```mermaid
graph LR
  A[API Gateway] -->|gRPC| B[Auth Service]
  A -->|gRPC| C[Movie Service]
  A -->|gRPC| D[Booking Service]
  B -->|PostgreSQL| E[Auth DB]
  C -->|PostgreSQL| F[Movie DB]
  D -->|PostgreSQL| G[Booking DB]
  D -->|Redis| H[Seat Holds]