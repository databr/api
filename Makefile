server:
	go get && fresh

deps:
	go get github.com/gin-gonic/gin
	go get github.com/pilu/fresh
	go get gopkg.in/mgo.v2
	go get github.com/camarabook/go-popolo
	go get github.com/fiam/gounidecode/unidecode
	go get

deploy:
	godep save
	git add --all Godeps/
	git commit -m "updated deps"
	git push heroku master
	
