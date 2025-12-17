"""Input reader utility for Advent of Code 2025."""

from pathlib import Path


def read_input(day: int) -> str:
    """Read input for a given day.

    Args:
        day: The day number (1-25)

    Returns:
        The input data as a string
    """
    input_file = Path(__file__).parent.parent / "input" / f"day{day:02d}.txt"

    if not input_file.exists():
        raise FileNotFoundError(f"Input file for day {day} not found at {input_file}")

    return input_file.read_text().strip()


def read_lines(day: int) -> list[str]:
    """Read input lines for a given day.

    Args:
        day: The day number (1-25)

    Returns:
        The input data as a list of lines
    """
    return read_input(day).split('\n')