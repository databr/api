# DataBR API [![Build Status](https://travis-ci.org/databr/api.svg?branch=master)](https://travis-ci.org/databr/api) [![Coverage Status](https://coveralls.io/repos/databr/api/badge.png)](https://coveralls.io/r/databr/api)

Join us on IRC at #databr to chat with other databr maintainers! ([web access](http://webchat.freenode.net/?channels=databr))

## Gerando TOKEN Private File

``` bash
$ openssl genrsa -out ./tmp/token-private-file.rsa 1024
$ openssl rsa -in ./tmp/token-private-file.rsa -pubout > ~/tmp/token-private-file.rsa.pub
```

## Configurando ENV

``` bash
export MONGO_DATABASE_NAME="databr"
export MONGO_URL="mongodb://dev"

export API_ROOT="http://localhost:3002"
export PORT=3002
export ENV="development"

export STATUSPAGEIO_ENABLE="false"

export PRIVATE_KEY="qweasdzxc"

```

## Rodando

``` bash
$ go get github.com/tools/godep
$ godep restore
$ go get github.com/pilu/fresh
$ fresh
```
