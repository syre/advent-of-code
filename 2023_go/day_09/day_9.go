package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
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

func is_all_zeroes(int_arr []int) bool {
	for _, value := range int_arr {
		if value != 0 {
			return false
		}
	}
	return true
}

func sum(arr []int) int {
	n := 0
	for _, v := range arr {
		n += v
	}
	return n
}

func part_one() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")
	input_values := [][]int {}
	for _, input_string := range input_splitted {
		ints, _ := ReadIntsFromStringArray(strings.Fields(input_string))
		input_values = append(input_values, ints)
	}

	next_value_list := []int {}
	for _, value_list := range input_values {
		history_list := [][]int {}
		history_list = append(history_list, value_list)
		skip_index := 0
		for !is_all_zeroes(history_list[len(history_list)-1]) {
			for _, history := range history_list[skip_index:] {
				append_list := []int {}
				for index, _ := range history {				
					if index == 0 {
						continue
					}
					append_list = append(append_list, history[index]-history[index-1])
				}
				history_list = append(history_list, append_list)
				skip_index += 1
			}
		}
		// add zero to zero-element list to the end and work our way up on the right side
		history_list[len(history_list)-1] = append(history_list[len(history_list)-1], 0)
		fmt.Println("history list:", history_list)
		for j := len(history_list)-1; j > 0; j-- {
			value_to_the_left := history_list[j-1][len(history_list[j-1])-1]
			value_below := history_list[j][len(history_list[j])-1]
			fmt.Println("value to the left:", value_to_the_left)
			fmt.Println("value below", value_below)
			history_list[j-1] = append(history_list[j-1], value_to_the_left + value_below)
		}
		next_value_list = append(next_value_list, history_list[0][len(history_list[0])-1])
		fmt.Println("new history_list", history_list)
	}
	fmt.Println(sum(next_value_list))
}

func part_two() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")
	input_values := [][]int {}
	for _, input_string := range input_splitted {
		ints, _ := ReadIntsFromStringArray(strings.Fields(input_string))
		input_values = append(input_values, ints)
	}

	next_value_list := []int {}
	for _, value_list := range input_values {
		history_list := [][]int {}
		history_list = append(history_list, value_list)
		skip_index := 0
		for !is_all_zeroes(history_list[len(history_list)-1]) {
			for _, history := range history_list[skip_index:] {
				append_list := []int {}
				for index, _ := range history {				
					if index == 0 {
						continue
					}
					append_list = append(append_list, history[index]-history[index-1])
				}
				history_list = append(history_list, append_list)
				skip_index += 1
			}
		}
		// add zero to zero-element list first and work our way up on the left side
		history_list[len(history_list)-1] = append([]int{0}, history_list[len(history_list)-1]...)
		fmt.Println("history list:", history_list)
		for j := len(history_list)-1; j > 0; j-- {
			value_to_the_left := history_list[j-1][0]
			value_below := history_list[j][0]
			fmt.Println("value to the left:", value_to_the_left)
			fmt.Println("value below", value_below)
			history_list[j-1] = append([]int{value_to_the_left - value_below}, history_list[j-1]...)
		}
		next_value_list = append(next_value_list, history_list[0][0])
		fmt.Println("new history_list", history_list)
	}
	fmt.Println(sum(next_value_list))
}


func main() {
	part_one()
	part_two()
}