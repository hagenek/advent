.PHONY: day test run build clean help

# Generate scaffolding for a new day
# Usage: make day DAY=15 or make 15
day:
ifndef DAY
	@echo "Error: DAY parameter is required"
	@echo "Usage: make day DAY=15 or make 15"
	@exit 1
endif
	@go run cmd/scaffold/main.go $(DAY)

# Generate scaffolding for days 1-25 (individual targets)
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25:
	@go run cmd/scaffold/main.go $@

# Run tests for a specific day
# Usage: make test DAY=5
test:
ifdef DAY
	@go test -v ./days/day$(DAY)
else
	@go test ./...
endif

# Run solution for a specific day
# Usage: make run DAY=5
run:
ifndef DAY
	@echo "Error: DAY parameter is required"
	@echo "Usage: make run DAY=5"
	@exit 1
endif
	@go run cmd/day$(DAY)/main.go

# Build all executables
build:
	@go build ./cmd/...

# Clean build artifacts
clean:
	@go clean ./...

# Show help
help:
	@echo "Available targets:"
	@echo "  day DAY=N    - Generate scaffolding for day N"
	@echo "  test [DAY=N] - Run tests (all or for specific day)"
	@echo "  run DAY=N    - Run solution for day N"
	@echo "  build        - Build all executables"
	@echo "  clean        - Clean build artifacts"
	@echo ""
	@echo "Examples:"
	@echo "  make 15            - Generate scaffolding for day 15"
	@echo "  make day DAY=15    - Generate scaffolding for day 15 (alternative)"
	@echo "  make test DAY=5    - Run tests for day 5"
	@echo "  make test          - Run all tests"
	@echo "  make run DAY=5     - Run solution for day 5"