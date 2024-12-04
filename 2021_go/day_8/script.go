package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func main() {
	// Initialization.
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input),"\n")
	// Part 1 https://adventofcode.com/2021/day/8
	num_ones := 0
	num_fours := 0
	num_sevens := 0
	num_eights := 0
	for _, line := range lines {
		output_strings := strings.Split(line, "|")[1]
		output_str_arr := strings.Split(output_strings, " ")
		fmt.Println(output_str_arr)

		for _, output_str := range output_str_arr {
			if len(output_str) == 2 { // this is a 1
				num_ones += 1
			} else if len(output_str) == 4 { // this is a 4
				num_fours += 1
			} else if len(output_str) == 3 { // this is a 7
				num_sevens += 1
			} else if len(output_str) == 7 { // this is a 8
				num_eights += 1
			}
		}
	}
	fmt.Println(num_ones + num_fours + num_sevens + num_eights)

	// Part 2 https://adventofcode.com/2021/day/8
	char_map := map[string]string {
	    SortString("acedgfb"): "8",
	    SortString("cdfbe"): "5",
	    SortString("gcdfa"): "2",
	    SortString("fbcad"): "3",
	    SortString("dab"): "7",
	    SortString("cefabd"): "9",
	    SortString("cdfgeb"): "6",
	    SortString("eafb"): "4",
	    SortString("cagedb"): "0",
	    SortString("ab"): "1",
	}
	output_total := 0
	for _, line := range lines {
		output_strings := strings.Split(line, "|")[1]
		output_str_arr := strings.Split(output_strings, " ")
		output_number_string := ""
		for _, output_str := range output_str_arr {
			if len(output_str) == 2 { // this is a 1
				output_number_string += "1"
			} else if len(output_str) == 4 { // this is a 4
				output_number_string += "4"
			} else if len(output_str) == 3 { // this is a 7
				output_number_string += "7"
			} else if len(output_str) == 7 { // this is a 8
				output_number_string += "8"
			} else {
				output_char := char_map[SortString(output_str)]
				output_number_string += output_char
			}
			fmt.Println(output_str,"->",output_number_string)
		}
		output_number, _ := strconv.Atoi(output_number_string)
		fmt.Println(output_number)
		output_total += output_number
	}
	fmt.Println(output_total)
}