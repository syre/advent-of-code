package main

import (
	"strings"
	"os"
	"fmt"
	"strconv"
)

// Maybe not needed.
type Coordinate struct {
    x, y interface{}
}

func Abs(x int) int {
	if x < 0 {
		return x*-1
	}
	return x
}

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

func CountOverlaps(diagram [][]int) int {
	overlap_counter := 0
	for _, row := range diagram {
		for _, number := range row {
			if number > 1 {
				overlap_counter += 1
			}
		}
	}
	return overlap_counter
}


func PrintDiagram(diagram [][]int) {
	for _, row := range diagram {
		fmt.Println(row)
	}
}

func main() {
	input, _ := os.ReadFile("input")
	lines := strings.Split(string(input),"\n")
	diagram := make([][]int, 1000)
	for i := range diagram {
	    diagram[i] = make([]int, 1000)
	}
	// Part 1 https://adventofcode.com/2021/day/5
	for _, line := range lines {
		coordinate_arr := strings.Replace(string(line), " -> ", ",", 1)
		coordinates := strings.Split(coordinate_arr, ",")
		x1, _ := strconv.Atoi(coordinates[0])
		y1, _ := strconv.Atoi(coordinates[1])
		x2, _ := strconv.Atoi(coordinates[2])
		y2, _ := strconv.Atoi(coordinates[3])
		if x1 == x2 || y1 == y2 {
			x_range := makeRange(Min(x1, x2), Max(x1, x2))
			y_range := makeRange(Min(y1, y2), Max(y1, y2))

			for _, y := range x_range {
				for _, x := range y_range {
					diagram[x][y] += 1
				}
			}
		}
	}
	overlap_counter := CountOverlaps(diagram)
	fmt.Println(overlap_counter)

	// Part 2 https://adventofcode.com/2021/day/5
	diagram = make([][]int, 1000)
	for i := range diagram {
	    diagram[i] = make([]int, 1000)
	}
	for _, line := range lines {
		coordinate_arr := strings.Replace(string(line), " -> ", ",", 1)
		coordinates := strings.Split(coordinate_arr, ",")
		x1, _ := strconv.Atoi(coordinates[0])
		y1, _ := strconv.Atoi(coordinates[1])
		x2, _ := strconv.Atoi(coordinates[2])
		y2, _ := strconv.Atoi(coordinates[3])
		if x1 == x2 || y1 == y2 {
			x_range := makeRange(Min(x1, x2), Max(x1, x2))
			y_range := makeRange(Min(y1, y2), Max(y1, y2))
			for _, y := range x_range {
				for _, x := range y_range {
					diagram[x][y] += 1
				}
			}
		}
		if y2 - y1 == x1 - x2 || y2 - y1 == x2 - x1 {
			x_range := makeRange(x1, x2)
			y_range := makeRange(y1, y2)
			for index, _ := range x_range {
				diagram[y_range[index]][x_range[index]] += 1
			}
		}
	}
	overlap_counter = CountOverlaps(diagram)
	fmt.Println(overlap_counter)
}

//8,0 0,8

//7,1
//6,2
//5,3
//4,4
//3,5
//2,6
//1,7

//6,4 -> 2,0

//5,3
//4,2
//3,1
//2,0