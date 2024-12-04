package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func multiply(arr []int) int {
	multiplied := 1
	for _, value := range arr {
		multiplied *= value
	}
	return multiplied
}

func part_one() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")
	time_string := input_splitted[0]
	distance_string := input_splitted[1]

	time_list := []int {}
	distance_list := []int {}
	// populate time list and distance list
	for _, val := range strings.Fields(time_string[6:]) {
			int_val, err := strconv.Atoi(string(val))
			if err == nil {
				time_list = append(time_list, int_val)
			}
	}
	for _, val := range strings.Fields(distance_string[9:]) {
			int_val, err := strconv.Atoi(string(val))
			if err == nil {
				distance_list = append(distance_list, int_val)
			}
	}

	number_of_ways_list := []int {}

	for pos, time := range time_list {
		number_of_ways := 0
		hold_milliseconds := 0
		travel_milliseconds := 0

		for i:=0; i < time; i++ {
			hold_milliseconds = time - i
			travel_milliseconds = time - hold_milliseconds
			travelled := hold_milliseconds * travel_milliseconds
			if travelled > distance_list[pos] {
				number_of_ways += 1
			}	
		}
		number_of_ways_list = append(number_of_ways_list, number_of_ways)
	}
	fmt.Println(multiply(number_of_ways_list))
}

func part_two() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")
	time_string := input_splitted[0]
	distance_string := input_splitted[1]

	time := 0
	distance := 0
	// populate time list and distance list
	int_time, err := strconv.Atoi(strings.ReplaceAll(time_string[6:], " ", ""))
	if err == nil {
		time = int_time
	}

	int_distance, err := strconv.Atoi(strings.ReplaceAll(distance_string[9:], " ", ""))
	if err == nil {
		distance = int_distance
	}

	number_of_ways_list := []int {}

	number_of_ways := 0
	hold_milliseconds := 0
	travel_milliseconds := 0

	for i:=0; i < time; i++ {
		hold_milliseconds = time - i
		travel_milliseconds = time - hold_milliseconds
		travelled := hold_milliseconds * travel_milliseconds
		if travelled > distance {
			number_of_ways += 1
		}	
	}
	number_of_ways_list = append(number_of_ways_list, number_of_ways)
	fmt.Println(multiply(number_of_ways_list))
}

func main() {
	part_one()
	part_two()
}