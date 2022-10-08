# {{.ProjectName}}

## Migrations

### How to create
```
docker compose run -v $(pwd)/db/migrations:/migrations migrate create -ext sql -dir migrations <migration_name>
```

### How to apply

```
docker compose run -v $(pwd)/db/migrations:/migrations migrate -path=/migrations/ -database "postgres://postgres@postgres:5432/postgres?sslmode=disable&search_path=public" up
```