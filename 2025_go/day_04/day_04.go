package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x int
	y int
}

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func out_of_bounds(arr [][]string, i, j int) bool {
	if i > len(arr)-1 || i < 0 {
		return true
	}
	if j > len(arr[0])-1 || j < 0 {
		return true
	}
	return false
}

func part_one() {
	input, _ := os.ReadFile("input")
	toilet_rolls_lists := strings.Split(string(input), "\n")
	toilet_rolls_map := [][]string{}
	// populate toilet roll map
	for _, toilet_roll_map_list := range toilet_rolls_lists {
		if len(toilet_roll_map_list) == 0 {
			continue
		}
		entry := strings.Split(toilet_roll_map_list, "")
		toilet_rolls_map = append(toilet_rolls_map, entry)
	}
	print_2d_array(toilet_rolls_map)

	toilet_rolls_accessable_by_fork_lift := 0
	for i := range toilet_rolls_map {
		for j := range toilet_rolls_map[0] {
			if toilet_rolls_map[i][j] != "@" {
				continue
			}
			toilet_roll_counter := 0

			if !out_of_bounds(toilet_rolls_map, i-1, j-1) && toilet_rolls_map[i-1][j-1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i-1, j) && toilet_rolls_map[i-1][j] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i-1, j+1) && toilet_rolls_map[i-1][j+1] == "@" {
				toilet_roll_counter++
			}

			if !out_of_bounds(toilet_rolls_map, i, j+1) && toilet_rolls_map[i][j+1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i, j-1) && toilet_rolls_map[i][j-1] == "@" {
				toilet_roll_counter++
			}

			if !out_of_bounds(toilet_rolls_map, i+1, j-1) && toilet_rolls_map[i+1][j-1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i+1, j) && toilet_rolls_map[i+1][j] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i+1, j+1) && toilet_rolls_map[i+1][j+1] == "@" {
				toilet_roll_counter++
			}
			if toilet_roll_counter < 4 {
				toilet_rolls_accessable_by_fork_lift += 1
			}
		}
	}
	fmt.Println(toilet_rolls_accessable_by_fork_lift)
}

func find_removed_toilet_roll_indices(toilet_rolls_map [][]string) []coord {
	coords := []coord{}
	for i := range len(toilet_rolls_map) {
		for j := range len(toilet_rolls_map[0]) {
			if toilet_rolls_map[i][j] != "@" {
				continue
			}
			toilet_roll_counter := 0

			if !out_of_bounds(toilet_rolls_map, i-1, j-1) && toilet_rolls_map[i-1][j-1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i-1, j) && toilet_rolls_map[i-1][j] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i-1, j+1) && toilet_rolls_map[i-1][j+1] == "@" {
				toilet_roll_counter++
			}

			if !out_of_bounds(toilet_rolls_map, i, j+1) && toilet_rolls_map[i][j+1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i, j-1) && toilet_rolls_map[i][j-1] == "@" {
				toilet_roll_counter++
			}

			if !out_of_bounds(toilet_rolls_map, i+1, j-1) && toilet_rolls_map[i+1][j-1] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i+1, j) && toilet_rolls_map[i+1][j] == "@" {
				toilet_roll_counter++
			}
			if !out_of_bounds(toilet_rolls_map, i+1, j+1) && toilet_rolls_map[i+1][j+1] == "@" {
				toilet_roll_counter++
			}
			if toilet_roll_counter < 4 {
				coords = append(coords, coord{i, j})
			}
		}
	}
	return coords
}

func part_two() {
	input, _ := os.ReadFile("input")
	toilet_rolls_lists := strings.Split(string(input), "\n")
	toilet_rolls_map := [][]string{}
	// populate toilet roll map
	for _, toilet_roll_map_list := range toilet_rolls_lists {
		if len(toilet_roll_map_list) == 0 {
			continue
		}
		entry := strings.Split(toilet_roll_map_list, "")
		toilet_rolls_map = append(toilet_rolls_map, entry)
	}
	print_2d_array(toilet_rolls_map)

	toilet_rolls_accessable_by_fork_lift := 0
	num_indices := -1
	for num_indices != 0 {
		indices := find_removed_toilet_roll_indices(toilet_rolls_map)
		for _, index := range indices {
			toilet_rolls_map[index.x][index.y] = "."
		}
		toilet_rolls_accessable_by_fork_lift += len(indices)
		num_indices = len(indices)
	}
	fmt.Println(toilet_rolls_accessable_by_fork_lift)
}

func main() {
	part_one()
	part_two()
}
