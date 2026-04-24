# 🎯 Turbo Taskify

A personal task management app with pastel design built with Next.js (frontend) + Go (backend) + MongoDB.

## 📁 Project Structure

```
turbo-taskify/
├── frontend/     # Next.js React app (port 3000)
├── backend/      # Go REST API (port 8080)
└── ...
```

## 🚀 Quick Start

### Frontend

```bash
cd frontend
pnpm install
pnpm dev
# http://localhost:3000
```

### Backend

```bash
cd backend
go mod download
go run cmd/api/main.go
# http://localhost:8080
```

### Backend with hot reload (recommended for dev)

```bash
go install github.com/air-verse/air@latest
cd backend
air
# http://localhost:8080 — rebuilds on every file save
```

### Backend with Docker (recommended)

Requires [Docker Desktop](https://www.docker.com/products/docker-desktop).

```bash
cd backend
cp .env.example .env   # then fill in your values
docker compose up --build
# API: http://localhost:8080
# MongoDB: localhost:27017
```

**Common commands:**

```bash
docker compose up --build     # start (rebuild image)
docker compose up -d          # start in background
docker compose down           # stop and remove containers
docker compose logs -f api    # stream API logs
docker compose logs -f mongo  # stream MongoDB logs
```

Verify the API is running:

```bash
curl http://localhost:8080/api/auth/me
# 401 Unauthorized → server is up
```
