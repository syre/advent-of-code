package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
	z int
}

func compute_distance(x coord, y coord) float64 {
	return math.Sqrt(math.Pow(float64(x.x-y.x), 2) + math.Pow(float64(x.y-y.y), 2) + math.Pow(float64(x.z-y.z), 2))
}

func main() {
	input, _ := os.ReadFile("test_input")
	splitted_input := strings.Split(string(input), "\n")
	coord_list := []coord{}
	for _, line := range splitted_input {
		line_splitted := strings.Split(line, ",")
		x, _ := strconv.Atoi(line_splitted[0])
		y, _ := strconv.Atoi(line_splitted[1])
		z, _ := strconv.Atoi(line_splitted[2])
		coord_list = append(coord_list, coord{x: x, y: y, z: z})
	}
	fmt.Println(coord_list)
}
