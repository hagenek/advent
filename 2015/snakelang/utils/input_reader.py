import os

def read_input(day: int) -> str:
    file_path = os.path.join(os.path.dirname(__file__), '..', 'inputs', f'day{day:02}.txt')
    with open(file_path, 'r') as file:
        return file.read().strip()
