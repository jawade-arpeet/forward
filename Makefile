.PHONY: run tidy fmt

run:
	go run cmd/main.go

tidy:
	go mod tidy

fmt:
	gofmt -w .