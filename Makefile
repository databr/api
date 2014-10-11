test:
	cd service && ginkgo -r --randomizeAllSpecs -cover -race

coverage: test
	go tool cover -html=service/service.coverprofile

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
