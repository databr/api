[![Stories in Ready](https://badge.waffle.io/databr/api.png?label=ready&title=Ready)](https://waffle.io/databr/api)
# DataBR API [![Build Status](https://travis-ci.org/databr/api.svg?branch=master)](https://travis-ci.org/databr/api) [![Coverage Status](https://coveralls.io/repos/databr/api/badge.png)](https://coveralls.io/r/databr/api)

[![Join the chat at https://gitter.im/databr/api](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/databr/api?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)


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

export PRIVATE_KEY="qweasdzxc"

```

## Rodando

``` bash
$ go get github.com/tools/godep
$ godep restore
$ go get github.com/pilu/fresh
$ fresh
```
