all:
	@printf "build                  Build the dugo binary\n"
	@printf "clean                  Clean build files\n"
	@printf "test-integration       Run integration tests\n"
	@printf "test-unit              Run unit tests\n"

build:
	go build dugo.go tree.go

clean:
	rm -f dugo

test-integration: build
	PATH=${PATH}:. go test

test-unit:
	go test -short
