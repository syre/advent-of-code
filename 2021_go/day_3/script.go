package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"io"
	"bufio"
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

func Transpose(x [][]int) [][]int {
	out := make([][]int, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func Sum(input []int) int {
    sum := 0
    for _, i := range input {
        sum += i
    }
    return sum
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
	// Part 1 https://adventofcode.com/2021/day/3
	transposed_matrix := Transpose(int_matrix)

	gamma_rate := ""
	for _, arr := range transposed_matrix {
		// 1 is dominant
		if sum(arr) > len(arr)/2 {
			gamma_rate += "1"
		} else {
			gamma_rate += "0"
		}
	}

	gamma_rate_dec, _ := strconv.ParseInt(gamma_rate, 2, 64) 
	fmt.Println("gamma rate binary:", gamma_rate)
	fmt.Println("gamma rate decimal:", gamma_rate_dec)

	epsilon_rate := ""
	for _, arr := range transposed_matrix {
		// 1 is dominant
		if Sum(arr) > len(arr)/2 {
			epsilon_rate += "0"
		} else {
			epsilon_rate += "1"
		}
	}
	epsilon_rate_dec, _ := strconv.ParseInt(epsilon_rate, 2, 64)
	fmt.Println("epsilon rate binary:", epsilon_rate)
	fmt.Println("epsilon rate decimal:", epsilon_rate_dec)

	fmt.Println("power consumption:", gamma_rate_dec * epsilon_rate_dec)

	// Part 2 https://adventofcode.com/2021/day/3
	oxygen_generator_slice := int_matrix[:][:]
	for gamma_index, value := range gamma_rate {
		value_int, _ := strconv.Atoi(string(value))

		for matrix_index, arr := range int_matrix {
			if arr[gamma_index] != value_int {
				fmt.Println("removing matrix index: ", matrix_index)
				oxygen_generator_slice = append(oxygen_generator_slice[:matrix_index], oxygen_generator_slice[matrix_index+1:]...)
			}
			if len(oxygen_generator_slice) == 1 {
				fmt.Println(oxygen_generator_slice)
				break
			}

		}
	} 
}
