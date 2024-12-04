package main

import (
	"bufio"
	"os"
	"log"
	"strconv"
	"fmt"
	"strings"
)

func sum(arr []int) int {
	n := 0
	for _, v := range arr {
		n += v
	}
	return n
}

func partOne() {
	numbersList := []int {}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)
		intRuneList := []string {}
		for _, ch := range text {
			if _, err := strconv.Atoi(string(ch)); err == nil {
				fmt.Println(string(ch))
				intRuneList = append(intRuneList, string(ch))
			}			
		}
		firstRune := intRuneList[0]
		lastRune := intRuneList[len(intRuneList)-1]
		fmt.Println(firstRune, lastRune)
		runeToInt, err := strconv.Atoi(firstRune+lastRune)
		fmt.Println(runeToInt)
		if err == nil {
			numbersList = append(numbersList, runeToInt)
		}
	}
	fmt.Println(numbersList)
	fmt.Println(sum(numbersList))
}

func findNumbersFromString(text string) []string {
	stringNumbersToStrings := map[string]string {
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
	intRuneList := []string {}

	for i := 0; i < len(text); i++ {
		if _, err := strconv.Atoi(string(text[i])); err == nil {
			intRuneList = append(intRuneList, string(text[i]))
			continue
		}

		for key, value := range stringNumbersToStrings {
			if strings.HasPrefix(text[i:], key) {
				intRuneList = append(intRuneList, value)
				continue
			}
		}
	}
	return intRuneList
}

func partTwo() {
	numbersList := []int {}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("-----------------------")
		fmt.Println(text)
		intRuneList := findNumbersFromString(text)
		firstRune := intRuneList[0]
		lastRune := intRuneList[len(intRuneList)-1]
		fmt.Println(firstRune, lastRune)
		runeToInt, err := strconv.Atoi(firstRune+lastRune)
		fmt.Println(runeToInt)
		if err == nil {
			numbersList = append(numbersList, runeToInt)
		}
	}
	fmt.Println(numbersList)
	fmt.Println(sum(numbersList))
}

func main() {
	fmt.Println("Part One:")
	partOne()
	fmt.Println("Part Two:")
	partTwo()
}