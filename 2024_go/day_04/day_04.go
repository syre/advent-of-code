package main

import (
	"fmt"
	"os"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func contains_horizontal_xmas(arr [][]string, i int, j int) bool {
	if len(arr[i])-j < 4 {
		return false
	}
	return arr[i][j] == "X" && arr[i][j+1] == "M" && arr[i][j+2] == "A" && arr[i][j+3] == "S"
}

func contains_vertical_xmas(arr [][]string, i int, j int) bool {
	if len(arr[j])-i < 4 {
		return false
	}
	return arr[i][j] == "X" && arr[i+1][j] == "M" && arr[i+2][j] == "A" && arr[i+3][j] == "S"
}

func contains_diagonal_xmas(arr [][]string, i int, j int) bool {
	return arr[i][j] == "X" && arr[i-1][j+1] == "M" && arr[i-2][j+2] == "A" && arr[i-3][j+3] == "S"
}

func contains_backwards_xmas(arr [][]string, i int, j int) bool {
	return arr[i][j] == "X" && arr[i][j-1] == "M" && arr[i][j-2] == "A" && arr[i][j-3] == "S"
}

func main() {
	input, _ := os.ReadFile("test_input")
	xmas_lists := strings.Split(string(input), "\n")
	xmas_2d_arr := [][]string{}
	// populate xmas_2d_arr
	for _, xmas_list := range xmas_lists {
		if len(xmas_list) == 0 {
			continue
		}
		entry := strings.Split(xmas_list, "")
		xmas_2d_arr = append(xmas_2d_arr, entry)
	}
	print_2d_array(xmas_2d_arr)

	//xmas_count := 0
	// horizontal test
	fmt.Println(xmas_2d_arr[0][5])
	fmt.Println(contains_horizontal_xmas(xmas_2d_arr, 0, 5))
	// vertical test
	fmt.Println(xmas_2d_arr[3][9])
	fmt.Println(contains_vertical_xmas(xmas_2d_arr, 3, 9))
	// diagonal test
	fmt.Println(xmas_2d_arr[9][5])
	fmt.Println(contains_diagonal_xmas(xmas_2d_arr, 9, 5))
	// written backwards test
	fmt.Println(xmas_2d_arr[1][4])
	fmt.Println(contains_backwards_xmas(xmas_2d_arr, 1, 4))

}
