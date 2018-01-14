.PHONY: build clean

default: build

OUTPUT_PATH=build/settings

clean:
	rm -rf build && mkdir build && touch build/.gitkeep

build: clean
	go build -o $(OUTPUT_PATH)

run:
	go run $(shell find . -name "*.go" | grep -v _test.go)

goRun:
	$(OUTPUT_PATH)