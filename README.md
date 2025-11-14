<h1 align="center">Golang Simple Restful API</h1>

<div align="center">

![Golang](https://img.shields.io/badge/-Golang-181818?logo=go)&nbsp;
![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-212121?logo=postgresql)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-F9FAFC?logo=docker)&nbsp;
![Bruno](https://img.shields.io/badge/-Bruno-1D1D1D?logo=bruno)&nbsp;

</div>

<p align="center">A simple RESTful API using **Golang**, **Gin**, and **GORM**</p>

---

## Table of Contents

-   [Features](#features)
-   [Requirements](#requirements)
-   [Project Structure](#project-structure)
-   [Docker Setup](#docker-setup)
-   [Running the Application](#running-the-application)
-   [Open Telemetry Integration](#open-telemetry-integration)

---

## Features

-   Simple Migration & Seeder
-   Simple clean architecture: handlers -> service/repository -> models
-   Logging using Uber Zap
-   Support build manual, binary release, and Docker image

## Requirements

-   Go 1.25+
-   PostgreSQL 17
-   Git
-   Docker & Docker Compose (optional)

## Runing the Application

### Development Mode

1. Install dependencies

```bash
go mod tidy
```

2. Running the application

> [!NOTE]
> Access the API at http://localhost:3090/api/students

```bash
go run cmd/server/main.go
```

### Compile Manual (Build Binary)

1. Compile to make executable file

> [!TIP]
> To compile for difference architecture (like Linux AMD64)
>
> ```bash
> GOOS=linux GOARCH=amd64 go build -o dist/golang-simple-restful-api ./cmd/server
> ```

```bash
go build -o dist/golang-simple-restful-api ./cmd/server
```

2. Running the executable file

```bash
./dist/golang-simple-restful-api
```

### Download Release (Prebuilt Binary)

> [!NOTE]
> This repository have a CI/CD pipeline that automate build, bump version, tagging, and release version using semantic versioning, so you can download the prebuilt binary on GitHub Release: [https://github.com/armandwipangestu/golang-simple-restful-api/releases](https://github.com/armandwipangestu/golang-simple-restful-api/releases)

After download di prebuilt binary, run this command

```bash
chmod +x golang-simple-restful-api
./golang-simple-restful-api
```

### Running with Docker

1. Build the image

```bash
docker build -t golang-simple-restful-api .
```

2. Run the image

```bash
docker run -p 3090:3090 --env-file .env golang-simple-restful-api
```

### Running with Docker Compose

1. Copy the `.env.example` and `.env.example.postgres`

```bash
cp .env.example .env
cp .env.example.postgres .env.postgres
```

2. Fill the value of `.env` and `.env.postgres` with your own configuration

```bash
APP_PORT=3090
DB_HOST=go_example_database
DB_PORT=5432
DB_USER=go
DB_PASS=golang
DB_NAME=go_simple_restful_api
DB_SSLMODE=disable
```

```bash
POSTGRES_DB="go_simple_restful_api"
POSTGRES_USER="go"
POSTGRES_PASSWORD="golang"
```

3. Runing the application using compose

```bash
docker compose up -d
```
