all:
	@printf "build                  Build the dugo binary\n"
	@printf "test-integration       Run integration tests\n"

build:
	go build dugo.go

test-integration: build
	PATH=${PATH}:. go test
