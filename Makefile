# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go install go.uber.org/mock/mockgen@latest
	@go mod download

# Run tests
test:
	@echo "Running tests..."
	@go test -v -count=1 ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# Lint the code
lint:
	@echo "Linting..."
	@golangci-lint run
