package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("test_input")
	problem_input := strings.Split(string(input), "\n")
	problem_matrix := [][]string{}
	// populate problem matrix
	for index, val := range problem_input {
		problem_matrix = append(problem_matrix, []string{})
		splitted := strings.Split(val, " ")
		for _, el := range splitted {
			trimmed_el := strings.TrimSpace(el)
			problem_matrix[index] = append(problem_matrix[index], trimmed_el)
		}
	}
	fmt.Println(problem_matrix)
	for i := 0; i < len(problem_matrix)-1; i++ {
		for j := 0; j < len(problem_matrix[0])-1; j++ {
			number, _ := strconv.Atoi(strings.TrimSpace(problem_matrix[j][i]))
			fmt.Println(number)
		}
	}
}
