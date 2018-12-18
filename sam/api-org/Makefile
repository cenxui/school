build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/orgGet org/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/orgList org/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/orgPut org/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/orgPost org/post/post.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/orgDelete org/delete/delete.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
