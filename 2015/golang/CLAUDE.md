# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is an Advent of Code 2015 solution repository written in Go. The project follows a structured approach where each day's challenge is implemented as a separate module with its own solution functions and test cases.

## Repository Structure

```
├── cmd/dayN/           # Executable main files for each day
├── days/dayN/          # Solution packages for each day
├── input/              # Puzzle input files (day1.txt, day2.txt, etc.)
└── utils/              # Shared utilities (input reading)
```

**Key Architecture Patterns:**
- Each day has its own package in `days/dayN/` with `Part1()` and `Part2()` functions
- Main executables in `cmd/dayN/` call the corresponding day's functions
- Input files are stored in `input/dayN.txt` and read via `utils.ReadInput(day)`
- All solutions follow the signature: `func PartN(input string) int`

## Common Commands

**Run a specific day's solution:**
```bash
go run cmd/day1/main.go
go run cmd/day5/main.go
```

**Run tests for a specific day:**
```bash
go test ./days/day1
go test ./days/day5
```

**Run all tests:**
```bash
go test ./...
```

**Run a single test function:**
```bash
go test -run TestPart1 ./days/day5
go test -run TestPart2 ./days/day5
```

**Run tests with verbose output:**
```bash
go test -v ./days/day5
```

**Build all executables:**
```bash
go build ./cmd/...
```

## Important Development Notes

- **NEVER edit solution files directly** - The user solves the puzzles themselves
- Solutions are in `days/dayN/dayN.go` with `Part1()` and `Part2()` functions
- Test files use table-driven testing pattern with individual test cases
- Input reading is handled by `utils.ReadInput(day)` which automatically finds project root
- All puzzle inputs are stored as `input/dayN.txt` files
- Tests use `github.com/stretchr/testify/assert` for assertions

## Module Information

- Module: `github.com/hagenek/advent/2015/golang`
- Go version: 1.25.3
- Main dependency: `github.com/stretchr/testify v1.11.1`

## Working with Tests

Test files follow the pattern `days/dayN/dayN_test.go` and use table-driven tests with individual test cases for each example from the problem description. When setting up tests, use the structure from existing day tests (like day3) with separate test cases for each input example.