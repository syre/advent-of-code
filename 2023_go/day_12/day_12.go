package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"reflect"
	"github.com/mowshon/iterium"
)

func read_ints_from_contiguous_group_text(str string) ([]int, error) {
    var result []int
    splitted := strings.Split(str, ",")
    for _, s := range splitted  {
        x, err := strconv.Atoi(s)
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, nil
}

func find_indexes_with_question_mark(record []string) []int {
	result := []int {}
	for index, _ := range record {
		if record[index] == "?" {
			result = append(result, index)
		}
	}
	return result
}

func record_satisfies_contiguous_group_sizes(record []string, contiguous_group_sizes string) bool {
	existing_contiguous_list, _ := read_ints_from_contiguous_group_text(contiguous_group_sizes)
	damaged_springs_counter := 0
	contiguous_list_found := []int {}

	for _, char := range record {
		if char == "#" {
			damaged_springs_counter += 1
		} else if char == "." {
			if damaged_springs_counter > 0 {
				contiguous_list_found = append(contiguous_list_found, damaged_springs_counter)
				damaged_springs_counter = 0
			}
		}
	}
	if damaged_springs_counter > 0 {
		contiguous_list_found = append(contiguous_list_found, damaged_springs_counter)
		damaged_springs_counter = 0
	}

	return reflect.DeepEqual(contiguous_list_found, existing_contiguous_list)
}

/*
[0 1 2]
[0 2 1]
[1 0 2]
[1 2 0]
[2 0 1]
[2 1 0]

[. . .]
[# . .]
[. # .]
[. . #]
[# # .]
[. # #]
[# . #]
[# # #]
*/

func main() {
	input, _ := os.ReadFile("test_input")
	condition_record_list := strings.Split(string(input), "\n")

	for _, record := range condition_record_list {
		record_splitted := strings.Split(record, " ")
		condition_record := strings.Split(record_splitted[0], "")
		fmt.Println(condition_record)
		//contiguous_group_sizes := string(record_splitted[1])

		//possible_symbols_list := []string {".", "#"}
		//satisfies_number := 0
		question_mark_indexes := find_indexes_with_question_mark(condition_record)
		fmt.Println(question_mark_indexes)
		permutations := iterium.Permutations(question_mark_indexes, len(question_mark_indexes))
		toSlice, _ := permutations.Slice()
		fmt.Println(toSlice)
		fmt.Println(permutations.Count())
		break
		// generate permutations of '.' and '#' from '?'
		// for char_index, char := range condition_record {
		// 	if char != "?" {
		// 		continue
		// 	}
		// 	previous_char := char
		// 	for _, symbol := range possible_symbols_list {
		// 		condition_record[char_index] = symbol
		// 		for other_char_index, other_char := range condition_record {
		// 			other_previous_char := other_char
		// 			if char_index == other_char_index {
		// 				continue
		// 			}
		// 			if other_char != "?" {
		// 				continue
		// 			}
		// 			for _, other_symbol := range possible_symbols_list {
		// 				condition_record[other_char_index] = other_symbol
		// 				fmt.Println("permutation tried:", condition_record)
		// 				satisfies := record_satisfies_contiguous_group_sizes(condition_record, contiguous_group_sizes)
		// 				if satisfies {
		// 					satisfies_number += 1
		// 					fmt.Println("permutation satisfies:", condition_record)
		// 				}
		// 			}
		// 			condition_record[other_char_index] = other_previous_char
		// 		}
		// 		condition_record[char_index] = previous_char
		// 	}
		// 	fmt.Println(satisfies_number)
		// 	satisfies_number = 0
		// }
		//fmt.Println(condition_record_list)
	}
}