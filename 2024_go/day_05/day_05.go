package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exists(l []int, e int) bool {
	for _, el := range l {
		if el == e {
			return true
		}
	}
	return false
}

func part_one() {
	input, _ := os.ReadFile("test_input")
	splitted_input := strings.Split(string(input), "\n\n")

	page_ordering_rules := splitted_input[0]
	updates := splitted_input[1]

	page_ordering_rules_list := strings.Split(page_ordering_rules, "\n")
	updates_list := strings.Split(updates, "\n")

	// values must be printed before the key
	page_ordering_map := make(map[int][]int)

	for _, page_ordering_rule := range page_ordering_rules_list {
		rule_split := strings.Split(page_ordering_rule, "|")
		key, _ := strconv.Atoi(rule_split[1])
		value, _ := strconv.Atoi(rule_split[0])
		page_ordering_map[key] = append(page_ordering_map[key], value)
	}
	fmt.Println(page_ordering_map)

	middle_element_sum := 0
	for _, update := range updates_list {
		update_strings_list := strings.Split(update, ",")
		update_int_list := []int{}
		for _, update_string := range update_strings_list {
			update_int, _ := strconv.Atoi(update_string)
			update_int_list = append(update_int_list, update_int)
		}

		obeys_rules := true
		for i := len(update_int_list) - 1; i > 0; i-- {
			key := update_int_list[i]
			for _, value := range update_int_list[:i] {
				if !exists(page_ordering_map[key], value) {
					obeys_rules = false
					break
				}
			}
		}
		if obeys_rules {
			middle_element := update_int_list[(len(update_int_list) / 2)]
			middle_element_sum += middle_element
		}
	}
	fmt.Println(middle_element_sum)

}

func part_two() {

}

func main() {
	//part_one()
	part_two()
}
