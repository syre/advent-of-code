package main
import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func Min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func Abs(x int) int {
	if x < 0 {
		return x*-1
	}
	return x
}

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

func makeRange(from, to int) []int {
	size := Max(from, to) - Min(from, to) + 1
    arr := make([]int, size)
    arr_index := 0
    if from <= to {
    	for i := from; i <= to; i++ {
    		arr[arr_index] = i
    		arr_index += 1
    	}
    } else {
    	for i := from; i >= to; i-- {
    		arr[arr_index] = i
    		arr_index += 1
    	}
    }
    return arr
}

func main() {
	input, _ := os.ReadFile("input")
	number_strings := strings.Split(string(input), ",")
	numbers, _ := ReadIntsFromStringArray(number_strings)
	fmt.Println(len(numbers))
	// Part 1 https://adventofcode.com/2021/day/7
	min_fuel := 9999999999999
	for _, i := range numbers {
		fuel := 0
		for _, j := range numbers {
			fuel += Abs(i-j)
		}
		if fuel < min_fuel {
			fmt.Println("lowest fuel found", fuel, "for number", i)
			min_fuel = fuel
		}
	}
	// Part 2 https://adventofcode.com/2021/day/7
	min_fuel = 9999999999999
	for _, i := range numbers {
		fuel := 0
		for _, j := range numbers {
			range_arr := makeRange(j, i)
			for index, _ := range range_arr {
				fuel += index
			}
		}
		if fuel < min_fuel {
			fmt.Println("lowest fuel found", fuel, "for number", i)
			min_fuel = fuel
		}
	}

}