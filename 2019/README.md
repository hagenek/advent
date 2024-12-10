# Advent of Code 2019 - Go Solutions

This repository contains my solutions for the [Advent of Code 2019](https://adventofcode.com/2019) challenges implemented in Go.

## Project Structure

```
.
├── main.go              # Main program to run solutions
├── days/               # Solution implementations
│   └── day01.go        # Day 1 solution
├── inputs/             # Input files
│   └── day01.txt       # Day 1 input
└── tests/              # Test files
    └── days/
        └── day01_test.go
```

## Setup

1. Clone the repository
2. Make sure you have Go installed
3. Initialize the module:
```bash
go mod init aoc2019-golang
```

## Usage

Run a specific day's solution:
```bash
go run main.go -day 1  # Runs day 1 solution
```

The program expects input files to be placed in the `inputs` directory, named as `day01.txt`, `day02.txt`, etc.

## Testing

Run all tests:
```bash
go test ./tests/...
```

## Adding New Solutions

1. Add your solution implementation in `days/dayXX.go`
2. Create corresponding test in `tests/days/dayXX_test.go`
3. Add your input file in `inputs/dayXX.txt`
4. Add a new case in `main.go`'s switch statement

## Implementation Notes

- Each day's solution should implement a function of the form `DayX(input string) (int, int)`
- Input files should use Unix-style line endings (\n)
- Empty lines at the end of input files are automatically trimmed
