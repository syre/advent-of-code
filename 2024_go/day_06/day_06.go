package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func find_guard_pos(arr [][]string) (x int, y int) {
	for row_index, row := range arr {
		for col_index, _ := range row {
			current_element := arr[row_index][col_index]
			if current_element == "^" || current_element == "<" || current_element == ">" || current_element == "v" {
				return row_index, col_index
			}
		}
	}
	return -1, -1
}

func rotate_guard(guard_symbol string) string {
	guard_symbol_map := map[string]string{">": "v", "v": "<", "^": ">", "<": "^"}
	return guard_symbol_map[guard_symbol]
}

func is_end_of_map(arr [][]string, guard_pos_x int, guard_pos_y int) bool {
	if guard_pos_y > len(arr)-1 {
		return true
	} else if guard_pos_y < 0 {
		return true
	} else if guard_pos_x < 0 {
		return true
	} else if guard_pos_x > len(arr[0])-1 {
		return true
	}
	return false
}

func main() {
	input, _ := os.ReadFile("test_input")
	guard_map_lists := strings.Split(string(input), "\n")
	guard_map := [][]string{}
	// populate guard_map
	for _, guard_map_list := range guard_map_lists {
		if len(guard_map_list) == 0 {
			continue
		}
		entry := strings.Split(guard_map_list, "")
		guard_map = append(guard_map, entry)
	}
	distinct_guard_positions := make(map[[2]int]bool)
	// this is ugly
	for {
		guard_pos_x, guard_pos_y := find_guard_pos(guard_map)
		distinct_guard_positions[[2]int{guard_pos_x, guard_pos_y}] = true
		guard_symbol := guard_map[guard_pos_x][guard_pos_y]
		if guard_pos_x == -1 && guard_pos_y == -1 {
			break
		}
		new_pos_x := -1
		new_pos_y := -1
		if guard_symbol == "^" {
			new_pos_x = guard_pos_x - 1
			new_pos_y = guard_pos_y
		} else if guard_symbol == ">" {
			new_pos_x = guard_pos_x
			new_pos_y = guard_pos_y + 1
		} else if guard_symbol == "v" {
			new_pos_x = guard_pos_x + 1
			new_pos_y = guard_pos_y
		} else if guard_symbol == "<" {
			new_pos_x = guard_pos_x
			new_pos_y = guard_pos_y - 1
		}
		if is_end_of_map(guard_map, new_pos_x, new_pos_y) {
			break
		}
		if guard_map[new_pos_x][new_pos_y] == "#" {
			guard_map[guard_pos_x][guard_pos_y] = rotate_guard(guard_map[guard_pos_x][guard_pos_y])
		} else {
			guard_map[new_pos_x][new_pos_y] = guard_symbol
			guard_map[guard_pos_x][guard_pos_y] = "."
		}
		//print_2d_array(guard_map)
		//fmt.Println("-------------------")
	}
	fmt.Println(len(slices.Collect(maps.Keys(distinct_guard_positions))))
}
