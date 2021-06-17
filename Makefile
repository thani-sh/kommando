.PHONY: build
build:
	go build -o ./build/kommando ./...
	cp build/kommando ./example/

.PHONY: release
release:
	GOOS=linux  GOARCH=amd64 go build -ldflags="-s -w" -o ./build/kommando.linux ./...
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./build/kommando.macos ./...
	cp build/* ./example/
