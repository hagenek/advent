from solutions.day01 import solve_part1, solve_part2
from utils.input_reader import read_input

def run_day(day: int):
    data = read_input(day)
    solution_module = __import__(f'solutions.day{day:02}', fromlist=['solve_part1', 'solve_part2'])
    print(f"Day {day} Solutions:")
    print("Part 1:", solution_module.solve_part1(data))
    print("Part 2:", solution_module.solve_part2(data))

if __name__ == "__main__":
    import sys
    day = int(sys.argv[1]) if len(sys.argv) > 1 else 1
    run_day(day)
