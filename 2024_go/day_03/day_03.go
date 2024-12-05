package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part_one() {
	input, _ := os.ReadFile("test_input")
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	instruction_list := re.FindAll(input, -1)

	instruction_sum := 0
	fmt.Println(len(instruction_list))
	for _, instruction := range instruction_list {
		string_instruction := string(instruction)
		first_number_start_index := 4
		last_number_end_index := len(instruction) - 1
		numbers_list := strings.Split(string_instruction[first_number_start_index:last_number_end_index], ",")
		first_number, _ := strconv.Atoi(numbers_list[0])
		second_number, _ := strconv.Atoi(numbers_list[1])
		instruction_sum += first_number * second_number
	}
	fmt.Println(instruction_sum)
}

func part_two() {
	input, _ := os.ReadFile("test_input")
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	instruction_list := re.FindAll(input, -1)

	run_instruction := true
	instruction_sum := 0
	for _, instruction := range instruction_list {
		string_instruction := string(instruction)
		fmt.Println(string_instruction)
		if string(string_instruction) == "don't()" {
			run_instruction = false
			fmt.Println("run instruction false")
			continue
		}
		if string(string_instruction) == "do()" {
			run_instruction = true
			fmt.Println("run instruction true")
			continue
		}
		if !run_instruction {
			continue
		}
		fmt.Println(string_instruction)
		first_number_start_index := 4
		last_number_end_index := len(string_instruction) - 1
		numbers_list := strings.Split(string_instruction[first_number_start_index:last_number_end_index], ",")
		first_number, _ := strconv.Atoi(numbers_list[0])
		second_number, _ := strconv.Atoi(numbers_list[1])
		instruction_sum += first_number * second_number
	}
	fmt.Println(instruction_sum)

}

func main() {
	//part_one()
	part_two()
}
