# Advent of Code 2025 - Python Solutions

## Setup

1. Install dependencies:
```bash
make install
```

## Usage

### Create a new day
```bash
make day DAY=1
# or simply
make 1
```

### Run a solution
```bash
make run-day DAY=1
# or directly
python3 day01.py
```

### Run tests
```bash
# All tests
make test

# Specific day
make test-day DAY=1
# or directly
pytest tests/test_day01.py
```

### Code quality
```bash
make format     # Format with black
make lint       # Lint with flake8
make type-check # Type check with mypy
```

## Project Structure
```
2025/python/
├── days/           # Solution modules
│   ├── day01.py
│   └── ...
├── input/          # Puzzle inputs
│   ├── day01.txt
│   └── ...
├── tests/          # Test files
│   ├── test_day01.py
│   └── ...
├── utils/          # Helper utilities
│   ├── input_reader.py
│   └── helpers.py
├── scaffold.py     # Scaffolding generator
├── requirements.txt
├── Makefile
└── day01.py        # Day runners
```