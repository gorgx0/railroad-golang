## postgres docker
```bash
docker run --rm  --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=railway -p 5432:5432 -it postgres 
```

## RUN PROJECT
```bash
go run main.go
```


## TODO
- make methods return err instead of panic
- add cli interface
- add db data persistence to docker
- externalize configuration
- add db versioning
- introduce proper tests
