## postgres docker



```bash
mkdir data
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=railway -p 5432:5432 -it -e PG_DATA=/var/lib/postgresql/data -v $(pwd)/data:/var/lib/postgresql/data postgres 
```

## RUN PROJECT
```bash
go run main.go
```

## DATABASE VERSIONING
Install goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
or
```bash
brew install goose
```

setup env
```bash
export GOOSE_DBSTRING="host=localhost user=postgres dbname=railway sslmode=disable password=postgres"
export GOOSE_DRIVER=postgres
```

Create migrtation
```bash
~/go/bin/goose -v -dir db.migrations create create_stations_table sql
```

status
```bash
goose -dir db.migrations postgres "host=localhost user=postgres dbname=railway sslmode=disable password=postgres" status
```

migrate
```bash
goose -dir db.migrations postgres "host=localhost user=postgres dbname=railway sslmode=disable password=postgres" up
```


## TODO
- make methods return err instead of panic
- introduce proper tests
- add window status line