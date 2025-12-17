.PHONY: help install test format lint type-check clean day run-day

help:
	@echo "Available commands:"
	@echo "  install      - Install dependencies"
	@echo "  test         - Run all tests"
	@echo "  format       - Format code with black"
	@echo "  lint         - Lint code with flake8"
	@echo "  type-check   - Type check with mypy"
	@echo "  clean        - Remove cache files"
	@echo "  day DAY=N    - Create scaffolding for day N"
	@echo "  run-day DAY=N - Run solution for day N"

install:
	pip install -r requirements.txt

test:
	pytest tests/

test-day:
	@if [ -z "$(DAY)" ]; then \
		echo "Please specify DAY=N"; \
		exit 1; \
	fi
	pytest tests/test_day$(shell printf "%02d" $(DAY)).py -v

format:
	black .

lint:
	flake8 . --max-line-length=100 --ignore=E203,W503

type-check:
	mypy . --ignore-missing-imports

clean:
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type d -name ".pytest_cache" -exec rm -rf {} +
	find . -type d -name ".mypy_cache" -exec rm -rf {} +

day:
	@if [ -z "$(DAY)" ]; then \
		echo "Please specify DAY=N"; \
		exit 1; \
	fi
	python3 scaffold.py $(DAY)

run-day:
	@if [ -z "$(DAY)" ]; then \
		echo "Please specify DAY=N"; \
		exit 1; \
	fi
	python3 day$(shell printf "%02d" $(DAY)).py

# Shortcuts for individual days
$(filter-out day,$(shell seq 1 25)):
	@python3 scaffold.py $@