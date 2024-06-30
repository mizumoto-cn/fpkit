.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: ut
ut:
	@echo "Running unit tests with race detector..."
	@go test -tags=goexperiment.arenas -race ./...

.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run 