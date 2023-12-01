import re

def sum_calibration_values_part1(lines):
    def extract_digits(line):
        digits = re.findall(r'\d', line)
        if len(digits) >= 2:
            return int(digits[0] + digits[-1])
        elif len(digits) == 1:
            return int(digits[0] * 2)  # Double the digit if it's the only one
        return 0

    return sum(extract_digits(line) for line in lines)

def sum_calibration_values_part2(lines):
    number_words = {
        'one': '1', 'two': '2', 'three': '3', 'four': '4', 'five': '5',
        'six': '6', 'seven': '7', 'eight': '8', 'nine': '9'
    }

    def extract_digits(line):
        for word, digit in number_words.items():
            line = line.replace(word, digit)
        digits = re.findall(r'\d', line)
        if len(digits) >= 2:
            return int(digits[0] + digits[-1])
        elif len(digits) == 1:
            return int(digits[0] * 2)  # Double the digit if it's the only one
        return 0

    return sum(extract_digits(line) for line in lines)

example_input_part1 = [   
    "1abc2",
    "pqr3stu8vwx",
    "a1b2c3d4e5f",
    "treb7uchet"
]

example_input_part2 = [
    "two1nine",
    "eightwothree",
    "abcone2threexyz",
    "xtwone3four",
    "4nineeightseven2",
    "zoneight234",
    "7pqrstsixteen"
]

print(sum_calibration_values_part1(example_input_part1)) # Output: 142
print(sum_calibration_values_part2(example_input_part2)) # Output: 211
