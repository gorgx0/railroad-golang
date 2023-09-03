## postgres docker
```bash
docker run --rm  --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=railway -p 5432:5432 -it postgres 
```

## TODO
- add cli interface
- add db data persistence to docker
- externalize configuration
- make methods return err instead of panic
- add db versioning
- introduce proper tests
- 