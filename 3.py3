#!/usr/bin/python3
import re

def import_file(filename: str) -> str:
    with open(filename, 'r') as file:
        return file.read()

def solve_1(data: str) -> int:
    result = 0
    for match in re.finditer('mul\((\d{1,3})\,(\d{1,3})\)', data):
        result += int(match.group(1)) * int(match.group(2))
    
    return result

def solve_2(data: str) -> int:
    result = 0
    skip_match = False
    for match in re.finditer('(.*?)(do\(\)|don\'t\(\))|(.*)', data, re.S):
        if False == skip_match:
            to_parse = match.group(3) if match.group(3) is not None else match.group(1)
            for inner_match in re.finditer('mul\((\d{1,3})\,(\d{1,3})\)', to_parse):
                result += int(inner_match.group(1)) * int(inner_match.group(2))
        skip_match = match.group(2) == 'don\'t()'
    
    return result

if __name__ == "__main__":
    data = import_file('./resources/3.txt')
    result_1 = solve_1(data)
    result_2 = solve_2(data)
    print(result_1)
    print(result_2)