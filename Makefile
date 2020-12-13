.DEFAULT_GOAL := build

build:
	@go build -o dist/rpn cmd/main.go

test:
	@go test github.com/matthewaj/rpncaclearlc/rpn/ -v

run:
	@go run cmd/main.go
