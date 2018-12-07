# Binary name
BINARY=app
# Builds the projectbuild:
build:
		go build -o ${BINARY}
		go test -v

# Installs our project: copies binaries
install:
		go install
release:
#		create out folder for package tar.gz files
		mkdir -p out
		# Clean
		go clean
		rm -rf *.gz

		# Build for mac
		go build -ldflags="-s -w" -o bin/app app.go
		tar czvf out/app-mac64-${VERSION}.tar.gz ./bin/app
		# Build for linux
		go clean
		env GOOS=linux go build -ldflags="-s -w" -o bin/app app.go
		tar czvf out/app-linux64-${VERSION}.tar.gz ./bin/app
		# Build for win
		go clean
		env GOOS=windows go build -ldflags="-s -w" -o bin/app.exe app.go
		tar czvf out/app-win64-${VERSION}.tar.gz ./bin/app.exe
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean

.PHONY:  clean build
