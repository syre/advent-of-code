package main

import (
	"os"
	"fmt"
	"strings"
)

func print_2d_array(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func get_load(platform_rock_map [][]string) int {
	// get load
    sum := 0
    for row_index, row := range platform_rock_map {
    	num_rounded_rocks := 0
    	for _, entry := range row {
    		if entry == "O" {
    			num_rounded_rocks += 1
    		}
    	}
    	sum += num_rounded_rocks * (len(platform_rock_map) - row_index)
    }
    return sum
}

func part_one() {
	input, _ := os.ReadFile("input")
	platform_rock_lists := strings.Split(string(input), "\n")
	platform_rock_map := [][]string {}
	for _, platform_rock_list := range platform_rock_lists {
		entry := strings.Split(platform_rock_list, "")
		platform_rock_map = append(platform_rock_map, entry)
	}
	print_2d_array(platform_rock_map)

	// tilt north
	for column_index := 0; column_index < len(platform_rock_map[0]); column_index++ {
	    for row_index, _ := range platform_rock_map {
	    	column_pointer := column_index
	        for column_pointer < len(platform_rock_map[0]) {
		        if column_pointer > 0 && platform_rock_map[column_pointer][row_index] == "O" && platform_rock_map[column_pointer-1][row_index] == "." {
		        	platform_rock_map[column_pointer][row_index] = "."
		        	platform_rock_map[column_pointer-1][row_index] = "O"
		        	column_pointer -= 1
		        } else {
		        	column_pointer += 1
		        }
	        }
	    }
	}
    fmt.Println("---------------------")
    fmt.Println("TILT NORTH MAP")
    print_2d_array(platform_rock_map)

	// get load
    sum := 0
    for row_index, row := range platform_rock_map {
    	num_rounded_rocks := 0
    	for _, entry := range row {
    		if entry == "O" {
    			num_rounded_rocks += 1
    		}
    	}
    	sum += num_rounded_rocks * (len(platform_rock_map) - row_index)
    }
    fmt.Println(sum)
}

func tilt_north(platform_rock_map[][]string) [][]string{
	// tilt north
	for column_index := 0; column_index < len(platform_rock_map[0]); column_index++ {
	    for row_index, _ := range platform_rock_map {
	    	column_pointer := column_index
	        for column_pointer < len(platform_rock_map[0]) {
		        if column_pointer > 0 && platform_rock_map[column_pointer][row_index] == "O" && platform_rock_map[column_pointer-1][row_index] == "." {
		        	platform_rock_map[column_pointer][row_index] = "."
		        	platform_rock_map[column_pointer-1][row_index] = "O"
		        	column_pointer -= 1
		        } else {
		        	column_pointer += 1
		        }
	        }
	    }
	}
	return platform_rock_map
}

func tilt_south(platform_rock_map[][]string) [][]string{
	// tilt south
	for row_index, _ := range platform_rock_map {
		for column_index := len(platform_rock_map[0]); column_index > 0; column_index-- {
	    	column_pointer := column_index
	    	//fmt.Println(column_index, row_index)
	    	if column_pointer < len(platform_rock_map[0])-1 {
	    		//fmt.Println("current", platform_rock_map[column_pointer][row_index], "previous", platform_rock_map[column_pointer+1][row_index])
	        }
	        for column_pointer >= 0 {
		        if column_pointer < len(platform_rock_map[0])-1 && platform_rock_map[column_pointer][row_index] == "O" && platform_rock_map[column_pointer+1][row_index] == "." {
		        	//fmt.Println("moved rock")
		        	platform_rock_map[column_pointer][row_index] = "."
		        	platform_rock_map[column_pointer+1][row_index] = "O"
		        	column_pointer += 1
		        } else {
		        	column_pointer -= 1
		        }
	        }
	    }
	}
	return platform_rock_map
}

func tilt_west(platform_rock_map[][]string) [][]string{
	// tilt west
	for column_index := 0; column_index < len(platform_rock_map[0]); column_index++ {
		for row_index := len(platform_rock_map); row_index > 0; row_index-- {
	    	row_pointer := row_index
	    	//fmt.Println(column_index, row_index)
	    	if row_pointer < len(platform_rock_map) {
	    		//fmt.Println("current", platform_rock_map[column_index][row_index], "previous", platform_rock_map[column_index][row_index-1])
	        }
	        for row_pointer > 0 {
		        if row_pointer < len(platform_rock_map) && platform_rock_map[column_index][row_pointer] == "O" && platform_rock_map[column_index][row_pointer-1] == "." {
		        	//fmt.Println("moved rock")
		        	platform_rock_map[column_index][row_pointer] = "."
		        	platform_rock_map[column_index][row_pointer-1] = "O"
		        	row_pointer += 1
		        } else {
		        	row_pointer -= 1
		        }
	        }
	    }
	}
	return platform_rock_map
}

func tilt_east(platform_rock_map[][]string) [][]string{
	// tilt east
	for column_index := 0; column_index < len(platform_rock_map[0]) ; column_index++ {
	for row_index, _ := range platform_rock_map {
	    	row_pointer := row_index
	    	//fmt.Println(column_index, row_index)
	    	if row_pointer > 0 {
	    		//fmt.Println("current", platform_rock_map[column_index][row_index], "previous", platform_rock_map[column_index][row_index-1])
	        }
	        for row_pointer < len(platform_rock_map[0])-1 {
		        if row_pointer >= 0 && platform_rock_map[column_index][row_pointer] == "O" && platform_rock_map[column_index][row_pointer+1] == "." {
		        	//fmt.Println("moved rock")
		        	platform_rock_map[column_index][row_pointer] = "."
		        	platform_rock_map[column_index][row_pointer+1] = "O"
		        	row_pointer -= 1
		        } else {
		        	row_pointer += 1
		        }
	        }
	    }
	}
	return platform_rock_map
}

func part_two() {
	input, _ := os.ReadFile("test_input")
	platform_rock_lists := strings.Split(string(input), "\n")
	platform_rock_map := [][]string {}
	for _, platform_rock_list := range platform_rock_lists {
		entry := strings.Split(platform_rock_list, "")
		platform_rock_map = append(platform_rock_map, entry)
	}
	print_2d_array(platform_rock_map)
	// check patterns?
	patterns_map := map[int]int {}
	// run cycles
	cycles := 1000
	for cycle := 0; cycle < cycles; cycle++ {
		platform_rock_map = tilt_north(platform_rock_map)
		platform_rock_map = tilt_west(platform_rock_map)
		platform_rock_map = tilt_south(platform_rock_map)
		platform_rock_map = tilt_east(platform_rock_map)

		//fmt.Println("CYCLE", cycle + 1)
    	//print_2d_array(platform_rock_map)
    	//fmt.Println("---------------------")
		load := get_load(platform_rock_map)
		_, ok := patterns_map[load]
		if !ok {
			patterns_map[load] = 0
		}
		patterns_map[load] += 1
		fmt.Println(cycle, load)
	}

	fmt.Println(patterns_map)
	// fmt.Println("---------------------")
	//print_2d_array((platform_rock_map)


}

func main() {
	//part_one()
	part_two()
}


// 101026 too high
// 101010 correct - 1000 cycles? a pattern of /1000?