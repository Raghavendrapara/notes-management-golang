# go-backend — Personal Notes API

A minimal, production-style Go backend for a **personal notes management system**. It serves HTTP GET, POST, PUT, and DELETE with in-memory storage, structured for scalability and maintenance.

## Layout

```
go-backend/
├── cmd/api/           # Application entry point
├── internal/          # Private app code (not importable by other modules)
│   ├── config/        # Configuration from environment
│   ├── handler/       # HTTP handlers and middleware
│   ├── models/        # Request/response and domain types
│   ├── server/        # Server wiring and routing
│   ├── service/       # Business logic
│   └── store/         # Data access (in-memory; swappable for DB later)
├── go.mod
└── README.md
```

- **Handlers** — Parse HTTP, call service, write JSON.
- **Service** — Business rules and ID generation; no HTTP.
- **Store** — Interface-based persistence; current implementation is in-memory with dummy seed data.

## Run

```bash
cd C:\Users\rsrsr\Programs\go-backend
go mod tidy
go run ./cmd/api
```

Server listens on `:8080` by default. Override with `PORT`:

```bash
set PORT=3000
go run ./cmd/api
```

## API

| Method | Path           | Description        |
|--------|----------------|--------------------|
| GET    | /health        | Liveness check     |
| GET    | /notes         | List all notes     |
| GET    | /notes/:id     | Get one note       |
| POST   | /notes         | Create note        |
| PUT    | /notes/:id     | Update note        |
| DELETE | /notes/:id     | Delete note        |

### Examples

**List notes**
```bash
curl http://localhost:8080/notes
```

**Create note**
```bash
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"title\":\"My note\",\"body\":\"Some text\"}"
```

**Update note**
```bash
curl -X PUT http://localhost:8080/notes/1 -H "Content-Type: application/json" -d "{\"title\":\"Updated title\",\"body\":\"Updated body\"}"
```

**Get one note**
```bash
curl http://localhost:8080/notes/1
```

**Delete note**
```bash
curl -X DELETE http://localhost:8080/notes/1
```

## Practices used

- **Standard project layout** — `cmd/` for entrypoints, `internal/` for private code.
- **Dependency injection** — Store and service are constructed in `server.New` and passed into handlers.
- **Interfaces** — `NoteStore` allows swapping in-memory store for a DB without changing handlers or service.
- **Structured logging** — `slog` with request logging middleware.
- **Graceful shutdown** — SIGINT/SIGTERM triggers `Shutdown` with a timeout.
- **Config from env** — `PORT`, `ENV` with defaults so the app runs without config files.

## Next steps

- Add a real DB (e.g. PostgreSQL) by implementing `store.NoteStore` in a new package (e.g. `internal/store/postgres`).
- Add validation (e.g. `go-playground/validator`) and request size limits.
- Add tests for handlers and service using the store interface with a fake in-memory store.
