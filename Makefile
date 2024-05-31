run: build
	@./bin/twitter

build:
	@go build -o bin/twitter main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

