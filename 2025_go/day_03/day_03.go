package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part_one() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), "\n")
	joltage_sum := 0
	for _, line := range string_lists {
		largest_joltage := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				joltage_string := string([]byte{line[i], line[j]})
				joltage, _ := strconv.Atoi(joltage_string)
				if joltage > largest_joltage {
					largest_joltage = joltage
				}
			}
		}
		joltage_sum += largest_joltage
	}
	println(joltage_sum)
}

func part_two() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), "\n")
	joltage_sum := 0
	for _, line := range string_lists {
		var joltage_str strings.Builder
		chars_left := 12
		// we know first number must be largest in the span of 0 -> len(line)-12
		// next number largest in the span of previous number index -> len(line)-11
		// and so on
		iterator_start := 0
		iterator_end := len(line) + 1 - chars_left

		for chars_left > 0 {
			largest_num := -1
			largest_num_index := -1
			for i := iterator_start; i < iterator_end; i++ {
				num, _ := strconv.Atoi(string(line[i]))
				if num > largest_num {
					largest_num = num
					largest_num_index = i
				}
			}
			joltage_str.WriteString(string(line[largest_num_index]))
			iterator_start = largest_num_index + 1
			iterator_end = len(line) + 2 - chars_left
			chars_left -= 1
		}
		joltage_int, _ := strconv.Atoi(joltage_str.String())
		fmt.Println("largest joltage:", joltage_str.String())
		joltage_sum += joltage_int
	}
	println("joltage sum:", joltage_sum)
}

func main() {
	//part_one()
	part_two()
}
