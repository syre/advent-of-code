package main

import (
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
	part_two_zero_counter := 0

	println("Dial starts at", dial_val)

	for _, line := range string_lists {
		if line == "" {
			continue
		}

		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		if direction == "L" {
			if dial_val-distance < 0 {
				part_two_zero_counter += int_abs((dial_val - distance) / 100)

				dial_val = (100 - int_abs(dial_val-distance)) % 100
			} else {
				dial_val -= distance
			}
		} else if direction == "R" {
			if dial_val+distance > 99 {
				part_two_zero_counter += (dial_val + distance) / 100

				dial_val = (dial_val + distance) % 100
			} else {
				dial_val += distance
			}
		}

		if dial_val == 0 {
			zero_counter += 1
		}

		println("dial is rotated", direction, distance, "to point at", dial_val)
		println("part two zero counter at", part_two_zero_counter)
	}
	println("zero counter at", zero_counter)
}

func main() {
	part_one()
}
