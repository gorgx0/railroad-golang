# Railway Golang Project

## Project Overview
This is a Go-based desktop application designed to manage and visualize railway station data. It features a Graphical User Interface (GUI) built with **Fyne**, allowing users to load station data from JSON files, persist it to a **PostgreSQL** database, and visualize station locations on a map.

## Tech Stack
*   **Language:** Go (Golang)
*   **GUI Framework:** [Fyne](https://fyne.io/) (v2)
*   **Database:** PostgreSQL
*   **Database Migrations:** [Goose](https://github.com/pressly/goose)
*   **Map Visualization:** `github.com/flopp/go-staticmaps`
*   **Containerization:** Docker (for the database)

## Prerequisites
*   **Go:** 1.20 or later
*   **Docker:** For running the PostgreSQL instance
*   **Goose:** Database migration tool
    ```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```

## Getting Started

### 1. Database Setup
Start a PostgreSQL instance using Docker:
```bash
mkdir data
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=railway -p 5432:5432 -it -e PG_DATA=/var/lib/postgresql/data -v $(pwd)/data:/var/lib/postgresql/data postgres
```

### 2. Configure Environment
Set up the environment variables required for `goose` migrations:
```bash
export GOOSE_DBSTRING="host=localhost user=postgres dbname=railway sslmode=disable password=postgres"
export GOOSE_DRIVER=postgres
```

### 3. Run Migrations
Apply the database schema:
```bash
# Check status
goose -dir db.migrations postgres "$GOOSE_DBSTRING" status

# Migrate up
goose -dir db.migrations postgres "$GOOSE_DBSTRING" up
```

### 4. Application Configuration
Ensure `config.json` exists in the root directory with the correct database credentials and input file path. Example structure based on `config/config.go`:
```json
{
  "database": {
    "host": "localhost",
    "port": 5432,
    "db": "railway",
    "user": "postgres",
    "password": "postgres"
  },
  "stations_file": "stations.json"
}
```

### 5. Run the Application
Start the GUI:
```bash
go run main.go
```

## Project Structure

*   `main.go`: Application entry point. Initializes the Fyne window and event handlers.
*   `config/`: Configuration loading logic (`config.json`).
*   `database/`: Database connection handling.
*   `db.migrations/`: SQL migration files managed by Goose.
*   `menu/`: Application logic for menu actions (loading, printing, removing stations).
*   `model/`: Data models (`Station`, `Map`) and database interaction logic.
*   `stations.json`: Sample dataset for railway stations.

## Key Features
*   **Load Stations:** Imports station data from a JSON file into the database.
*   **Print Stations:** Retrieves and displays station data from the database.
*   **Remove Stations:** Clears all station data from the database.
*   **Show Map:** Generates and displays a map visualization of stored stations.

## Development Notes
*   **Testing:** Currently, there are no explicit tests. Adding unit tests for the `model` package is a recommended TODO.
*   **Error Handling:** The application currently relies heavily on `log.Panicf` for error handling. Refactoring to more graceful error management is advised.
