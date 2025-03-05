# 📌 Project Name

This is a backend service built with **Golang** using the **Gin** web framework. It follows a structured architecture with **controllers, services, and repositories**, utilizing **SQLC** for database interactions with **PostgreSQL**. The project includes JWT-based authentication, logging, and environment configuration.

## 🛠️ Tech Stack

- **Golang 1.23.3**
- **Gin** - Web framework
- **PostgreSQL** - Database
- **SQLC** - SQL query generator
- **Golang Migrate** - Database migration tool
- **JWT** - Authentication & Authorization
- **Viper** - Configuration management
- **Lumberjack** - Log rotation
- **CORS Middleware** - `gin-contrib/cors`
- **PGX** - PostgreSQL driver
- **ULID** - Unique ID generator

---

## Requirements

- make sure to install [SQLC](https://docs.sqlc.dev/en/stable/overview/install.html)
- make sure to install [Golang Migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) as migrations

## 🚀 How to Run

### 1️⃣ Clone the Repository

```sh
git clone git@github.com:Alter-IO/boilerplate-golang.git
cd boilerplate-golang
```

### 2️⃣ Setup Environment Variables

Copy a `example.toml` file and configure it:

```sh
mv server/config/example.toml server/config/app.toml
```

### 3️⃣ Install Dependencies

```sh
go mod tidy
```

### 4️⃣ Run Database Migration

Make sure you have **golang-migrate** installed. If not, install it:

```sh
brew install golang-migrate  # macOS
sudo apt install golang-migrate  # Linux
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest # Golang install
```

Then, run the migrations, for command you can check on Makefile in root directory:

```sh
make migrate-up
make migrate-down
```

### 5️⃣ Start the Server

```sh
go run server/cmd/http/main.go
```

---

## 📂 Project Structure

```
📦 project-directory
├── 📂 db                 # Database-related files
│   ├── 📂 migrations     # Database migration files
│   └── 📂 query          # SQLC query files
├── 📂 docker             # Docker-related files
├── 📂 server             # Main application source code
│   ├── 📂 cmd/http       # Entry point for the HTTP server
│   │   ├── 📂 modules    # Application modules
│   │   └── main.go       # Main entry file
│   ├── 📂 config        # Configuration files
│   ├── 📂 controllers   # Handles HTTP requests
│   ├── 📂 helpers       # Utility functions
│   ├── 📂 logs          # Logging files
│   ├── 📂 repositories  # Database interactions (SQLC generated)
│   ├── 📂 routes        # API route definitions
│   └── 📂 service       # Business logic layer
├── .gitignore           # Git ignore rules
├── .air.toml            # Air (Live reload tool) configuration
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── Makefile             # Build and run automation
├── README.md            # Project documentation
└── sqlc.yaml            # SQLC configuration file
```
