test:
	cd lib && go test

fmt:
	go fmt ./...

.PHONY: test fmt
