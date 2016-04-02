# DataBR API [![Build Status](https://travis-ci.org/databr/api.svg?branch=master)](https://travis-ci.org/databr/api) [![Coverage Status](https://coveralls.io/repos/databr/api/badge.png)](https://coveralls.io/r/databr/api)

Junte-se a nós no Slack [clicando aqui](http://databr.herokuapp.com/), venha conversar com os mantenedores e interessados em API e dados publicos 


## Usando Docker

```
$ docker-compose up // acesse http://localhost:3000
```

## Sem Docker

### Configurando ENV

``` bash
export MONGO_DATABASE_NAME="databr"
export MONGO_URL="mongodb://dev"

export API_ROOT="http://localhost:3002"
export PORT=3002
export ENV="development"

export STATUSPAGEIO_ENABLE="false"

export PRIVATE_KEY="qweasdzxc"

```

### Rodando

``` bash
$ go get github.com/tools/godep
$ go get github.com/pilu/fresh
$ godep restore
$ fresh
```
