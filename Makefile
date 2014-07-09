test:
	@cd lib && go test

fmt:
	@go fmt ./...

build:
	@go build -o resize.out resize.go

.PHONY: test fmt
