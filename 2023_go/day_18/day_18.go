package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

type Coords struct {
	x int
	y int
}

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func calculate_cubic_meters(arr [][]string) int {
	sum := 0
	for _, row := range arr {
		for _, col := range row {
			if col == "#" {
				sum += 1
			}
		}
	}
	return sum
}

func map_column(map_obj [][]string, column_index int) (column []string) {
    column = make([]string, 0)
    for _, row := range map_obj {
        column = append(column, row[column_index])
    }
    return column
}

func remove_ground_level_terrain_axes(dig_map [][]string) [][]string {
	// filter empty rows
	row_skip_index := 0
	for row_index, _ := range dig_map {
		if is_all_ground_level_terrain(dig_map[row_index+row_skip_index]) {
			dig_map = append(dig_map[:row_index+row_skip_index], dig_map[row_index+1+row_skip_index:]...)
			row_skip_index -= 1
		}
	}
	// filter empty columns
	col_skip_index := 0
	for column_index, _ := range dig_map[0] {
		if is_all_ground_level_terrain(map_column(dig_map, column_index+col_skip_index)) {
			for row_index, _ := range dig_map {
		    	dig_map[row_index] = append(dig_map[row_index][:column_index+col_skip_index], dig_map[row_index][column_index+1+col_skip_index:]...)
		    }
		    col_skip_index -= 1
		}
	}
	return dig_map
}

func is_all_ground_level_terrain(arr []string) bool {
	for _,el := range arr {
		if el != "." {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("test_input")
	dig_plan_lists := strings.Split(string(input), "\n")

	rows := 550
	cols := 550
	dig_map := make([][]string, rows)
	for row_index, _ := range dig_map {
		dig_map[row_index] = make([]string, cols)
		for col_index, _ := range dig_map[row_index] {
			dig_map[row_index][col_index] = "."
		}
	}
	current_point := Coords{x:rows/2, y:cols/2}
	coords_to_hex_color_map := map[Coords]string {}
	// draw edge
	for _, dig_plan := range dig_plan_lists {
		direction := string(dig_plan[0])
		meters, _ := strconv.Atoi(dig_plan[2:strings.IndexRune(dig_plan, '(')-1])
		hex_color_code := dig_plan[strings.IndexRune(dig_plan, '(')+1:strings.IndexRune(dig_plan, ')')]

		if direction == "U" {
			for i := 0; i < meters; i++ {
				dig_map[current_point.x][current_point.y] = "#"
				coords_to_hex_color_map[current_point] = hex_color_code
				current_point.x -= 1
			}

		} else if direction == "D" {
			for i := 0; i < meters; i++ {
				dig_map[current_point.x][current_point.y] = "#"
				coords_to_hex_color_map[current_point] = hex_color_code
				current_point.x += 1
			}

		} else if direction == "L" {
			for i := 0; i < meters; i++ {
				dig_map[current_point.x][current_point.y] = "#"
				coords_to_hex_color_map[current_point] = hex_color_code
				current_point.y -= 1
			}

		} else if direction == "R" {
			for i := 0; i < meters; i++ {
				dig_map[current_point.x][current_point.y] = "#"
				coords_to_hex_color_map[current_point] = hex_color_code
				current_point.y += 1
			}
		}
	}
	dig_map = remove_ground_level_terrain_axes(dig_map)
	print_2d_array(dig_map)
	// fill out
	for row_index, _ := range dig_map {
		start_coords := Coords{x:row_index, y:0}
		start_coords_set := false
		for col_index, _ := range dig_map[row_index] {
			if dig_map[row_index][col_index] == "#" {
				if !start_coords_set {
					start_coords.y = col_index
					start_coords_set = true
				} else {
					for i := col_index; i > start_coords.y; i-- {
						dig_map[row_index][i] = "#"
					}
					start_coords_set = false
				}
			}
		}

	}
	//print_2d_array(dig_map)
	fmt.Println(calculate_cubic_meters(dig_map))
}