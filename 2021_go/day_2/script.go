package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input),"\n")
	direction_map := map[string]int{
		"forward": 0,
		"down": 0,
		"up": 0,
	}
	// Part 1 https://adventofcode.com/2021/day/1
	for _, s := range lines {
		line := strings.Split(s, " ")
		direction := line[0]
		step, _ := strconv.Atoi(line[1])
		direction_map[direction] += step
	}
	fmt.Println(direction_map["forward"]*(direction_map["down"]-direction_map["up"]))
	// Part 2 https://adventofcode.com/2021/day/1
	direction_map = map[string]int{
		"forward": 0,
		"down": 0,
		"up": 0,
		"aim": 0,
	}
	for _, s := range lines {
		line := strings.Split(s, " ")
		direction := line[0]
		step, _ := strconv.Atoi(line[1])
		if direction == "down" {
			direction_map["aim"] += step
		} else if direction == "up" {
			direction_map["aim"] -= step
		} else if direction == "forward" {
			direction_map["forward"] += step
			direction_map["down"] += direction_map["aim"] * step
		}
	}
	fmt.Println(direction_map["forward"]*(direction_map["down"]-direction_map["up"]))
}