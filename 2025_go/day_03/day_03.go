package main

import (
	"os"
	"strconv"
	"strings"
)

func part_one() {
	input, _ := os.ReadFile("input")
	string_lists := strings.Split(string(input), "\n")
	joltage_sum := 0
	for _, line := range string_lists {
		largest_joltage := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				joltage_string := string([]byte{line[i], line[j]})
				joltage, _ := strconv.Atoi(joltage_string)
				if joltage > largest_joltage {
					largest_joltage = joltage
				}
			}
		}
		joltage_sum += largest_joltage
	}
	println(joltage_sum)
}

func part_two() {
	// turn off 3 permutations
	input, _ := os.ReadFile("test_input")
	string_lists := strings.Split(string(input), "\n")
	joltage_sum := 0
	for _, line := range string_lists {
		largest_joltage := 0
		// oh god what have i done
		for i0 := 0; i0 < len(line)-1; i0++ {
			for i1 := i0 + 1; i1 < len(line); i1++ {
				for i2 := i0 + 2; i2 < len(line); i2++ {
					for i3 := i0 + 3; i3 < len(line); i3++ {
						for i4 := i0 + 4; i4 < len(line); i4++ {
							for i5 := i0 + 5; i5 < len(line); i5++ {
								for i6 := i0 + 6; i6 < len(line); i6++ {
									for i7 := i0 + 7; i7 < len(line); i7++ {
										for i8 := i0 + 8; i8 < len(line); i8++ {
											for i9 := i0 + 9; i9 < len(line); i9++ {
												for i10 := i0 + 10; i10 < len(line); i10++ {
													joltage_string := string([]byte{
														line[i0],
														line[i1],
														line[i2],
														line[i3],
														line[i4],
														line[i5],
														line[i6],
														line[i7],
														line[i8],
														line[i9],
														line[i10],
													})
													joltage, _ := strconv.Atoi(joltage_string)
													if joltage > largest_joltage {
														largest_joltage = joltage
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		joltage_sum += largest_joltage
	}
	println(joltage_sum)
}

func main() {
	//part_one()
	part_two()
}
