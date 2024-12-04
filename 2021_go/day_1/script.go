package main

import (
	"os"
	"strings"
	"fmt"
	"io"
	"bufio"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}

func sum(input []int) int {
    sum := 0
    for _, i := range input {
        sum += i
    }
    return sum
}

func main() {
	input, _ := os.ReadFile("input")
	lines, _ := ReadInts(strings.NewReader(string(input)))

	// Part 1 https://adventofcode.com/2021/day/1
	counter := 0
	for i := 1; i < len(lines); i++ {
		current := lines[i]
		previous := lines[i-1]
		if current > previous {
			counter += 1
		}
	}
	fmt.Println(counter)

	// Part 2 https://adventofcode.com/2021/day/1
	counter = 0
	for i := 1; i < len(lines); i++ {
		current := lines[i:i+3]
		previous := lines[i-1:i+2]

		if sum(current) > sum(previous) {
			counter += 1
		}
	}
	fmt.Println(counter)
}