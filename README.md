## docker setting
```shell
$ docker pull postgres:12-alpine
$ docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:14-alpine
$ docker exec -it postgres psql -U root
$ docker stop postgres
```

## migrate setting
```shell
$ docker exec -it postgres /bin/sh
$ brew install golang-migrate
$ migrate create -ext sql -dir db/migration -seq ini_bank
$ migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank" -verbose up
# disable ssl
$ migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

```


## sqlc setting
```shell
$ brew install sqlc

```


## for test on postgres
```shell
$ go get github.com/lib/pq
```


## testify install
```shell
$ go get github.com/stretchr/testify
```


## install gin
```shell
$ go get -u github.com/gin-gonic/gin
```

## install viper
```shell
$ go get github.com/spf13/viper
```


## install goMock
```shell
$ go install github.com/golang/mock/mockgen@v1.6.0
```


## install uuid
```shell
$ go get github.com/google/uuid
```


## install jwt-go
```shell
$ go get github.com/dgrijalva/jwt-go
```