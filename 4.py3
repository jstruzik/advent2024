#!/usr/bin/python3
import re

def import_file(filename: str) -> str:
    with open(filename, 'r') as file:
        return file.read()
    
def create_matrix(data: str) -> list[list[str]]:
    char_matrix = [[]]
    index = 0
    for char in data:
        if char == '\n':
            char_matrix.append([])
            index += 1
        else: char_matrix[index].append(char)
    return char_matrix

def solve_1(data: str) -> int:
    result = 0
    char_matrix = create_matrix(data)  
    height = len(char_matrix)

    for row_idx, row in enumerate(char_matrix):
        width = len(row)
        for col_idx, col in enumerate(row):
            if col == 'X':
                # Traverse!!!
                trav_right = col_idx+3 < width and (row[col_idx] + row[col_idx+1] + row[col_idx+2] + row[col_idx+3]) == 'XMAS'
                trav_left = col_idx-3 >= 0 and (row[col_idx] + row[col_idx-1] + row[col_idx-2] + row[col_idx-3]) == 'XMAS'
                trav_up = row_idx-3 >= 0 and (char_matrix[row_idx][col_idx] + char_matrix[row_idx-1][col_idx] + char_matrix[row_idx-2][col_idx] + char_matrix[row_idx-3][col_idx]) == 'XMAS'
                trav_down = row_idx+3 < height and (char_matrix[row_idx][col_idx] + char_matrix[row_idx+1][col_idx] + char_matrix[row_idx+2][col_idx] + char_matrix[row_idx+3][col_idx]) == 'XMAS'
                trav_up_right = col_idx+3 < width and row_idx-3 >= 0 and (char_matrix[row_idx][col_idx] + char_matrix[row_idx-1][col_idx+1] + char_matrix[row_idx-2][col_idx+2] + char_matrix[row_idx-3][col_idx+3]) == 'XMAS'
                trav_up_left = row_idx-3 >= 0 and col_idx-3 >= 0 and (char_matrix[row_idx][col_idx] + char_matrix[row_idx-1][col_idx-1] + char_matrix[row_idx-2][col_idx-2] + char_matrix[row_idx-3][col_idx-3]) == 'XMAS'
                trav_down_right = row_idx+3 < height and col_idx+3 < width and (char_matrix[row_idx][col_idx] + char_matrix[row_idx+1][col_idx+1] + char_matrix[row_idx+2][col_idx+2] + char_matrix[row_idx+3][col_idx+3]) == 'XMAS'
                trav_down_left = row_idx+3 < height and col_idx-3 >= 0 and (char_matrix[row_idx][col_idx] + char_matrix[row_idx+1][col_idx-1] + char_matrix[row_idx+2][col_idx-2] + char_matrix[row_idx+3][col_idx-3]) == 'XMAS'
                result += sum([trav_right, trav_left, trav_up, trav_down, trav_up_right, trav_up_left, trav_down_right, trav_down_left])

    return result

def solve_2(data: str) -> int:
    result = 0
    char_matrix = create_matrix(data)
    height = len(char_matrix)
    mas_ord = ord('M') + ord('S')

    for row_idx, row in enumerate(char_matrix):
        width = len(row)
        for col_idx, col in enumerate(row):
            if col == 'A':
                # Traverse!!!
                trav_diag_left = row_idx - 1 >= 0 and col_idx + 1 < width and row_idx + 1 < height and col_idx - 1 >= 0 and (ord(char_matrix[row_idx-1][col_idx+1]) + ord(char_matrix[row_idx+1][col_idx-1])) == mas_ord
                trav_diag_right = row_idx - 1 >= 0 and col_idx - 1 >= 0 and row_idx + 1 < height and col_idx + 1 < width and (ord(char_matrix[row_idx-1][col_idx-1]) + ord(char_matrix[row_idx+1][col_idx+1])) == mas_ord
                
                if trav_diag_left and trav_diag_right:
                    result += 1
    
    return result

if __name__ == "__main__":
    data = import_file('./resources/4.txt')
    result_1 = solve_1(data)
    result_2 = solve_2(data)
    print(result_1)
    print(result_2)