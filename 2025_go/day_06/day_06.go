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

func part_one() {
	input, _ := os.ReadFile("test_input")
	problem_input := strings.Split(string(input), "\n")
	problem_matrix := [][]string{}
	// populate problem matrix
	for index, val := range problem_input {
		problem_matrix = append(problem_matrix, []string{})
		// didn't know about strings.Fields!
		splitted := strings.Fields(val)
		for _, el := range splitted {
			problem_matrix[index] = append(problem_matrix[index], el)
		}
	}
	print_2d_array(problem_matrix)

	total_sum := 0
	for i := 0; i < len(problem_matrix[0]); i++ {
		operator := problem_matrix[len(problem_matrix)-1][i]
		var column_sum int
		if operator == "*" {
			column_sum = 1
		} else if operator == "+" {
			column_sum = 0
		}
		for j := 0; j < len(problem_matrix)-1; j++ {
			number, _ := strconv.Atoi(problem_matrix[j][i])
			if operator == "*" {
				column_sum *= number
			} else if operator == "+" {
				column_sum += number
			}
		}
		total_sum += column_sum
	}
	fmt.Println(total_sum)
}

func part_two() {
	input, _ := os.ReadFile("test_input")
	problem_input := strings.Split(string(input), "\n")
	problem_matrix := [][]string{}
	// populate problem matrix - this time just split each character
	for index, val := range problem_input {
		problem_matrix = append(problem_matrix, []string{})
		splitted := strings.Split(val, "")
		for _, el := range splitted {
			problem_matrix[index] = append(problem_matrix[index], el)
		}
	}
	print_2d_array(problem_matrix)

	number_sum := 0
	var number_str strings.Builder
	numbers_list := []int{}
	for j := len(problem_matrix[0]) - 1; j > 0; j-- {
		for i := 0; i < len(problem_matrix); i++ {
			el := problem_matrix[i][j]
			if i == len(problem_matrix)-1 {
				number, err := strconv.Atoi(strings.TrimSpace(number_str.String()))
				number_str.Reset()
				if err == nil {
					numbers_list = append(numbers_list, number)
				}
			}

			if el == "" {
				continue
			} else if el == "*" {
				fmt.Println(numbers_list)
				mult_sum := 1
				for _, num := range numbers_list {
					mult_sum *= num
				}
				number_sum += mult_sum
				numbers_list = []int{}
			} else if el == "+" {
				fmt.Println(numbers_list)
				for _, num := range numbers_list {
					number_sum += num
				}
				numbers_list = []int{}
			}
			number_str.WriteString(el)
		}
	}
	fmt.Println(number_sum)
}

func main() {
	//part_one()
	part_two()
}
