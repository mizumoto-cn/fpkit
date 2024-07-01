.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: race
ut:
	@echo "Running unit tests with race detector..."
	@go test -tags=goexperiment.arenas -race ./...

.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run 

.PHONY: benchmark
benchmark:
	@echo "Running benchmarks..."
	@go test -bench=. ./...

.PHONY: coverage
coverage:
	@echo "Running coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out