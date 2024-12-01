from python import Python
from utils.static_tuple import StaticTuple

fn calculate_floor(instructions: String) raises -> Int:
    """
    Calculate the final floor based on parentheses instructions.
    '(' means go up one floor
    ')' means go down one floor
    """

    var floor: Int

    for c in range(len(instructions)):
        instruction = instructions[c]
        if instruction == '(':
            floor += 1
        elif instruction == ')':
            floor -= 1
    
    return floor

fn test_cases() raises:
    """Run test cases to verify the implementation."""
    var test_inputs = StaticTuple[5]("(())", "()()", "(((", "(()(()(", "))(((((" )
    var expected_outputs = StaticTuple[5](0, 0, 3, 3, 3)
    
    for i in range(len(test_inputs)):
        var result = calculate_floor(test_inputs[i])
        print("Input:", test_inputs[i], "Expected:", expected_outputs[i], "Got:", result)
        assert result == expected_outputs[i]
    
    # Test basement cases
    let basement_inputs = StaticTuple[4]("())", "))(", ")))", ")())())")
    let basement_outputs = StaticTuple[4](-1, -1, -3, -3)
    
    for i in range(len(basement_inputs)):
        let result = calculate_floor(basement_inputs[i])
        print("Input:", basement_inputs[i], "Expected:", basement_outputs[i], "Got:", result)
        assert result == basement_outputs[i]

fn main() raises:
    # First run the test cases
    print("Running test cases...")
    test_cases()
    print("\nAll test cases passed!")
    
    # For solving the actual puzzle, you would read the input file
    # Note: In Mojo, we'll use Python's built-in file operations
    let python = Python.import_module("builtins")
    let file = python.open("input.txt", "r")
    let puzzle_input = String(file.read().strip())
    file.close()
    
    let final_floor = calculate_floor(puzzle_input)
    print("\nSanta's final floor:", final_floor)