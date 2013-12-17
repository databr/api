# Social Câmara
## A rede social preferida dos deputados da cârama!

## API

Esse repositório consiste na API do projeto,([caso esteja interessado no frontend](https://github.com/dukex/socialcamara-site)), é responsabilidade dessa API prover todos os dados que o frontend precisa, isso inclui códigos do bot que busca os dados e o servidor para responder as chamadas do frontend


Bot

* [/data](https://github.com/dukex/socialcamara-api/tree/master/data) Tem os parsers que buscam dados no site da câmara, esses arquivos tem a responsabilidade apenas de trazer os dados, sem intereção com o banco de dados.
* [/lib/services](https://github.com/dukex/socialcamara-api/tree/master/lib/services) Esse Services utilizam os dados que os parsers trazem e salvam no banco de dados correto

Server

* [/api.rb](https://github.com/dukex/socialcamara-api/blob/master/api.rb) Contem o código do servidor web, para fazer a api utilizamos o [Grape](https://github.com/intridea/grape), esse arquivo busca no banco de dados e responde

* [/lib/models/](https://github.com/dukex/socialcamara-api/tree/master/lib/models) Os modelos, usando [activerecord](http://api.rubyonrails.org/classes/ActiveRecord/Base.html)


### Instalação

```
  $ git clone https://github.com/dukex/socialcamara-api
  $ cd socialcamara-api
  $ bundle install
  $ rake db:create db:migrate db:test:prepare
  $ bundle exec rspec spec
```

### Bots

Para rodar os bots é bem simples, rode:

```
  rake data:all
```

Para mais informações veja o arquivo [Rakefile](https://github.com/dukex/socialcamara-api/blob/master/Rakefile#L7-L33)


### Contribua

1. Fork o projeto
2. Crie um feature-branch ($ git checkout -b myfeaturebranch master)
3. Trabalhe com commits pequenos e objetivos
4. Atualiza esse branch no github ($ git push origin myfeaturebranch)
5. Mande um pull-request
