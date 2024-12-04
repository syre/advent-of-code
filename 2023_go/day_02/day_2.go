package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func partOne() {
	// form: map(gameId -> list(map(color -> number of cubes)))
	gamesMap := make(map[int][]map[string]int)
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		gameIdAndGamesSplitted := strings.Split(text, ":")
		gameIdStr := strings.Trim(gameIdAndGamesSplitted[0][4:], " ")
		gameId, _ := strconv.Atoi(gameIdStr)
		gamesString := gameIdAndGamesSplitted[1]
		gamesStringSplitted := strings.Split(strings.Trim(gamesString, " "), ";")

		gamesMap[gameId] = []map[string]int {}
		fmt.Println("Game:", gameId)
		fmt.Println("Game string:", gamesString)
		fmt.Println("-----------------")
		for _, gameString := range gamesStringSplitted {
			gameMap := map[string]int {
				"blue": 0,
				"green": 0,
				"red": 0,
			}
			gameStringSplitted := strings.Split(gameString, ",")
			fmt.Println(gameStringSplitted)
			for _, colorNumberString := range gameStringSplitted {
				colorNumberStringSplitted := strings.Split(strings.Trim(colorNumberString, " "), " ")
				number, _ := strconv.Atoi(colorNumberStringSplitted[0])
				color := colorNumberStringSplitted[1]
				gameMap[color] = number
			}
			gamesMap[gameId] = append(gamesMap[gameId], gameMap)
		}
	}
	sum := 0
	for gameId, gameList := range gamesMap {
		possible := true
		for _, gameMap := range gameList {
			if gameMap["red"] > 12 {
				possible = false
			}
			if gameMap["green"] > 13 {
				possible = false
			}
			if gameMap["blue"] > 14 {
				possible = false
			}
		}
		if possible {
			sum += gameId
		}
	}
	fmt.Println(sum)
}

func partTwo() {
	// form: map(gameId -> list(map(color -> number of cubes)))
	gamesMap := make(map[int][]map[string]int)
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		gameIdAndGamesSplitted := strings.Split(text, ":")
		gameIdStr := strings.Trim(gameIdAndGamesSplitted[0][4:], " ")
		gameId, _ := strconv.Atoi(gameIdStr)
		gamesString := gameIdAndGamesSplitted[1]
		gamesStringSplitted := strings.Split(strings.Trim(gamesString, " "), ";")

		gamesMap[gameId] = []map[string]int {}
		fmt.Println("Game:", gameId)
		fmt.Println("Game string:", gamesString)
		fmt.Println("-----------------")
		for _, gameString := range gamesStringSplitted {
			gameMap := map[string]int {
				"blue": 0,
				"green": 0,
				"red": 0,
			}
			gameStringSplitted := strings.Split(gameString, ",")
			fmt.Println(gameStringSplitted)
			for _, colorNumberString := range gameStringSplitted {
				colorNumberStringSplitted := strings.Split(strings.Trim(colorNumberString, " "), " ")
				number, _ := strconv.Atoi(colorNumberStringSplitted[0])
				color := colorNumberStringSplitted[1]
				gameMap[color] = number
			}
			gamesMap[gameId] = append(gamesMap[gameId], gameMap)
		}
	}
	sum := 0
	for _, gameList := range gamesMap {
		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, gameMap := range gameList {
			if gameMap["red"] > minRed {
				minRed = gameMap["red"]
			}
			if gameMap["green"] > minGreen {
				minGreen = gameMap["green"]
			}
			if gameMap["blue"] > minBlue {
				minBlue = gameMap["blue"]
			}
		}
		sum += minRed * minGreen * minBlue
	}
	fmt.Println(sum)
}

func main() {
	partOne()
	partTwo()
}