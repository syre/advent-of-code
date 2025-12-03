package main

import (
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
	input, _ := os.ReadFile("test_input")
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

func main() {
	//part_one()
	part_two()
}
