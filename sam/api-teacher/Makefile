build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/teacherGet teacher/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/teacherPut teacher/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/teacherList teacher/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/teacherPost teacher/post/post.go
    env GOOS=linux go build -ldflags="-s -w" -o bin/teacherDelete teacher/delete/delete.go

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
