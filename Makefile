test:
	cd service &&	ginkgo -r --randomizeAllSpecs -cover
	cd service && gover
	cd service && goveralls -service travis-ci -coverprofile=service.coverprofile $COVERALLS_TOKEN

server:
	go get && fresh

deps:
	go get github.com/tools/godep
	godep restore
	go get

save_deps:
	godep save
	git add --all Godeps/
	git commit -m "updated deps"

deploy:
	make save_deps -i
	git push heroku master
	git push origin master
