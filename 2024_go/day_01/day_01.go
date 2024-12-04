package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs_diff_int(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func part_one() {
	input, _ := os.ReadFile("test_input")
	string_lists := strings.Split(string(input), "\n")

	left_numbers_list := []int{}
	right_numbers_list := []int{}

	for _, string_list := range string_lists {
		if string_list == "" {
			continue
		}
		splitted := strings.Split(string(string_list), "   ")
		left_number, _ := strconv.Atoi(splitted[0])
		right_number, _ := strconv.Atoi(splitted[1])
		left_numbers_list = append(left_numbers_list, left_number)
		right_numbers_list = append(right_numbers_list, right_number)
	}
	slices.Sort(left_numbers_list)
	slices.Sort(right_numbers_list)
	if len(left_numbers_list) != len(right_numbers_list) {
		log.Fatal("length of numbers lists differ")
	}

	difference_sum := 0
	for i := 0; i < len(left_numbers_list); i++ {
		difference_sum += abs_diff_int(left_numbers_list[i], right_numbers_list[i])

	}
	fmt.Println(difference_sum)
}

func part_two() {
	input, _ := os.ReadFile("test_input")
	string_lists := strings.Split(string(input), "\n")

	left_numbers_list := []int{}
	right_numbers_list := []int{}
	right_numbers_map := make(map[int]int)

	for _, string_list := range string_lists {
		if string_list == "" {
			continue
		}
		splitted := strings.Split(string(string_list), "   ")
		left_number, _ := strconv.Atoi(splitted[0])
		right_number, _ := strconv.Atoi(splitted[1])
		left_numbers_list = append(left_numbers_list, left_number)
		right_numbers_list = append(right_numbers_list, right_number)
		right_numbers_map[right_number] += 1
	}
	if len(left_numbers_list) != len(right_numbers_list) {
		log.Fatal("length of numbers lists differ")
	}

	similarity_score_sum := 0
	for i := 0; i < len(left_numbers_list); i++ {
		similarity_score_sum += left_numbers_list[i] * right_numbers_map[left_numbers_list[i]]

	}
	fmt.Println(similarity_score_sum)

}

func main() {
	part_one()
	part_two()
}
