#!/usr/bin/env python3
"""Run solution for Day 1."""

from days.day01 import part1, part2
from utils.input_reader import read_input


def main():
    """Run both parts of the solution."""
    input_data = read_input(1)

    result1 = part1(input_data)
    print(f"Part 1: {result1}")

    result2 = part2(input_data)
    print(f"Part 2: {result2}")


if __name__ == "__main__":
    main()
