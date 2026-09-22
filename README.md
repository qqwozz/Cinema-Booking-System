# 🎬 Cinema Booking System

<p align="center">
  <strong>Concurrent cinema seat booking backend built with Go</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/Concurrency-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Concurrency">
  <img src="https://img.shields.io/badge/Race%20Detector-Tested-2ea44f?style=for-the-badge" alt="Race Detector">
</p>

<p align="center">
  A backend system for managing cinema seat reservations with a focus on
  <strong>concurrency safety</strong>, <strong>clean architecture</strong>
  and <strong>reliable storage</strong>.
</p>

---

## 📌 Overview

**Cinema Booking System** is a backend service written in Go for handling cinema seat reservations.

The main engineering challenge of the project is **concurrent booking**: multiple users may attempt to reserve the same seat at approximately the same time.

The system must guarantee that:

> **One seat → one successful booking.**

Even when multiple goroutines perform the same operation concurrently, the application must prevent double booking and maintain a consistent state.

---

## ✨ Features

* 🎟️ Cinema seat booking
* 🔒 Protection against double booking
* ⚡ Concurrent request handling
* 🧵 Goroutine-based concurrency
* 💾 Redis storage adapter
* 🧩 Storage abstraction through interfaces
* 🐳 Docker / Docker Compose
* 🧪 Automated tests
* 🚨 Go Race Detector support
* 🏗️ Separation of business logic and infrastructure
* 🔌 Replaceable storage implementation

---

## 🛠️ Tech Stack

| Technology         | Purpose                |
| ------------------ | ---------------------- |
| **Go**             | Backend application    |
| **Redis**          | External storage       |
| **Docker**         | Containerization       |
| **Docker Compose** | Local infrastructure   |
| **net/http**       | HTTP server            |
| **Go testing**     | Automated testing      |
| **Race Detector**  | Concurrency validation |

---

# 🏛️ Architecture

The project separates the HTTP layer, business logic and storage implementation.

```text
┌─────────────────────────────────────────────┐
│                  HTTP Layer                 │
│                                             │
│                HTTP Handlers                │
└──────────────────────┬──────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────┐
│               Business Layer                │
│                                             │
│                Booking Service              │
│                                             │
│  • booking logic                            │
│  • availability checks                      │
│  • concurrency control                      │
└──────────────────────┬──────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────┐
│                Storage Layer                │
│                                             │
│              Storage Interface              │
└───────────────┬─────────────────┬───────────┘
                │                 │
                ▼                 ▼
        ┌──────────────┐   ┌──────────────┐
        │ Memory Store │   │ Redis Store  │
        └──────────────┘   └──────────────┘
```

The business layer works with an abstraction rather than directly depending on Redis.

This makes the storage implementation replaceable and keeps infrastructure concerns isolated from the core booking logic.

---

# 📂 Project Structure

```text
Cinema-Booking-System/
│
├── cmd/
│   └── ...
│
├── internal/
│   ├── booking/
│   │   └── ...
│   │
│   └── adapters/
│       └── redis/
│           └── ...
│
├── static/
│   └── ...
│
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

### `cmd/`

Application entry points.

### `internal/booking/`

Core booking domain and business logic.

Responsible for:

* checking seat availability;
* creating bookings;
* validating booking operations;
* preventing duplicate bookings;
* interacting with storage.

### `internal/adapters/redis/`

Redis-specific storage implementation.

This adapter isolates Redis infrastructure from the rest of the application.

---

# ⚡ Concurrency

Concurrency is one of the main technical aspects of the project.

Consider two users trying to book the same seat:

```text
User A ────────────────┐
                       │
                       ▼
                  ┌───────────┐
                  │  Booking  │
                  │  Service  │
                  └─────┬─────┘
                       ▲
                       │
User B ────────────────┘
```

Without proper synchronization, both requests could observe the seat as available:

```text
Goroutine A ──► check ──► available ──► book
Goroutine B ──► check ──► available ──► book
                                      │
                                      ▼
                              DOUBLE BOOKING ❌
```

The desired behavior is:

```text
Goroutine A ──► SUCCESS
Goroutine B ──► CONFLICT
```

Only one concurrent operation is allowed to successfully reserve the same seat.

---

## Concurrent Booking Test

The project contains a dedicated concurrency test:

```text
TestConcurrentBooking_ExactlyOneWins
```

Multiple goroutines attempt to book the same resource simultaneously.

```text
                 ┌─ Goroutine 1 ─┐
                 ├─ Goroutine 2 ─┤
                 ├─ Goroutine 3 ─┤
Booking Request ─┼─ Goroutine 4 ─┼──► Booking Service
                 ├─ Goroutine 5 ─┤
                 └─ Goroutine N ─┘
```

Expected result:

```text
Successful bookings = 1
Failed bookings     = N - 1
```

This test is designed to verify that the booking operation remains correct under concurrent access.

---

# 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test ./... -v
```

Run tests without cached results:

```bash
go test ./... -v -count=1
```

Run with the Go Race Detector:

```bash
go test ./... -race
```

Recommended development command:

```bash
go test ./... -v -count=1 -race
```

The test suite focuses on:

* booking logic;
* seat availability;
* duplicate bookings;
* concurrent operations;
* storage behavior;
* race conditions.

---

# 🐳 Docker

Redis can be started using Docker Compose.

Start infrastructure:

```bash
docker compose up -d
```

Check running containers:

```bash
docker compose ps
```

Stop containers:

```bash
docker compose down
```

Rebuild and start:

```bash
docker compose up --build
```

---

# 🚀 Quick Start

## Requirements

Make sure the following tools are installed:

* Go 1.22+
* Docker
* Docker Compose

### 1. Clone the repository

```bash
git clone https://github.com/qqwozz/Cinema-Booking-System.git
cd Cinema-Booking-System
```

### 2. Download dependencies

```bash
go mod download
```

### 3. Start infrastructure

```bash
docker compose up -d
```

### 4. Run the application

```bash
go run ./cmd/...
```

> If the project contains a specific server entry point inside `cmd`, run that package directly.

### 5. Run tests

```bash
go test ./... -v -count=1 -race
```

---

# 🔄 Booking Flow

A typical booking operation follows this flow:

```text
                         Client
                           │
                           │ booking request
                           ▼
                    ┌─────────────┐
                    │ HTTP Handler│
                    └──────┬──────┘
                           │
                           ▼
                    ┌─────────────┐
                    │   Booking   │
                    │   Service   │
                    └──────┬──────┘
                           │
                 ┌─────────┼─────────┐
                 │         │         │
                 ▼         ▼         ▼
              Validate   Check    Reserve
              request    seat      seat
                           │
                           ▼
                    ┌─────────────┐
                    │   Storage   │
                    └──────┬──────┘
                           │
                           ▼
                        Response
```

If the requested seat is already reserved, the operation is rejected instead of creating another booking.

---

# 🧩 Design Principles

## Separation of Concerns

HTTP handling, business logic and infrastructure are separated into different layers.

## Dependency Inversion

The booking logic depends on storage abstractions instead of a concrete Redis implementation.

## Concurrency Safety

Shared state is protected from unsafe concurrent access.

## Testability

Business logic can be tested independently from external infrastructure.

## Replaceable Infrastructure

Storage implementations can be changed without rewriting the core booking logic.

---

# 🔌 Storage

The application is designed around a storage abstraction.

Conceptually:

```text
             ┌─────────────────┐
             │ Storage         │
             │ Interface       │
             └────────┬────────┘
                      │
             ┌────────┴────────┐
             │                 │
             ▼                 ▼
      Memory Store         Redis Store
```

This approach allows the application to use an in-memory implementation during development and testing while supporting Redis as an external storage backend.

---

# 📊 Performance & Benchmarks

The project can be benchmarked using Go's built-in benchmarking tools:

```bash
go test ./... -bench=. -benchmem
```

Example benchmark output can be added here once reproducible benchmark results are available.

> Benchmark numbers should be generated on a controlled environment and should not be hard-coded without a reproducible test.

---

# 🔍 Code Quality

Recommended checks before submitting changes:

```bash
gofmt -w .
```

```bash
go vet ./...
```

```bash
go test ./... -race
```

A typical development pipeline is:

```text
      Code
       │
       ▼
    gofmt
       │
       ▼
   go vet
       │
       ▼
   go test
       │
       ▼
   go test -race
       │
       ▼
     Build
```

---

# 🏗️ Architecture Decisions

### Why Go?

Go provides lightweight concurrency through goroutines and channels, making it a good fit for backend services that need to process multiple requests concurrently.

### Why Redis?

Redis provides fast in-memory data access and can be used as an external storage layer for booking state.

### Why interfaces?

Using interfaces decouples the business layer from infrastructure implementations.

This makes the application easier to test and allows storage implementations to be replaced without changing the booking logic.

### Why test with the Race Detector?

A booking system is particularly sensitive to race conditions because multiple requests may modify the same state simultaneously.

The Go Race Detector helps identify unsafe concurrent memory access during testing.

---

# 🗺️ Roadmap

* [ ] Complete API documentation
* [ ] PostgreSQL persistence
* [ ] User authentication
* [ ] JWT authentication
* [ ] Movie management
* [ ] Cinema management
* [ ] Showtimes
* [ ] Booking expiration
* [ ] Payment integration
* [ ] Redis distributed locking
* [ ] Message broker integration
* [ ] Prometheus metrics
* [ ] Grafana dashboards
* [ ] Structured logging
* [ ] OpenTelemetry
* [ ] GitHub Actions CI
* [ ] Load testing
* [ ] Testcontainers
* [ ] Kubernetes deployment

---

# 📚 Engineering Topics

This project is intended to provide practical experience with:

```text
Go
├── Goroutines
├── Mutex / RWMutex
├── Race Conditions
├── Interfaces
├── Dependency Injection
└── Error Handling

Backend
├── HTTP APIs
├── Service Layer
├── Storage Abstraction
└── Redis

Infrastructure
├── Docker
└── Docker Compose

Testing
├── Unit Tests
├── Integration Tests
├── Concurrent Tests
└── Race Detector
```

---

# 👨‍💻 Author

**qqwozz**

Backend Developer

GitHub: **[@qqwozz](https://github.com/qqwozz)**

---

<p align="center">
  <sub>Built with Go, Redis and Docker.</sub>
</p>