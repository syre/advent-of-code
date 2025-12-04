package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func int_abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part_one() {
	input, _ := os.ReadFile("test_input")
	string_lists := strings.Split(string(input), "\n")
	dial_val := 50
	zero_counter := 0

	println("dial starts at", dial_val)
	for _, line := range string_lists {
		if line == "" {
			continue
		}

		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		if direction == "L" {
			if dial_val-distance < 0 {
				dial_val = (100 - int_abs(dial_val-distance)) % 100
			} else {
				dial_val -= distance
			}
		} else if direction == "R" {
			if dial_val+distance > 99 {
				dial_val = (dial_val + distance) % 100
			} else {
				dial_val += distance
			}
		}

		if dial_val == 0 {
			zero_counter += 1
		}

		println("dial is rotated", direction, distance, "to point at", dial_val)
	}
	println("zero counter at", zero_counter)
}

func part_two() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), "\n")
	dial_val := 50
	zero_counter := 0

	println("dial starts at", dial_val)
	for _, line := range string_lists {
		if line == "" {
			continue
		}

		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		if direction == "L" {
			for i := distance; i >= 0; i-- {
				if dial_val == -1 {
					fmt.Println("incrementing zero counter")
					zero_counter += 1
					dial_val = 99
					continue
				}
				dial_val -= 1
			}
		} else if direction == "R" {
			for i := 0; i < distance; i++ {
				if dial_val == 99 {
					fmt.Println("incrementing zero counter")
					zero_counter += 1
					dial_val = 0
					continue
				}
				dial_val += 1
			}
		}
		println("dial is rotated", direction, distance, "to point at", dial_val)
	}
	println("zero counter at", zero_counter)
}

func main() {
	//part_one()
	part_two()
}
