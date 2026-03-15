# Trackr — Backend

REST API for the Trackr personal expense tracker, built with Go and Gin, backed by PostgreSQL via Supabase.

## Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) — for generating type-safe SQL queries
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) — for running database migrations
- A Supabase project (or any PostgreSQL instance)

## Getting started

### 1. Clone and navigate

```bash
git clone https://github.com/santtisosa/trackr.git
cd trackr/backend
```

### 2. Set up environment variables

```bash
cp .env.example .env
```

Fill in the values in `.env`:

| Variable                   | Description                              |
|----------------------------|------------------------------------------|
| `PORT`                     | Port the server listens on (default 8080)|
| `DATABASE_URL`             | PostgreSQL connection string             |
| `SUPABASE_URL`             | Your Supabase project URL                |
| `SUPABASE_ANON_KEY`        | Supabase anon/public key                 |
| `SUPABASE_SERVICE_ROLE_KEY`| Supabase service role key                |
| `OPENAI_API_KEY`           | OpenAI API key (for AI features)         |
| `GOOGLE_CLOUD_VISION_KEY`  | Google Cloud Vision key (for OCR)        |

### 3. Install dependencies

```bash
go mod download
```

### 4. Run database migrations

Install the migrate CLI if you haven't already:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run migrations:

```bash
migrate -path db/migrations -database "$DATABASE_URL" up
```

### 5. Generate sqlc code (optional, if you modify SQL queries)

Install sqlc:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Generate:

```bash
sqlc generate
```

### 6. Run the server

```bash
go run ./cmd/api
```

The server starts on `http://localhost:8080`.

## Health check

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

## Project structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go          # Entry point
├── internal/
│   ├── handler/             # HTTP handlers (controllers)
│   ├── service/             # Business logic
│   ├── repository/          # Database access layer
│   └── model/               # Domain models and types
├── db/
│   └── migrations/          # SQL migration files
├── .env.example
├── .gitignore
├── go.mod
└── go.sum
```
