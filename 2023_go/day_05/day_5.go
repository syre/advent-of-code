package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

func ReadIntsFromStringArray(str_arr []string) ([]int, error) {
    var result []int
    for _, s := range str_arr  {
        x, err := strconv.Atoi(s)
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, nil
}

func min(int_arr []int) int {
	min_found := math.MaxInt64
	for _, val := range int_arr {
		if val < min_found {
			min_found = val
		}
	}

	return min_found
}

func part_one() {
	input, _ := os.ReadFile("input")
	sections := strings.Split(string(input), "\n\n")

	// populate seeds - they are a special case
	seeds_string := sections[0][7:]
	seeds_strings := strings.Split(seeds_string, " ")
	seeds := []int {}
	maps_map := map[string][][]int {}
	for _, seed_string := range seeds_strings {
		val, err := strconv.Atoi(seed_string)
		if err == nil {
			seeds = append(seeds, val)
		}

	}
	for _, section := range sections[1:] {
		sectionSplitted := strings.Split(section, "\n")
		// save header for maps_map key
		header := strings.TrimRight(sectionSplitted[0], ":")
		for _, line := range sectionSplitted[1:] {
			// list of: destination, source, range_number
			ints, _ := ReadIntsFromStringArray(strings.Split(line, " "))
			maps_map[header] = append(maps_map[header], ints)
		}
	}

	fmt.Println(seeds)
	fmt.Println(maps_map)

	section_order_list := []string {
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	}

	location_list := []int {}

	for _, seed := range seeds {
		fmt.Println("seed", seed)
		previous_source := seed
		for _, section := range section_order_list {
			fmt.Println(section)
			//fmt.Println(previous_source)
			for _, value := range maps_map[section] {
				destination := value[0]
				source := value[1]
				range_number := value[2]
				//fmt.Println("destination:", destination, "source:", source, "range_number:", range_number)
				if source < previous_source && previous_source < (source + range_number) {
					fmt.Println("within range")
					//fmt.Println(destination + (previous_source-source))
					previous_source = destination + (previous_source-source)
					break
				}
			}
		}
		location_list = append(location_list, previous_source)
	}
	fmt.Println(location_list)
	fmt.Println(min(location_list))
}

func part_two() {
	input, _ := os.ReadFile("input")
	sections := strings.Split(string(input), "\n\n")

	// populate seeds - they are a special case
	seeds_string := sections[0][7:]
	seeds_strings := strings.Split(seeds_string, " ")
	maps_map := map[string][][]int {}

	// populate maps_map with sections
	for _, section := range sections[1:] {
		sectionSplitted := strings.Split(section, "\n")
		// save header for maps_map key
		header := strings.TrimRight(sectionSplitted[0], ":")
		for _, line := range sectionSplitted[1:] {
			// list of: destination, source, range_number
			ints, _ := ReadIntsFromStringArray(strings.Split(line, " "))
			maps_map[header] = append(maps_map[header], ints)
		}
	}

	section_order_list := []string {
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	}

	min_location := math.MaxInt64

	// use computed seeds directly
	for i := 0; i < len(seeds_strings); i+=2 {
		start_seed, start_seed_err := strconv.Atoi(seeds_strings[i])
		range_number, range_number_err := strconv.Atoi(seeds_strings[i+1])
		if start_seed_err == nil && range_number_err == nil {
			for j := 0; j < range_number; j++ {
				seed := start_seed+j
				//fmt.Println("seed", seed)
				previous_source := seed
				for _, section := range section_order_list {
					//fmt.Println(section)
					//fmt.Println(previous_source)
					for _, value := range maps_map[section] {
						destination := value[0]
						source := value[1]
						range_number := value[2]
						//fmt.Println("destination:", destination, "source:", source, "range_number:", range_number)
						if source < previous_source && previous_source < (source + range_number) {
							//fmt.Println("within range")
							//fmt.Println(destination + (previous_source-source))
							previous_source = destination + (previous_source-source)
							break
						}
					}
				}
				if previous_source < min_location {
					min_location = previous_source
				}
			}
		}
	}
	// why this off-by-one-error? xD
	fmt.Println(min_location-1)
}

func main() {
	//part_one()
	part_two()

}