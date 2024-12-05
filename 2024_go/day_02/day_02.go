package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_report_safe(report string) bool {
	levels := strings.Split(string(report), " ")
	first_level, _ := strconv.Atoi(levels[0])
	second_level, _ := strconv.Atoi(levels[1])

	prev_level := first_level
	// increasing
	safe := true
	if first_level < second_level {
		for _, level := range levels[1:] {
			curr_level, _ := strconv.Atoi(level)
			if curr_level-prev_level > 3 || curr_level-prev_level <= 0 {
				safe = false
			}
			prev_level = curr_level
		}
		// decreasing
	} else {
		for _, level := range levels[1:] {
			curr_level, _ := strconv.Atoi(level)
			if prev_level-curr_level > 3 || prev_level-curr_level <= 0 {
				safe = false
			}
			prev_level = curr_level
		}
	}
	return safe
}

func del_element(s []string, index int) []string {
	return append(s[0:index], s[index+1:]...)
}

func part_one() {
	input, _ := os.ReadFile("test_input")

	report_list := strings.Split(string(input), "\n")

	safe_reports_count := 0
	for _, report := range report_list {
		if report == "" {
			continue
		}
		is_safe := is_report_safe(report)
		if is_safe {
			fmt.Println(report, " is safe")
			safe_reports_count += 1
		}
	}
	fmt.Println(safe_reports_count)

}

func part_two() {
	input, _ := os.ReadFile("test_input")

	report_list := strings.Split(string(input), "\n")

	safe_reports_count := 0
	for _, report := range report_list {
		if report == "" {
			continue
		}
		is_safe := is_report_safe(report)
		if is_safe {
			fmt.Println(report, " is safe")
			safe_reports_count += 1
			continue
		}
		report_splitted := strings.Split(report, " ")
		for i := 0; i < len(report_splitted); i++ {
			// copy report_splitted
			report_copy := make([]string, len(report_splitted))
			copy(report_copy, report_splitted)

			one_elem_deleted_report := strings.Join(del_element(report_copy, i), " ")
			fmt.Println("testing", one_elem_deleted_report)
			is_safe := is_report_safe(one_elem_deleted_report)
			if is_safe {
				fmt.Println(report, " is safe")
				safe_reports_count += 1
				break
			}
		}
	}
	fmt.Println(safe_reports_count)
}

func main() {
	//part_one()
	part_two()
}
