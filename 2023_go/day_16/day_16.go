package main

import (
	"os"
	"fmt"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func main() {
	input, _ := os.ReadFile("test_input")
	beam_lists := strings.Split(string(input), "\n")
	beam_map := [][]string {}
	for _, beam_list := range beam_lists {
		entry := strings.Split(beam_list, "")
		beam_map = append(beam_map, entry)
	}
	print_2d_array(beam_map)
}