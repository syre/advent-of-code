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

func find_root_pos(arr [][]string) coord {
	for i, _ := range arr {
		for j, _ := range arr[i] {
			if arr[i][j] == "S" {
				return coord{i, j}
			}
		}
	}
	return coord{}
}
func advance_beams(beams_arr []coord, tachyon_diagram [][]string) (bool, int, []coord, [][]string) {
	new_beams_arr := []coord{}
	finished := false
	beams_split := 0
	for _, beam := range beams_arr {
		el := tachyon_diagram[beam.x+1][beam.y]
		if el == "." {
			new_beams_arr = append(new_beams_arr, coord{beam.x + 1, beam.y})
			tachyon_diagram[beam.x+1][beam.y] = "|"
		} else if el == "^" {
			beams_split += 1
			new_beams_arr = append(new_beams_arr, coord{beam.x + 1, beam.y + 1})
			tachyon_diagram[beam.x+1][beam.y+1] = "|"
			new_beams_arr = append(new_beams_arr, coord{beam.x + 1, beam.y - 1})
			tachyon_diagram[beam.x+1][beam.y-1] = "|"
		}
	}
	// check if end is reached:
	for _, beam := range new_beams_arr {
		if beam.x == len(tachyon_diagram)-1 {
			finished = true
		}
	}
	return finished, beams_split, new_beams_arr, tachyon_diagram
}

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
	fmt.Println(strings.Repeat("-", len(arr)*2))
}

func part_one() {
	input, _ := os.ReadFile("input")
	tachyon_input := strings.Split(string(input), "\n")
	tachyon_diagram := [][]string{}
	// populate tachyon diagram matrix
	for index, val := range tachyon_input {
		tachyon_diagram = append(tachyon_diagram, []string{})
		splitted := strings.Split(val, "")
		for _, el := range splitted {
			tachyon_diagram[index] = append(tachyon_diagram[index], el)
		}
	}
	print_2d_array(tachyon_diagram)

	// find root node pos
	root_pos := find_root_pos(tachyon_diagram)
	beams_arr := []coord{root_pos}
	// advance beams
	finished := false
	beams_split_sum := 0
	beams_split := 0
	for finished != true {
		finished, beams_split, beams_arr, tachyon_diagram = advance_beams(beams_arr, tachyon_diagram)
		beams_split_sum += beams_split
		print_2d_array(tachyon_diagram)
	}
	fmt.Println(beams_split_sum)
}

func main() {
	part_one()
}
