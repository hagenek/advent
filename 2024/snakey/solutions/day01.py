import os

def read_input(day: int) -> str:
    file_path = os.path.join(os.path.dirname(__file__), '..', 'inputs', f'day{day:02}.txt')
    with open(file_path, 'r') as file:
        return file.read().strip()

def parse_and_sort_columns(data: str):
    lines = []
    for line in data.split("\n"):
        x, y = map(int, line.split())
        lines.append((x, y))

    col_one = []
    col_two = []
    for x1, y1 in lines:
        col_one.append(x1)
        col_two.append(y1)

    col_one.sort()
    col_two.sort()

    return col_one, col_two

def solve_part1(data: str):
    col_one, col_two = parse_and_sort_columns(data)

    difference = 0
    for i in range(len(col_one)):
        if col_one[i] == col_two[i]:
            continue
        else:
            difference += abs(col_one[i] - col_two[i])

    return difference

def solve_part2(data: str):
    col_one, col_two = parse_and_sort_columns(data)

    frequency_dict = {}
    for n in col_two:
        if n in frequency_dict:
            frequency_dict[n] += 1
        else:
            frequency_dict[n] = 1

    sim_index = 0
    for n in col_one:
        if n in frequency_dict:
            sim_index += n * frequency_dict[n]

    return sim_index

demo_data_part_one = """3   4
4   3
2   5
1   3
3   9
3   3"""

demo_data_part_two = """3   4
4   3
2   5
1   3
3   9
3   3"""

solved1 = solve_part1(read_input(1))
solved2 = solve_part2(read_input(1))
print(solved1)
print(solved2)
