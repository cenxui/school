build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneList zone/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneGet zone/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zonePut zone/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zonePost zone/post/post.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneDelete zone/delete/delete.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/positionList position/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/positionGet position/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/positionPut position/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/positionPost position/post/post.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/positionDelete position/delete/delete.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneTemplateList zone-template/list/list.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneTemplateGet zone-template/get/get.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneTemplatePut zone-template/put/put.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneTemplatePost zone-template/post/post.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/zoneTemplateDelete zone-template/delete/delete.go
.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
