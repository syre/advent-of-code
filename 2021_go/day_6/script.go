package main
import (
	"os"
	"strings"
	"fmt"
	"strconv"
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

func PrintFishMap(fish_map map[int]int) {
	for key, value := range fish_map {
		for i := 0; i < value; i++ {
			fmt.Print(key, ",")
		}
	}
	fmt.Println()
}

func main() {
	input, _ := os.ReadFile("input")
	number_strings := strings.Split(string(input), ",")
	numbers, _ := ReadIntsFromStringArray(number_strings)
	fmt.Println(len(numbers))
	// Part 1 https://adventofcode.com/2021/day/6
	for i := 0; i < 18; i++ {
		new_fish_arr := []int{}
		for index, _ := range numbers {
			if numbers[index] == 0 {
				new_fish_arr = append(new_fish_arr, 8)
				numbers[index] = 7
			}
			numbers[index] -= 1
		}
		for _, new_fish := range new_fish_arr {
			numbers = append(numbers, new_fish)
		}
	}
	fmt.Println(len(numbers))

	// Part 2 https://adventofcode.com/2021/day/6
	input, _ = os.ReadFile("input")
	number_strings = strings.Split(string(input), ",")
	numbers, _ = ReadIntsFromStringArray(number_strings)
	fish_map := map[int]int {
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
	for _, number := range numbers {
		fish_map[number] += 1
	}
	fmt.Println(fish_map)
	for i := 0; i < 256; i++ {
		fish_to_add := 0
		fish_to_reset := 0
		// use for-loop as iteration order of maps is random.
		for i := 0; i <= 8; i++ {
			if i == 0 {
				fish_to_add = fish_map[i]
				fish_to_reset = fish_map[i]
				fish_map[i] = 0

			} else {
				fish_map[i-1] = fish_map[i]
			}
		}
		fish_map[8] = fish_to_add
		fish_map[6] += fish_to_reset
		//fmt.Println("day", i)
		//PrintFishMap(fish_map)
	}
	
	num_fish := 0
	for _, value := range fish_map {
		num_fish += value
	}
	fmt.Println(num_fish)
}