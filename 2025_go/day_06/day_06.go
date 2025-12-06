package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func main() {
	input, _ := os.ReadFile("test_input")
	problem_input := strings.Split(string(input), "\n")
	problem_matrix := [][]string{}
	// populate problem matrix
	for index, val := range problem_input {
		problem_matrix = append(problem_matrix, []string{})
		splitted := strings.Fields(val)
		for _, el := range splitted {
			problem_matrix[index] = append(problem_matrix[index], el)
		}
	}
	print_2d_array(problem_matrix)

	total_sum := 0
	for i := 0; i < len(problem_matrix)-1; i++ {
		column_sum := 0
		for j := 0; j < len(problem_matrix[i])-1; j++ {
			number, _ := strconv.Atoi(problem_matrix[j][i])

			operator := problem_matrix[len(problem_matrix[0])-1][i]
			if operator == "*" {
				column_sum *= number
			} else if operator == "+" {
				column_sum += number
			}
		}
		fmt.Println(column_sum)
		total_sum += column_sum
	}
	fmt.Println(total_sum)
}
