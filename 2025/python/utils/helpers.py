"""Helper functions for Advent of Code 2025."""

import re
from typing import Any, Callable, Iterator, TypeVar

T = TypeVar('T')


def extract_numbers(text: str) -> list[int]:
    """Extract all numbers from a string.

    Args:
        text: Input string

    Returns:
        List of integers found in the string
    """
    return [int(x) for x in re.findall(r'-?\d+', text)]


def chunk(lst: list[T], size: int) -> Iterator[list[T]]:
    """Split a list into chunks of specified size.

    Args:
        lst: Input list
        size: Size of each chunk

    Yields:
        Chunks of the specified size
    """
    for i in range(0, len(lst), size):
        yield lst[i:i + size]


def sliding_window(lst: list[T], size: int) -> Iterator[tuple[T, ...]]:
    """Generate sliding windows of specified size.

    Args:
        lst: Input list
        size: Size of the window

    Yields:
        Sliding windows as tuples
    """
    for i in range(len(lst) - size + 1):
        yield tuple(lst[i:i + size])


def memoize(func: Callable) -> Callable:
    """Simple memoization decorator.

    Args:
        func: Function to memoize

    Returns:
        Memoized function
    """
    cache = {}

    def wrapper(*args, **kwargs):
        key = (args, frozenset(kwargs.items()) if kwargs else ())
        if key not in cache:
            cache[key] = func(*args, **kwargs)
        return cache[key]

    wrapper.__name__ = func.__name__
    wrapper.__doc__ = func.__doc__
    return wrapper


def manhattan_distance(p1: tuple[int, int], p2: tuple[int, int]) -> int:
    """Calculate Manhattan distance between two points.

    Args:
        p1: First point (x, y)
        p2: Second point (x, y)

    Returns:
        Manhattan distance
    """
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])


def neighbors_4(x: int, y: int) -> list[tuple[int, int]]:
    """Get 4-connected neighbors (up, down, left, right).

    Args:
        x: X coordinate
        y: Y coordinate

    Returns:
        List of neighbor coordinates
    """
    return [(x-1, y), (x+1, y), (x, y-1), (x, y+1)]


def neighbors_8(x: int, y: int) -> list[tuple[int, int]]:
    """Get 8-connected neighbors (including diagonals).

    Args:
        x: X coordinate
        y: Y coordinate

    Returns:
        List of neighbor coordinates
    """
    return [
        (x-1, y-1), (x-1, y), (x-1, y+1),
        (x, y-1),             (x, y+1),
        (x+1, y-1), (x+1, y), (x+1, y+1)
    ]