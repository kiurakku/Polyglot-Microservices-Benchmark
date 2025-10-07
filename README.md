# Polyglot Microservices Benchmark

An educational and demo project showing how **Go**, **Rust**, and **Elixir** can work together in a microservice architecture.  
Includes a **Docker Compose** example for local deployment.

---

## 🧠 Purpose

- Demonstrate messaging patterns between services.  
- Compare performance/latency across languages for common workloads.  
- Provide a hands-on learning example for developers.

---

## 🧩 Components

- **API Gateway (Go)** — handles HTTP requests and routes them to the message queue.  
- **Compute Service (Rust)** — performs CPU-bound and IO-bound computations.  
- **Event Processor (Elixir)** — processes real-time events and aggregates results.  
- **Message Broker** — *NATS* or *Kafka* (used for pub/sub demonstration).  
- **Database** — *Postgres* or *Redis* for result storage.

---

## 🚀 Example Scenario

1. The user sends a `POST /tasks` request to the Go API.  
2. Go pushes the task into NATS.  
3. The Rust service picks up the task, performs heavy computation (e.g., hashing or parsing), and publishes the result to another topic.  
4. The Elixir service subscribes to the result events, aggregates them, and pushes updates via WebSocket to the UI.

---

## 🐳 Docker Compose (short version)

```yaml
version: "3.8"
services:
  nats:
    image: nats:latest
    ports: ["4222:4222"]

  api:
    build: ./services/api-go
    ports: ["8080:8080"]
    depends_on: ["nats"]

  compute:
    build: ./services/compute-rust
    depends_on: ["nats"]

  processor:
    build: ./services/processor-elixir
    ports: ["4000:4000"]
    depends_on: ["nats"]

  db:
    image: postgres:15
    environment:
      POSTGRES_PASSWORD: example
