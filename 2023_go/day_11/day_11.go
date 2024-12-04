package main

import (
	"os"
	"fmt"
	"strings"
	"math"
)

type Coords struct {
	x int
	y int
}

type GalaxyPair struct {
	a Coords
	b Coords
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(int_arr []int) int {
    min_found := math.MaxInt64
    for _, val := range int_arr {
        if val < min_found {
            min_found = val
        }
    }
    return min_found
}

func contains_no_galaxies(arr []string) bool {
	for _, value := range arr {
		if value != "." {
			return false
		}
	}
	return true
}

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func part_one() {
	input, _ := os.ReadFile("test_input")
	galaxy_lists := strings.Split(string(input), "\n")
	galaxy_map := [][]string {}
	for _, galaxy_list := range galaxy_lists {
		entry := strings.Split(galaxy_list, "")
		galaxy_map = append(galaxy_map, entry)
	}
	print_2d_array(galaxy_map)

	// expand at row-level
	skip_index := 0
	for row_index, _ := range galaxy_map {
		if contains_no_galaxies(galaxy_map[row_index+skip_index]) {
			galaxy_map = append(galaxy_map[:row_index+1+skip_index], galaxy_map[row_index+skip_index:]...)
			galaxy_map[row_index+skip_index] = galaxy_map[row_index+skip_index]
			skip_index += 1
		}
	}
	fmt.Println("-------------------")
	print_2d_array(galaxy_map)

	
	// fetch columns that need to be expanded to expand_column_indexes
	expand_column_indexes := []int {}
	for column_index := 0; column_index < len(galaxy_map[0]); column_index++ {
	    column := []string {}
	    for _, row := range galaxy_map {
	        column = append(column, row[column_index])
	    }
	    if contains_no_galaxies(column) {
	    	expand_column_indexes = append(expand_column_indexes, column_index)
	    	fmt.Println(column_index, column)
	    }
	}
	// expand at column-level
	col_skip_index := 0
	for _, column_index := range expand_column_indexes {
		for row_index, _ := range galaxy_map {
	    	galaxy_map[row_index] = append(galaxy_map[row_index][:column_index+1+col_skip_index], galaxy_map[row_index][column_index+col_skip_index:]...)
	    	galaxy_map[row_index][column_index+col_skip_index] = galaxy_map[row_index][column_index+col_skip_index]
	    }
	    col_skip_index += 1
	}
	fmt.Println("----------------------")
	print_2d_array(galaxy_map)

	// build a list of galaxy-coords
	galaxy_coords_list := []Coords {}
	// build a map of (galaxy-pair-coords -> distance
	coords_shortest_path_map := map[GalaxyPair]int {}
	for row_index, _ := range galaxy_map {
		for col_index, _ := range galaxy_map[0] {
			if galaxy_map[row_index][col_index] == "#" {
				fmt.Println(row_index, col_index)
				galaxy_coords_list = append(galaxy_coords_list, Coords{x:row_index, y:col_index})
			}
		}
	}
	for _, galaxy := range galaxy_coords_list {
		for _, other_galaxy := range galaxy_coords_list {
			if galaxy.x == other_galaxy.x && galaxy.y == other_galaxy.y {
				continue
			}
			_, a_b_pair_exists := coords_shortest_path_map[GalaxyPair{a:Coords{x:galaxy.x, y:galaxy.y}, b:Coords{x:other_galaxy.x, y:other_galaxy.y}}]
			_, b_a_pair_exists := coords_shortest_path_map[GalaxyPair{a:Coords{x:other_galaxy.x, y:other_galaxy.y}, b:Coords{x:galaxy.x, y:galaxy.y}}]
			if a_b_pair_exists || b_a_pair_exists {
				continue
			}
			distance := abs(galaxy.x - other_galaxy.x) + abs(galaxy.y - other_galaxy.y)
			coords_shortest_path_map[GalaxyPair{a:Coords{x:galaxy.x, y:galaxy.y}, b:Coords{x:other_galaxy.x, y:other_galaxy.y}}] = distance
		}
	}
	fmt.Println(galaxy_coords_list)
	fmt.Println(coords_shortest_path_map)
 	sum := 0
 	for _, value := range coords_shortest_path_map {
 		sum += value
 	}
 	fmt.Println(sum)
}

func part_two() {
	input, _ := os.ReadFile("input")
	galaxy_lists := strings.Split(string(input), "\n")
	galaxy_map := [][]string {}
	for _, galaxy_list := range galaxy_lists {
		entry := strings.Split(galaxy_list, "")
		galaxy_map = append(galaxy_map, entry)
	}

	// expand at row-level
	expand_row_indexes := []int {}
	for row_index, _ := range galaxy_map {
		if contains_no_galaxies(galaxy_map[row_index]) {
			expand_row_indexes = append(expand_row_indexes, row_index)
		}
	}
	
	// fetch columns that need to be expanded to expand_column_indexes
	expand_column_indexes := []int {}
	for column_index := 0; column_index < len(galaxy_map[0]); column_index++ {
	    column := []string {}
	    for _, row := range galaxy_map {
	        column = append(column, row[column_index])
	    }
	    if contains_no_galaxies(column) {
	    	expand_column_indexes = append(expand_column_indexes, column_index)
	    	//fmt.Println(column_index, column)
	    }
	}

	// build a list of galaxy-coords
	galaxy_coords_list := []Coords {}
	for row_index, _ := range galaxy_map {
		for col_index, _ := range galaxy_map[0] {
			if galaxy_map[row_index][col_index] == "#" {
				//fmt.Println(row_index, col_index)
				galaxy_coords_list = append(galaxy_coords_list, Coords{x:row_index, y:col_index})
			}
		}
	}

	// build a map of (galaxy-pair-coords -> distance
	coords_shortest_path_map := map[GalaxyPair]int {}
	for _, galaxy := range galaxy_coords_list {
		for _, other_galaxy := range galaxy_coords_list {
			if galaxy.x == other_galaxy.x && galaxy.y == other_galaxy.y {
				continue
			}
			_, a_b_pair_exists := coords_shortest_path_map[GalaxyPair{a:Coords{x:galaxy.x, y:galaxy.y}, b:Coords{x:other_galaxy.x, y:other_galaxy.y}}]
			_, b_a_pair_exists := coords_shortest_path_map[GalaxyPair{a:Coords{x:other_galaxy.x, y:other_galaxy.y}, b:Coords{x:galaxy.x, y:galaxy.y}}]
			if a_b_pair_exists || b_a_pair_exists {
				continue
			}
			add_to_x := 0
			add_to_y := 0
			for _, index := range expand_row_indexes {
				if galaxy.x > other_galaxy.x {
					if other_galaxy.x < index && index < galaxy.x {
						add_to_x += 999999
					}
				} else {
					if galaxy.x < index && index < other_galaxy.x {
						add_to_x += 999999
					}

				}
			}
			for _, index := range expand_column_indexes {
				if galaxy.y > other_galaxy.y {
					if other_galaxy.y < index && index < galaxy.y {
						add_to_y += 999999
					}
				} else {
					if galaxy.y < index && index < other_galaxy.y {
						add_to_y += 999999
					}

				}
			}
			distance := abs(galaxy.x - other_galaxy.x) + add_to_x + abs(galaxy.y - other_galaxy.y) + add_to_y
			coords_shortest_path_map[GalaxyPair{a:Coords{x:galaxy.x, y:galaxy.y}, b:Coords{x:other_galaxy.x, y:other_galaxy.y}}] = distance
		}
	}
	//fmt.Println(galaxy_coords_list)
	//fmt.Println(coords_shortest_path_map)
 	sum := 0
 	for _, value := range coords_shortest_path_map {
 		sum += value
 	}
 	fmt.Println(sum)
}

func main() {
	//part_one()
	part_two()
}