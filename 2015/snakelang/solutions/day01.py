import os

def read_input(day: int) -> str:
    file_path = os.path.join(os.path.dirname(__file__), '..', 'inputs', f'day{day:02}.txt')
    with open(file_path, 'r') as file:
        return file.read().strip()

def str_to_int(s: str) -> int:
    if (s == '('):
        return 1
    return -1

def solve_part1(data: str):
    return data.count("(") - data.count(")")

def solve_part2(data: str):
    floor = 0
    for index, char in enumerate(data, 1):  
        if char == "(":
            floor += 1
        elif char == ")":
            floor -= 1
            if floor < 0:
                return index 

solved1 = solve_part1(read_input(1))
solved2 = solve_part2(read_input(1))
print(solved1)
print(solved2)