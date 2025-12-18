package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func int_abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

type coord struct {
	x int
	y int
}

func part_one() {
	input, _ := os.ReadFile("input")
	tiles_input := strings.Split(string(input), "\n")
	tiles_list := []coord{}
	// load in tiles
	for _, val := range tiles_input {
		x_y_val := strings.Split(val, ",")
		x, _ := strconv.Atoi(x_y_val[0])
		y, _ := strconv.Atoi(x_y_val[1])
		tiles_list = append(tiles_list, coord{x: x, y: y})
	}

	max_area := 0
	for _, coord1 := range tiles_list {
		for _, coord2 := range tiles_list {
			if coord1 == coord2 {
				continue
			}
			area := int_abs(coord1.x+1-coord2.x) * int_abs(coord1.y+1-coord2.y)
			if area > max_area {
				max_area = area
				fmt.Println(coord1, coord2, area)
			}
		}
	}
}

func check_

func part_two() {
	input, _ := os.ReadFile("test_input")
	tiles_input := strings.Split(string(input), "\n")
	tiles_list := []coord{}
	// load in tiles
	for _, val := range tiles_input {
		x_y_val := strings.Split(val, ",")
		x, _ := strconv.Atoi(x_y_val[0])
		y, _ := strconv.Atoi(x_y_val[1])
		tiles_list = append(tiles_list, coord{x: x, y: y})
	}

	max_area := 0
	for _, coord1 := range tiles_list {
		for _, coord2 := range tiles_list {
			if coord1 == coord2 {
				continue
			}
			area := int_abs(coord1.x+1-coord2.x) * int_abs(coord1.y+1-coord2.y)
			if area > max_area {
				max_area = area
				fmt.Println(coord1, coord2, area)
			}
		}
	}

}

func main() {
	//part_one()
	part_two()
}
