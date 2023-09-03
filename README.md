## postgres docker



```bash
mkdir data
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=railway -p 5432:5432 -it -e PG_DATA=/var/lib/postgresql/data -v $(pwd)/data:/var/lib/postgresql/data postgres 
```

## RUN PROJECT
```bash
go run main.go
```


## TODO
- make methods return err instead of panic
- externalize configuration
- add db versioning
- introduce proper tests
