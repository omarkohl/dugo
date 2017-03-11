all:
	@printf "build                  Build the dugo binary\n"
	@printf "clean                  Clean build files\n"
	@printf "test-integration       Run integration tests\n"

build:
	go build dugo.go

clean:
	rm -f dugo

test-integration: build
	PATH=${PATH}:. go test
