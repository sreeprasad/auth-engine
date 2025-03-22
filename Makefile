.PHONY: build clean test coverage bench fmt lint

GOFLAGS := -v

build:
	go build $(GOFLAGS) -o bin/auth-engine

clean:
	rm -rf bin/
	go clean

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

bench:
	go test ./... -bench=. -benchmem

fmt:
	go fmt ./...

lint:
	go vet ./...
	@if command -v golint > /dev/null; then \
		golint ./...; \
	else \
		echo "golint not installed. Run: go install golang.org/x/lint/golint@latest"; \
	fi

check: fmt lint test
