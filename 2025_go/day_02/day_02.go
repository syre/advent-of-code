package main

import (
	"os"
	"strconv"
	"strings"
)

func is_invalid_part_one(a int) bool {
	string_integer := strconv.Itoa(a)
	left_part := string_integer[0 : len(string_integer)/2]
	right_part := string_integer[len(string_integer)/2:]

	return left_part == right_part
}
func split_every(s string, n int) []string {
	var chunks []string
	for len(s) > 0 {
		if len(s) < n {
			chunks = append(chunks, s)
			break
		}
		chunks = append(chunks, s[:n])
		s = s[n:]
	}
	return chunks
}
func all_same_strings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func is_invalid_part_two(a int) bool {
	string_integer := strconv.Itoa(a)
	for i := len(string_integer) / 2; i > 0; i-- {
		chunks := split_every(string_integer, i)
		if all_same_strings(chunks) {
			return true
		}
	}
	return false
}

func part_one() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), ",")

	invalid_ids_sum := 0
	for _, line := range string_lists {
		left, right, _ := strings.Cut(line, "-")
		left_number, _ := strconv.Atoi(strings.TrimSpace(left))
		right_number, _ := strconv.Atoi(strings.TrimSpace(right))
		for i := left_number; i < right_number+1; i++ {
			if is_invalid_part_one(i) {
				println(i, "is invalid")
				invalid_ids_sum += i
			}
		}
	}
	println(invalid_ids_sum)
}

func part_two() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), ",")

	invalid_ids_sum := 0
	for _, line := range string_lists {
		left, right, _ := strings.Cut(line, "-")
		left_number, _ := strconv.Atoi(strings.TrimSpace(left))
		right_number, _ := strconv.Atoi(strings.TrimSpace(right))
		println(left_number, right_number)
		for i := left_number; i < right_number+1; i++ {
			if is_invalid_part_two(i) {
				println(i, "is invalid")
				invalid_ids_sum += i
			}
		}
	}
	println(invalid_ids_sum)
}

func main() {
	//part_one()
	part_two()
}
