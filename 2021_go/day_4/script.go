package main

import (
	"strings"
	"os"
	"strconv"
	"fmt"
	"io"
	"bufio"
)

func ReadIntsFromStringArray(str_arr []string) ([]int, error) {
    var result []int
    for _, s := range str_arr  {
        x, err := strconv.Atoi(s)
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, nil
}

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    // ScanRunes to split on each utf-8 rune.
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {

        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func Sum(input []int) int {
    sum := 0
    for _, i := range input {
        sum += i
    }
    return sum
}

func SumUnmarkedBoard(board [][]int) int {
	sum := 0
	for _, row := range board {
		for _, number := range row {
			if number != -1 {
				sum += number
			}
		}
	}
	return sum
}

func Transpose(x [][]int) [][]int {
	out := make([][]int, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func CheckForBingo(board [][]int) bool {
	for _, row := range board {
		if Sum(row) == -5 {
			return true
		}
	}
	transposed_board := Transpose(board)
	for _, col := range transposed_board {
		if Sum(col) == -5 {
			return true
		}
	}
	return false
}

func main() {
	// Initialization.
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input),"\n")
	number_strings := strings.Split(lines[0], ",")
	bingo_numbers, _ := ReadIntsFromStringArray(number_strings)
	fmt.Println(bingo_numbers)

	// Populate board map, could refactor into its own function.
	board_map := make(map[int][][]int)
	var board [][]int
	board_counter := 0
	for index, line := range lines[1:] {
		// bypass first empty line
		if index == 0 {
			continue
		}
		// leave out every empty row (6th row when board is 5x5)
		if index % 6 == 0 {
			board_map[board_counter] = board
			board_counter++
			board = nil
			continue
		}
		ints, _ := ReadInts(strings.NewReader(line))
		board = append(board, ints)

	}
	// Part 1 https://adventofcode.com/2021/day/4
	// Part 2 https://adventofcode.com/2021/day/4
	for _, bingo_number := range bingo_numbers {
		for key, board := range board_map {
			if CheckForBingo(board) == true {
				continue
			}
			for _, row := range board {
				for row_index, number := range row {
					if number == bingo_number {
						// We can set matched number to -1.
						row[row_index] = -1
					}
				}
			}
			if CheckForBingo(board) {
				fmt.Println("Board", key, "has Bingo!")
				fmt.Println("Bingo number:", bingo_number, "* Sum of unmarked board:", SumUnmarkedBoard(board), "=", bingo_number * SumUnmarkedBoard(board))
			}
		}
	}
}
