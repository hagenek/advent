#!/usr/bin/env python3
"""Scaffold tool for creating new day solutions."""

import os
import sys
from pathlib import Path

DAY_TEMPLATE = '''"""Day {day} solution for Advent of Code 2025."""


def parse(input_data: str):
    """Parse the input data."""
    return input_data.strip().split('\\n')


def part1(input_data: str) -> int:
    """Solve part 1."""
    # TODO: Implement part 1
    return 0


def part2(input_data: str) -> int:
    """Solve part 2."""
    # TODO: Implement part 2
    return 0


if __name__ == "__main__":
    from utils.input_reader import read_input

    input_data = read_input({day})
    print(f"Part 1: {{part1(input_data)}}")
    print(f"Part 2: {{part2(input_data)}}")
'''

TEST_TEMPLATE = '''"""Tests for Day {day}."""

import pytest
from days.day{day:02d} import part1, part2


class TestDay{day:02d}:
    """Test cases for Day {day}."""

    @pytest.fixture
    def example_input(self):
        """Example input for testing."""
        return """test input here"""

    def test_part1_example(self, example_input):
        """Test part 1 with example input."""
        assert part1(example_input) == 0  # Update with expected result

    def test_part2_example(self, example_input):
        """Test part 2 with example input."""
        assert part2(example_input) == 0  # Update with expected result

    @pytest.mark.skip(reason="Not implemented yet")
    def test_part1_actual(self):
        """Test part 1 with actual input."""
        from utils.input_reader import read_input
        input_data = read_input({day})
        result = part1(input_data)
        assert result == 0  # Update with expected result

    @pytest.mark.skip(reason="Not implemented yet")
    def test_part2_actual(self):
        """Test part 2 with actual input."""
        from utils.input_reader import read_input
        input_data = read_input({day})
        result = part2(input_data)
        assert result == 0  # Update with expected result
'''

MAIN_TEMPLATE = '''#!/usr/bin/env python3
"""Run solution for Day {day}."""

from days.day{day:02d} import part1, part2
from utils.input_reader import read_input


def main():
    """Run both parts of the solution."""
    input_data = read_input({day})

    result1 = part1(input_data)
    print(f"Part 1: {{result1}}")

    result2 = part2(input_data)
    print(f"Part 2: {{result2}}")


if __name__ == "__main__":
    main()
'''


def create_scaffolding(day: int) -> None:
    """Create scaffolding for a new day."""
    base_dir = Path(__file__).parent

    # Create directories if they don't exist
    (base_dir / "days").mkdir(exist_ok=True)
    (base_dir / "input").mkdir(exist_ok=True)
    (base_dir / "tests").mkdir(exist_ok=True)

    # Create solution file
    day_file = base_dir / "days" / f"day{day:02d}.py"
    if day_file.exists():
        print(f"Day {day} already exists. Skipping...")
        return

    day_file.write_text(DAY_TEMPLATE.format(day=day))
    print(f"Created: {day_file}")

    # Create test file
    test_file = base_dir / "tests" / f"test_day{day:02d}.py"
    test_file.write_text(TEST_TEMPLATE.format(day=day))
    print(f"Created: {test_file}")

    # Create main runner
    main_file = base_dir / f"day{day:02d}.py"
    main_file.write_text(MAIN_TEMPLATE.format(day=day))
    main_file.chmod(0o755)
    print(f"Created: {main_file}")

    # Create empty input file
    input_file = base_dir / "input" / f"day{day:02d}.txt"
    input_file.touch()
    print(f"Created: {input_file}")

    # Create instructions file
    instructions_file = base_dir / "days" / f"day{day:02d}_instructions.md"
    instructions_file.touch()
    print(f"Created: {instructions_file}")

    print(f"\\nScaffolding for Day {day} created successfully!")
    print("\\nTo get started:")
    print(f"  1. Add your puzzle input to input/day{day:02d}.txt")
    print(f"  2. Implement solutions in days/day{day:02d}.py")
    print(f"  3. Add test cases to tests/test_day{day:02d}.py")
    print(f"  4. Run tests: pytest tests/test_day{day:02d}.py")
    print(f"  5. Run solution: python day{day:02d}.py")


def main():
    """Main entry point."""
    if len(sys.argv) != 2:
        print(f"Usage: {sys.argv[0]} <day>")
        print(f"Example: {sys.argv[0]} 12")
        sys.exit(1)

    try:
        day = int(sys.argv[1])
        if not 1 <= day <= 25:
            raise ValueError("Day must be between 1 and 25")
    except ValueError as e:
        print(f"Error: {e}")
        sys.exit(1)

    create_scaffolding(day)


if __name__ == "__main__":
    main()