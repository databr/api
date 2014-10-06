# DataBR API


## [Trello com Tarefas](https://trello.com/b/3WLlqXpX/databr)

## Gerando TOKEN Private File

``` bash
$ openssl genrsa -out ./tmp/token-private-file.rsa 1024
$ openssl rsa -in ./tmp/token-private-file.rsa -pubout > ~/tmp/token-private-file.rsa.pub
```

## Configurando ENV

``` bash
export MONGO_DATABASE_NAME="databr"
export MONGO_URL="mongodb://dev"

export INFLUXDB_HOST="dev:8086"
export INFLUXDB_USERNAME="user1"
export INFLUXDB_PASSWORD="user1"
export INFLUXDB_DATABASE="databr"

export API_ROOT="http://localhost:3002"
export PORT=3002
export ENV="development"

export STATUSPAGEIO_ENABLE="false"
export INFLUXDB_ENABLE="false"

export MEMCACHE_URL="dev:11211"

export DATABASE_URL="postgres://duke:duke@dev/databr?sslmode=disable"

export TOKEN_PRIVATE_FILE="./tmp/token-private-file.rsa"
```

## Rodando

``` bash
$ go get github.com/tools/godep
$ godep restore
$ go get github.com/pilu/fresh
$ fresh
```
