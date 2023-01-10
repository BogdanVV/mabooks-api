# Basic REST API built with `gin` + `Postgres`

I'm using this project for learning Go and getting more familiar with SQL.

The idea is to create an app (with both BE and FE (perhaps, even with mobile app) parts) where you can track the books you've read. Maybe with some additional stuff, dunno yet.

Not sure what the business logic will be exactly though. Gonna figure it out during the process.

## Prerequisites:

- Go version >= `1.19.2`
- Docker

## 1. Create and run Postgres in Docker container:

### If you're running this project for the first time:

```bash
$ docker run --name mabooks-db -p 5430:5431 -e POSTGRES_PASSWORD=qweqwe -d postgres
```

_If you don't have [postgres docker image](https://hub.docker.com/_/postgres) installed, this command will pull it for you from dockerhub automatically. At least, I think so. Too lazy to check that :) If I'm wrong, run `$ docker pull postgres` first\_

### If you already have the container and just wanna start it:

```bash
$ docker start mabooks-db
```

## 2. Do initial migration for db:

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
