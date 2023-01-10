# REST API with Postgres I'm using for learning Go

The idea is to create an app (with both BE and FE (perhaps, even with mobile app) parts) where you can track the books you've read.

Not sure what the business logic will be exactly though. Gonna figure it out during the process.

## 1. Create and run Postgres in Docker container:

### If you're running this project for the first time:

```bash
$ docker run --name mabooks-db -p 5430:5431 -e POSTGRES_PASSWORD=qweqwe -d postgres
```

### If you want to start this app again:

```bash
$ docker start mabooks-db
```

## 2. Do migrations for db:

```bash
$ migrate -path ./migrations -database 'postgres://postgres:qweqwe@localhost:5430/postgres?sslmode=disable' up
```

## 3. Run the app:

```bash
$ go run cmd/main.go
```

## 4. Make call to API:

Call the API through `localhost:9999`.

---

## _Side notes_

- ### To drop the data in Postgres:

```bash
$ migrate -path ./migrations -database 'postgres://postgres:qweqwe@localhost:5430/postgres?sslmode=disable' down
```

- ### To connect to Postgres inside container's terminal in Docker:

```bash
$ psql -U postgres
```

- ### To make migration clean (whatever it means). It may be needed for running `migration down` script

```bash
$ update schema_migrations set version='000001', dirty=false;
```
