"""Tests for Day 1."""

import pytest
from days.day01 import part1, part2


class TestDay01:
    """Test cases for Day 1."""

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
        input_data = read_input(1)
        result = part1(input_data)
        assert result == 0  # Update with expected result

    @pytest.mark.skip(reason="Not implemented yet")
    def test_part2_actual(self):
        """Test part 2 with actual input."""
        from utils.input_reader import read_input
        input_data = read_input(1)
        result = part2(input_data)
        assert result == 0  # Update with expected result
