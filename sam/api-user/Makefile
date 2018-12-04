build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/userGet user/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/userPut user/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/userList user/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/userPost user/post/post.go
    env GOOS=linux go build -ldflags="-s -w" -o bin/userDelete user/delete/delete.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupGet group/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupPut group/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupList group/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/groupPost group/post/post.go
    env GOOS=linux go build -ldflags="-s -w" -o bin/groupDelete group/delete/delete.go    	

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
