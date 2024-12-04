package main

import (
	"os"
	"fmt"
	"strings"
	"slices"
)

func part_one() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")

	left_right_instructions := input_splitted[0]
	//fmt.Println(left_right_instructions)

	network_map := map[string][]string {}
	network_list := []string {}
	// populate network map with node -> left node, right node
	for _, node := range input_splitted[1:] {
		if node == "" {
			continue
		}
		current_node := node[0:3]
		left_node := node[7:10]
		right_node := node[12:15]
		network_list = append(network_list, current_node)
		network_map[current_node] = []string {left_node, right_node}
	}
	//fmt.Println(network_map)
	//fmt.Println(network_list)

	index := slices.Index(network_list, "AAA")
	zzz_found := false
	steps := 0
	for zzz_found == false {
		for _, instruction := range left_right_instructions {
			steps += 1
			new_node := ""
			if instruction == 'L' {
				new_node = network_map[network_list[index]][0]
			} else if instruction == 'R' {
				new_node = network_map[network_list[index]][1]
			}
			index = slices.Index(network_list, new_node)
			if network_list[index] == "ZZZ" {
				zzz_found = true
			}
		}
	}
	fmt.Println(steps)
}

func network_list_contain_all_z_nodes(network_list []string, indexes []int) bool {
	for _, index := range indexes {
		if network_list[index][2] != 'Z' {
			return false
		}
	}
	return true
}

func part_two_bruteforce() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")

	left_right_instructions := input_splitted[0]
	fmt.Println(left_right_instructions)

	network_map := map[string][]string {}
	network_list := []string {}
	// populate network map with node -> left node, right node
	for _, node := range input_splitted[1:] {
		if node == "" {
			continue
		}
		current_node := node[0:3]
		left_node := node[7:10]
		right_node := node[12:15]
		network_list = append(network_list, current_node)
		network_map[current_node] = []string {left_node, right_node}
	}
	// populate indexes with nodes ending in A
	indexes := []int {}
	for index, node := range network_list {
		if node[2] == 'A' {
			indexes = append(indexes, index)
		}
	}
	//fmt.Println(network_list)
	all_z_nodes_found := false
	steps := 0
	for all_z_nodes_found == false {
		for _, instruction := range left_right_instructions {
			steps += 1
			for index_index, index := range indexes {
				new_node := ""
				if instruction == 'L' {
					new_node = network_map[network_list[index]][0]
				} else if instruction == 'R' {
					new_node = network_map[network_list[index]][1]
				}
				indexes[index_index] = slices.Index(network_list, new_node)
			}
			if network_list_contain_all_z_nodes(network_list, indexes) {
				all_z_nodes_found = true
			}
		}
	}
	fmt.Println(steps)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func part_two_lcm() {
	input, _ := os.ReadFile("input")
	input_splitted := strings.Split(string(input), "\n")

	left_right_instructions := input_splitted[0]
	fmt.Println(left_right_instructions)

	network_map := map[string][]string {}
	network_list := []string {}
	// populate network map with node -> left node, right node
	for _, node := range input_splitted[1:] {
		if node == "" {
			continue
		}
		current_node := node[0:3]
		left_node := node[7:10]
		right_node := node[12:15]
		network_list = append(network_list, current_node)
		network_map[current_node] = []string {left_node, right_node}
	}
	// populate indexes with nodes ending in A
	indexes := []int {}
	for index, node := range network_list {
		if node[2] == 'A' {
			indexes = append(indexes, index)
		}
	}
	all_z_nodes_found := false
	steps := 0
	steps_map := map[int]int {}
	for all_z_nodes_found == false {
		for _, instruction := range left_right_instructions {
			for index_index, index := range indexes {
				_, ok := steps_map[index_index]
				if ok {
					continue
				}
				if network_list[index][2] == 'Z' {
					steps_map[index_index] = steps
					continue
				}
				new_node := ""
				if instruction == 'L' {
					new_node = network_map[network_list[index]][0]
				} else if instruction == 'R' {
					new_node = network_map[network_list[index]][1]
				}
				indexes[index_index] = slices.Index(network_list, new_node)
			}
			if len(steps_map) == len(indexes) {
				all_z_nodes_found = true
			}
			steps += 1
		}
	}
	values := []int {}
	for _, steps := range steps_map {
		values = append(values, steps)
	}
	fmt.Println(steps_map)
	fmt.Println(LCM(values[0], values[1], values[2:]...))
}

func main() {
	//part_one()
	//part_two_bruteforce()
	// had to check subreddit for this, had no idea about LCM 
	part_two_lcm()
}