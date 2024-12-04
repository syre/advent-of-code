package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

type LabelFocalPair struct {
	label string
	focal_length int
}

func getLabelIndex(arr []LabelFocalPair, label string) int {
	for index, el := range arr {
		if el.label == label {
			return index
		}
	}
	return -1
}

func part_one() {
	input, _ := os.ReadFile("input")
	string_list := strings.Split(string(input), ",")

	fmt.Println(string_list)

	value_sum := 0
	current_value := 0
	for _, string := range string_list {
		for _, char := range string {
			current_value += int(char)
			current_value *= 17
			current_value = current_value % 256
		}
		value_sum += current_value
		current_value = 0
	}
	fmt.Println(value_sum)
}

func part_two() {
	input, _ := os.ReadFile("input")
	string_list := strings.Split(string(input), ",")

	box_to_label_focal_length_map := map[int][]LabelFocalPair {}
	for _, string_obj := range string_list {
		minus_index := strings.IndexRune(string_obj, '-')
		equals_index := strings.IndexRune(string_obj, '=')
		current_value := 0
		number := 0

		if equals_index != -1 {
			number, _ = strconv.Atoi(string_obj[equals_index+1:])
			fmt.Println(number)
			for _, char := range string_obj[:equals_index] {
				current_value += int(char)
				current_value *= 17
				current_value = current_value % 256
			}
			_, ok := box_to_label_focal_length_map[current_value]
			if !ok {
				box_to_label_focal_length_map[current_value] = []LabelFocalPair {}
			}
			label_index := getLabelIndex(box_to_label_focal_length_map[current_value], string_obj[:equals_index])
			if label_index != -1 {
				box_to_label_focal_length_map[current_value][label_index].focal_length = number
			} else {
				box_to_label_focal_length_map[current_value] = append(box_to_label_focal_length_map[current_value], LabelFocalPair{label:string_obj[:equals_index], focal_length:number})
			}
		}

		if minus_index != -1 {
			number, _ = strconv.Atoi(string_obj[minus_index+1:])
			for _, char := range string_obj[:minus_index] {
				current_value += int(char)
				current_value *= 17
				current_value = current_value % 256
			}
			_, ok := box_to_label_focal_length_map[current_value]
			if !ok {
				box_to_label_focal_length_map[current_value] = []LabelFocalPair {}
				continue
			}
			// delete element
			delete_index := getLabelIndex(box_to_label_focal_length_map[current_value], string_obj[:minus_index])
			if delete_index == -1 {
				continue
			}
			box_to_label_focal_length_map[current_value] = append(box_to_label_focal_length_map[current_value][:delete_index], box_to_label_focal_length_map[current_value][delete_index+1:]...)
		}
	}
	fmt.Println(box_to_label_focal_length_map)
	sum := 0
	for box_number, label_focal_map := range box_to_label_focal_length_map {
		for index, label_focal_element := range label_focal_map {
			fmt.Println("box number:", 1+box_number, "index:", index+1, "focal length:", label_focal_element.focal_length)
			sum += (1+box_number)*(index+1)*label_focal_element.focal_length
		}
	}
	fmt.Println(sum)
}

func main() {
	//part_one()
	part_two()
}