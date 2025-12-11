package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toggle_button(button_wiring_schematics string) string {
	return ""
}

func main() {
	input, _ := os.ReadFile("test_input")
	splitted_input := strings.Split(string(input), "\n")

	for _, line := range splitted_input {
		splitted := strings.Split((line), " ")
		button_wiring_schematics := splitted[0]
		indicator_light_diagrams := splitted[1:]
		indicator_light_diagrams_list := [][]int{}

		joltage_requirements_str := indicator_light_diagrams[len(indicator_light_diagrams)-1]
		fmt.Println(joltage_requirements_str)
		// populate indicator light diagrams list
		for _, indicator_light_diagram := range indicator_light_diagrams[:len(indicator_light_diagrams)-1] {
			lights_splitted := strings.Split(indicator_light_diagram[1:len(indicator_light_diagram)-1], ",")
			indicator_lights_line_list := []int{}
			for _, light_str_val := range lights_splitted {
				light_val, _ := strconv.Atoi(light_str_val)
				indicator_lights_line_list = append(indicator_lights_line_list, light_val)
			}
			indicator_light_diagrams_list = append(indicator_light_diagrams_list, indicator_lights_line_list)
		}
		fmt.Println(button_wiring_schematics)
		fmt.Println(indicator_light_diagrams_list)
	}

}
