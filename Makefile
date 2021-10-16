.PHONY: build
build:
	go build -o ./build/kommando ./cmd/kommando/...
	cp build/kommando ./example/

.PHONY: clean
clean:
	rm -rf build
	rm example/kommando*

.PHONY: test
test:
	alias
