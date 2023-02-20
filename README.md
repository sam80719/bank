## docker setting
```shell
$ docker pull postgres:12-alpine
$ docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:12-alpine
$ docker exec -it postgres psql -U root
$ docker stop postgres
```

## migrate setting
```shell
$ docker exec -it postgres /bin/sh
$ brew install golang-migrate
$ migrate create -ext sql -dir database/migration -seq ini_bank
$ migrate -path database/migration -database "postgresql://root:secret@localhost:5432/simple_bank" -verbose up
# disable ssl
$ migrate -path database/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

```