# AGENTS.md - Railway Golang Project

## Project Overview

This is a Go desktop application using [Fyne](https://fyne.io/) for GUI, PostgreSQL for data storage, and [go-staticmaps](https://github.com/flopp/go-staticmaps) for map rendering. It manages railway stations and displays them on a map.

## Build Commands

```bash
# Run the application
go run main.go

# Build the binary
go build -o railway .

# Run a single test (if tests exist)
go test -v -run TestName ./...

# Run all tests
go test ./...

# Format code
go fmt ./...

# Run go vet
go vet ./...

# Run golangci-lint (if installed)
golangci-lint run
```

### Database Migrations (goose)

```bash
# Set environment variables
export GOOSE_DBSTRING="host=localhost user=postgres dbname=railway sslmode=disable password=postgres"
export GOOSE_DRIVER=postgres

# Check migration status (use -driver flag)
goose -dir db.migrations -driver postgres "$GOOSE_DBSTRING" status

# Run migrations
goose -dir db.migrations -driver postgres "$GOOSE_DBSTRING" up

# Create new migration
goose -dir db.migrations -driver postgres "$GOOSE_DBSTRING" create migration_name sql
```

## Code Style Guidelines

### Imports

Organize imports in three groups with blank lines between:
1. Standard library packages
2. External/third-party packages
3. Internal/railway packages

```go
import (
    "database/sql"
    "encoding/json"
    "fmt"

    "fyne.io/fyne/v2"
    "github.com/flopp/go-staticmaps"

    "railway/config"
    "railway/database"
    "railway/model"
)
```

### Formatting

- Use `go fmt` for automatic formatting
- 4-space indentation (standard Go)
- No trailing whitespace
- Group related constants and variables

### Naming Conventions

- **Types**: PascalCase (e.g., `Station`, `DatabaseConfig`, `Rectangle`)
- **Functions/Methods**: PascalCase (e.g., `GetDBConnection`, `LoadStationsFromJsonFile`)
- **Variables**: camelCase (e.g., `dbConfig`, `stations`, `err`)
- **Constants**: PascalCase or camelCase depending on export status
- **Packages**: lowercase, short names (e.g., `config`, `menu`, `model`)
- **Database tables**: snake_case (e.g., `stations`)

### Types

- Use explicit types; avoid `var x int` when `x := 0` works
- Use appropriate integer sizes (e.g., `int16` for small IDs if space matters)
- Use `float64` for coordinates
- Use `sql.DB` for database connections (pass as pointer)

### Error Handling

- **Return errors instead of panicking** (see project TODO)
- Always check `err` after function calls that can fail
- Use `log.Panic` or `log.Panicf` only for truly unrecoverable errors
- In database operations, close resources with `defer` and check close errors
- Propagate errors up the call stack with `return err` or `return nil, err`

```go
// Good
func GetDBConnection(dbConfig config.DatabaseConfig) (*sql.DB, error) {
    db, err := sql.Open("postgres", dbString)
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Good - deferred close with error check
defer func(db *sql.DB) {
    if err := db.Close(); err != nil {
        log.Println(err)
    }
}(db)
```

### Database Operations

- Use parameterized queries with `$1`, `$2`, etc. (not string formatting)
- Always close `sql.Rows` with `defer`
- Use `db.Query()` for queries that return rows
- Use `db.Exec()` for commands that don't return data

### Structs and JSON

- Use struct tags for JSON mapping (e.g., `json:"field_name"`)
- Keep JSON field names lowercase in tags
- Use nested structs for related data (see `Station.Location`)

### Logging

- Use `log.Println` for non-fatal errors
- Use `log.Panicf` or `log.Panic` for fatal errors that should stop the program
- Use `fmt.Println` / `fmt.Printf` for user-facing output

### Fyne GUI

- Use `fyne.NewSize(width, height)` for dimensions
- Use `container.NewVBox()`, `container.NewHBox()` for layouts
- Use `widget.NewButton()`, `widget.NewLabel()` for UI elements
- Create new windows with `app.NewWindow(title)`

### General Patterns

- One function per responsibility
- Keep functions focused and concise
- Use early returns to reduce nesting
- Close resources (files, connections) with `defer`
- Avoid global state where possible

## Project Structure

```
railroad-golang/
├── main.go              # Application entry point
├── config/
│   └── config.go        # Configuration reading
├── database/
│   └── connection.go    # PostgreSQL connection
├── model/
│   ├── station.go       # Station model and operations
│   └── map.go           # Map rendering
├── menu/
│   └── menu.go          # Menu actions
├── db.migrations/       # Database migrations (goose)
├── config.json          # Application config
└── stations.json        # Station data
```

## Common Tasks

### Adding a new station field
1. Update `Station` struct in `model/station.go`
2. Add JSON tag for config loading
3. Add database column in migration
4. Update `Store()` method
5. Update `GetAllStations()` query

### Adding a new database migration
```bash
goose -dir db.migrations -driver postgres "$GOOSE_DBSTRING" create add_column sql
# Edit migration file in db.migrations/
goose -dir db.migrations -driver postgres "$GOOSE_DBSTRING" up
```

### Adding a new UI button
Add to the appropriate menu or create new container in `main.go`.
