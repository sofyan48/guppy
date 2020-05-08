build-linux:
	env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/cli/guppy guppy-cli/main.go

build-mac:
	env GOOS=darwin go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/cli/guppy guppy-cli/main.go

clean:
	rm -rf ./bin

deploy: clean build-$(os)