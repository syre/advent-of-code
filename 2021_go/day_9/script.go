package main

import (
 	"io"
 	"strconv"
 	"strings"
 	"bufio"
 	"os"
 	"fmt"
)

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    // ScanRunes to split on each utf-8 rune.
    scanner.Split(bufio.ScanRunes)
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

func main() {
	// Initialization.
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input), "\n")

	int_matrix := make([][]int, 0, len(lines))

	for _, line := range lines {
		ints, _ := ReadInts(strings.NewReader(line))
		int_matrix = append(int_matrix, ints)
	}

	fmt.Println(int_matrix)
	// Part 1 https://adventofcode.com/2021/day/9
	risk_counter := 0
	for i := 0; i < len(int_matrix); i++ {
		for j := 0; j < len(int_matrix[0]); j++ {
			value := int_matrix[i][j]

			if i+1 < len(int_matrix) {
				top_neighbor := int_matrix[i+1][j]
				if top_neighbor < value {
					continue
				}
			}
			if i-1 >= 0 {
				bottom_neighbor := int_matrix[i-1][j]
				if bottom_neighbor < value {
					continue
				}
			}
			if j+1 < len(int_matrix[0]) {
				right_neighbor := int_matrix[i][j+1]
				if right_neighbor < value {
					continue
				}
			}
			if j-1 >= 0 {
				left_neighbor := int_matrix[i][j-1]
				if left_neighbor < value {
					continue
				}
			}
			// 9 is a special case for this way of doing it (highest value so cannot be lower than neighbours).
			if value == 9 {
				continue
			}
			fmt.Println(value,"is lowest point at",i, j,"with risk", 1+value)
			risk_counter += 1+value
		}
	}
	fmt.Println(risk_counter)
	// Part 2 https://adventofcode.com/2021/day/9
}