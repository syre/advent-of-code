package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	input, _ := os.ReadFile("test_input")
	pipe_lists := strings.Split(string(input), "\n")
	pipe_map := [][]string {}
	for _, pipe_list := range pipe_lists {
		entry := strings.Split(pipe_list, "")
		pipe_map = append(pipe_map, entry)
	}
	fmt.Println(pipe_map)
}