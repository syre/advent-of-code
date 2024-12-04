package main

import (
	"os"
	"fmt"
	"strings"
	"reflect"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func map_column(map_obj [][]string, column_index int) (column []string) {
    column = make([]string, 0)
    for _, row := range map_obj {
        column = append(column, row[column_index])
    }
    return column
}

func main() {
	input, _ := os.ReadFile("test_input_2")
	terrain_lists := strings.Split(string(input), "\n\n")

	summarizing_sum := 0
	for _, terrain_list := range terrain_lists {
		terrain_map := [][]string {}
		for _, terrain_entry := range strings.Split(terrain_list, "\n") {
			entry := strings.Split(string(terrain_entry), "")
			terrain_map = append(terrain_map, entry)
		}

		print_2d_array(terrain_map)

		// check for row equality
		max_num_rows_equal := 0
		rows_above := 0
		for row_index := 0; row_index < len(terrain_map)-1; row_index++ {
			if reflect.DeepEqual(terrain_map[row_index], terrain_map[row_index+1]) {
				fmt.Println("row indexes", row_index, row_index+1, "has line of reflection")
				row_upper := row_index-1
				row_lower := row_index+2
				num_rows_equal := 1
				fmt.Println("row_upper", row_upper, "row_lower", row_lower)
				for row_upper >= 0 && row_lower < len(terrain_map) && reflect.DeepEqual(terrain_map[row_upper], terrain_map[row_lower]) {
					fmt.Println("comparing rows", terrain_map[row_upper], terrain_map[row_lower])
					num_rows_equal += 1
					row_upper -= 1
					row_lower += 1
				}
				if num_rows_equal > max_num_rows_equal && ((row_upper == 0 && row_lower == num_rows_equal+num_rows_equal+1) || (row_upper == len(terrain_map)-num_rows_equal-num_rows_equal-1 && row_lower == len(terrain_map))) {
					max_num_rows_equal = num_rows_equal
					rows_above = row_index + 1
					fmt.Println("CHOOSING: ", "number of rows", num_rows_equal, ",rows above", rows_above, "row_upper", row_upper, "row_lower", row_lower)
				}
				fmt.Println("number of rows", num_rows_equal, ",rows above", rows_above)
			}
		}

		// check for column equality
		max_num_columns_equal := 0
		columns_to_the_left := 0
		for column_index := 0; column_index < len(terrain_map[0])-1; column_index++ {
			column_arr := []string {}
			column_arr_next := []string {}
			for row_index := 0; row_index < len(terrain_map); row_index++ {
				column_arr = append(column_arr, terrain_map[row_index][column_index])
				column_arr_next = append(column_arr_next, terrain_map[row_index][column_index+1])
			}
			if reflect.DeepEqual(column_arr, column_arr_next) {
				fmt.Println("column indexes", column_index, column_index+1, "has line of reflection")
				column_upper := column_index-1
				column_lower := column_index+2
				num_columns_equal := 1
				fmt.Println("column_upper", column_upper, "column_lower", column_lower)
				for column_upper >= 0 && column_lower < len(terrain_map[0]) && reflect.DeepEqual(map_column(terrain_map, column_upper), map_column(terrain_map, column_lower)) {
					fmt.Println("comparing columns", map_column(terrain_map, column_upper), map_column(terrain_map, column_lower))
					num_columns_equal += 1
					column_upper -= 1
					column_lower += 1
				}
				if num_columns_equal > max_num_columns_equal && ((column_upper == 0 && column_lower == num_columns_equal+num_columns_equal+1) || (column_upper == len(terrain_map[0])-num_columns_equal-num_columns_equal-1 && column_lower == len(terrain_map[0]))) {
					max_num_columns_equal = num_columns_equal
					columns_to_the_left = column_index + 1
					fmt.Println("CHOOSING: ", "number of columns", num_columns_equal, ",columns to the left", columns_to_the_left, "column_upper", column_upper, "column_lower", column_lower)
				}
				fmt.Println("number of columns", num_columns_equal, ",columns to the left", columns_to_the_left)
			}		
		}
		fmt.Println("----------------------------------")
		if max_num_columns_equal > max_num_rows_equal {
			summarizing_sum += columns_to_the_left
		} else {
			summarizing_sum += rows_above * 100
		}
	}
	fmt.Println(summarizing_sum)
}