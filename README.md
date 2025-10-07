Polyglot Microservices Benchmark

Навчальний та демонстраційний проєкт, який показує, як Go, Rust і Elixir можуть співпрацювати в мікросервісній архітектурі. Включає приклад із Docker Compose для локального підняття.

Мета

показати pattern’и обміну повідомленнями між сервісами;

порівняти продуктивність/затримки між мовами для типових завдань;

дати навчальний приклад для розробників.

Компоненти

API Gateway (Go) — приймає HTTP-запити, маршрутизує в чергу.

Compute Service (Rust) — виконує CPU-bound/IO-bound обчислення.

Event Processor (Elixir) — обробляє події у реальному часі, агрегує результат.

Message Broker — NATS або Kafka (для демонстрації pub/sub).

DB — Postgres або Redis для зберігання результатів.

Приклад сценарію

Користувач робить POST /tasks до Go API.

Go заливає задачу в NATS.

Rust сервіс підбирає задачу, виконує обчислення (наприклад, heavy hashing або парсинг), повертає результат у іншу тему.

Elixir підписаний на результатні події, агрегує їх і пушить обновлення через WebSocket до UI.

Docker Compose (скорочена схема)
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

Приклад API (Go)
// POST /tasks
type TaskRequest struct {
  Input string `json:"input"`
}

Метрики та бенчмарки

Вимірювати latency від POST до фінального збереження.

Кількість завдань/сек для кожного сервісу.

CPU/RAM usage під навантаженням.

Можна вбудувати Grafana/Prometheus для візуалізації.

Навчальна цінність

Показує, як вибирати мову для певної задачі: Rust для обчислень, Go для API, Elixir для concurrency/real-time.

Демонструє патерни комунікації: pub/sub, request/reply, event sourcing.

Quickstart

Клон:

git clone https://github.com/you/polyglot-benchmark.git
cd polyglot-benchmark


Запусти:

docker-compose up --build


curl -X POST http://localhost:8080/tasks -d '{"input":"hello"}' -H "Content-Type: application/json"

Далі (ідеї)

Додати benchmarking suite (wrk/hey) та автоматичний звіт.

Підключити CI для запуску автоматичних тестів та бенчів.

Розширити сценарії: stream processing, ML inference.

Contributing

PR’и вітаються. Окремі папки для кожної мови з доками по локальному запуску та розробці.

License

MIT

Ліцензія та атрибуція

Локальні файли ліцензії для цього підпроєкту: див. `./LICENSE` і `./NOTICE`.
Обов'язкова атрибуція: автор — Architecture No. 7; модифікатори мають бути вказані.