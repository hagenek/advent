from utils.input_reader import read_input

def solve_part1(data: str) -> int:
    # Implement your part 1 logic here
    return sum(map(int, data.split()))

def solve_part2(data: str) -> int:
    # Implement your part 2 logic here
    return max(map(int, data.split()))

if __name__ == "__main__":
    input_data = read_input(1)
    print("Part 1:", solve_part1(input_data))
    print("Part 2:", solve_part2(input_data))
