# Todo Go API

## Requirement

- Go 1.22+
- Docker

## How to run

**1. Clone repo**
```bash
git clone https://github.com/okarinn06/todo-go.git
cd todo-go
```

**2. Create file `.env`**
```bash
cp .env.example .env
```

**3. Run PostgreSQL in Docker**
```bash
docker compose up -d
```

**4. Download dependencies**
```bash
go mod tidy
```

**5. Run server**
```bash
go run main.go
```

Server run at: `http://localhost:8080`
