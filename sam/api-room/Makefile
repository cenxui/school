build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/roomGet room/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/roomPut room/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/roomList room/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/roomPost room/post/post.go
    env GOOS=linux go build -ldflags="-s -w" -o bin/roomDelete room/delete/delete.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
