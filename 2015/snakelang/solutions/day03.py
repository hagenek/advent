demo_data_3 = "^v"
demo_data_3 = "^>v<"
demo_data_11 = "^v^v^v^v^v "

def char_to_coordinate(char, x, y):
    match char:
        case "^":
            return x, y + 1  # Up
        case "v":
            return x, y - 1  # Down
        case ">":
            return x + 1, y  # Right
        case "<":
            return x - 1, y  # Left
        case _:
            return x, y      # Default case for any other character

def part1(data):
    visited = {(0,0)}
    for char in data:
        new_coord = char_to_coordinate(char)
        visited.add(new_coord)
    
    return len(visited)

# Test with demo data
print(f"Demo 1 (^v): {part1('^v')}")           # Should visit 2 houses
print(f"Demo 2 (^>v<): {part1('^>v<')}")       # Should visit 4 houses
print(f"Demo 3 (^v^v^v^v^v): {part1('^v^v^v^v^v')}")  # Should visit 2 houses