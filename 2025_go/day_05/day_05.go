package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ingredient_range struct {
	min int
	max int
}

func part_one() {
	input, _ := os.ReadFile("input")
	ingredients_input := strings.Split(string(input), "\n\n")

	fresh_ingredient_ranges_str := ingredients_input[0]
	avaiable_ingredient_ids_str := ingredients_input[1]

	fresh_ingredient_id_ranges := strings.Split(fresh_ingredient_ranges_str, "\n")
	available_ingredient_id_ranges := strings.Split(avaiable_ingredient_ids_str, "\n")

	ingredients_list := []ingredient_range{}

	// populate ingredients_list
	for _, val := range fresh_ingredient_id_ranges {
		val := strings.Split(val, "-")
		start_range, _ := strconv.Atoi(val[0])
		end_range, _ := strconv.Atoi(val[1])
		ingredients_list = append(ingredients_list, ingredient_range{start_range, end_range})
	}

	fresh_ingredients_counter := 0
	for _, available_ingredient := range available_ingredient_id_ranges {
		available_ingredient_int, _ := strconv.Atoi(available_ingredient)
		for _, ingredient_range := range ingredients_list {
			if ingredient_range.max >= available_ingredient_int && available_ingredient_int >= ingredient_range.min {
				fresh_ingredients_counter += 1
				break
			}
		}
	}
	fmt.Println(fresh_ingredients_counter)
}

func part_two() {
	input, _ := os.ReadFile("input")
	ingredients_input := strings.Split(string(input), "\n\n")

	fresh_ingredient_ranges_str := ingredients_input[0]

	fresh_ingredient_id_ranges := strings.Split(fresh_ingredient_ranges_str, "\n")

	ingredients_range_list := []ingredient_range{}

	// populate ingredients_range_list
	for _, val := range fresh_ingredient_id_ranges {
		val := strings.Split(val, "-")
		start_range, _ := strconv.Atoi(val[0])
		end_range, _ := strconv.Atoi(val[1])
		ingredients_range_list = append(ingredients_range_list, ingredient_range{start_range, end_range})
	}

	indices_to_be_removed := []int{}
	for _, x := range ingredients_range_list {
		for j, y := range ingredients_range_list {
			if x.min < y.min && x.max > y.max {
				indices_to_be_removed = append(indices_to_be_removed, j)
			}
		}
	}
	fmt.Println(len(ingredients_range_list))
	fmt.Println(indices_to_be_removed)
}

func main() {
	//part_one()
	part_two()
}
