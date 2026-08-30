.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: deps
deps:
	go mod tidy
