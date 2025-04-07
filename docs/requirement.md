# Movie Ticket Booking System – Software Requirements

## Technology Stack

| Component      | Technology            |
|----------------|------------------------|
| Backend        | Golang (Fiber/Gin/Echo) |
| Database       | PostgreSQL              |
| Deployment     | Docker + Docker Compose |
| Auth           | JWT (or OAuth2)         |
| Caching        | Redis *(Optional)*      |
| Messaging      | Kafka / RabbitMQ *(Optional for future scaling)* |
| Monitoring     | Prometheus + Grafana *(Optional)* |

---

## Feature Prioritization

| Priority Level | Feature |
|----------------|---------|
| ⭐️⭐️⭐️ (Must Have - Phase 1) | User Registration/Login, Movie Listing, Theater & Showtimes, Seat Selection, Ticket Booking |
| ⭐️⭐️ (Should Have - Phase 2) | Booking Expiry (Temp Hold), Payment Gateway Integration, Booking History |
| ⭐️ (Nice to Have - Phase 3) | Admin Dashboard, Reports, Notifications, Promo Code System |

---

## Functional Requirements

### 1. User Management

- **User Roles**: Customer, Staff, Admin
- **Features**:
  - Register
  - Login / Logout
  - JWT-based Authentication
  - View/Edit profile (name, email, phone)
  - View booking history

---

### 2. Movie Management

- **Admin Features**:
  - Add/Edit/Delete movies
- **Movie Data**:
  - Title, Description
  - Duration, Language, Genre
  - Age Rating, Release Date
  - Poster URL, Trailer URL
- **User Features**:
  - View all movies (now showing / upcoming)
  - Filter by genre / date / location

---

### 3. Theater & Room Management

- **Admin Features**:
  - Add/Edit/Delete theaters
  - Add/Edit rooms per theater
- **Room Details**:
  - Name, Capacity, Seat Layout (grid-based)
  - Audio/Visual system support

---

### 4. Showtimes Management

- **Admin Features**:
  - Create showtimes: Movie + Theater Room + Time
- **User Features**:
  - View showtimes by date
  - Filter by movie or location

---

### 5. Seat Selection & Ticket Booking

- **User Flow**:
  1. Choose movie → select showtime
  2. See seat layout (available/reserved/booked)
  3. Select seats
  4. Confirm booking
- **System Logic**:
  - Concurrency-safe seat reservation
  - Use Redis or DB transactions for lock
  - Ticket contains: QR Code, booking code, movie/showtime info

---

### 6. Temporary Seat Hold

- When a user selects seats, hold them for **X minutes**.
- Auto-release if timeout or payment fails.
- Can use Redis with expiry keys.

---

### 7. Payment (Optional for simulation)

- Integration with payment gateway (e.g., Stripe, Momo, PayPal)
- Simulated payment for MVP
- Confirm booking only after payment success

---

### 8. Booking History

- User can view past & upcoming bookings
- Admin/Staff can view all bookings
- Cancel booking (if allowed)

---

### 9. Admin Panel (Phase 2+)

- Dashboard: stats of movies, revenue, seats sold
- Manage users, movies, theaters, showtimes, bookings

---

## Non-Functional Requirements

### 1. Performance & Scalability

- RESTful stateless APIs
- Horizontal scaling (via Docker/K8s)
- Redis caching for performance
- Load balancing supported

---

### 2. Security

- JWT Auth with refresh token (optional)
- Input validation (SQL injection, XSS, etc.)
- Rate limiting
- Secure password storage (bcrypt)

---

### 3. Deployment & Config

- Dockerize all services
- Docker Compose for local development
- Config via `.env` and config files
- Support CI/CD (optional)

---

### 4. Observability & Monitoring

- Logs: request logs, error logs, audit trails
- Health check endpoints
- Metrics exposed for Prometheus (optional)
- Alerting (optional)

---

### 5. Extensibility

- Codebase modularized
- Clear separation of domain logic
- Easy to plugin microservices later (e.g., notifications, analytics)