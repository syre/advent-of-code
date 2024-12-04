package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
	"slices"
)

type Coords struct {
	x int
	y int
}

func remove_duplicates[T string | int](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}


func sum(arr []int) int {
	n := 0
	for _, v := range arr {
		n += v
	}
	return n
}

func printEngineSchematicMatrix(engineSchematicMatrix [][]string) {
	for _, row := range engineSchematicMatrix {
		fmt.Println(row)
	}
}

func partOne() {
	engineSchematicMatrix := [][]string {}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// load in engine schematic in an array of arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charList := []string {}
		text := scanner.Text()
		for _, ch := range text {
			charList = append(charList, string(ch))
		}
		engineSchematicMatrix = append(engineSchematicMatrix, charList)
	}
	printEngineSchematicMatrix(engineSchematicMatrix)

	numberList := []int {}
	symbolList := []string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}
	X := len(engineSchematicMatrix)
	Y := len(engineSchematicMatrix[0])

	numberString := ""
	adjacentToSymbol := false
	for rowPos, row := range engineSchematicMatrix {
		if numberString != "" && adjacentToSymbol {
			if intValue, err := strconv.Atoi(numberString); err == nil {
				numberList = append(numberList, intValue)
			}
		}
		// reset numberString and adjacentToSymbol on new row
		numberString = ""
		adjacentToSymbol = false
		for charPos, char := range row {

			// if char is a number we append to numberString
			// and check adjacent entries (with bounds-checking)
			if _, err := strconv.Atoi(char); err == nil {
				numberString += char

				if (charPos < X-1 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos][charPos+1])) ||
				(charPos > 0 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos][charPos-1])) ||
				(rowPos < X-1 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos+1][charPos])) ||
				(rowPos < Y-1 && charPos < Y-1 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos+1][charPos+1])) ||
				(rowPos < X-1 && charPos > 0 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos+1][charPos-1])) ||
				(rowPos > 0 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos-1][charPos])) ||
				(rowPos > 0 && charPos < Y-1 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos-1][charPos+1])) ||
				(rowPos > 0 && charPos > 0 && !slices.Contains(symbolList, engineSchematicMatrix[rowPos-1][charPos-1])) {
					adjacentToSymbol = true
				}
			} else {
			// if char is not a number
				if numberString != "" && adjacentToSymbol {
					if intValue, err := strconv.Atoi(numberString); err == nil {
						numberList = append(numberList, intValue)
					}
				}
				// reset numberString on non-number
				numberString = ""
				adjacentToSymbol = false
			} 


		}
	}
	fmt.Println(numberList)
	fmt.Println(len(engineSchematicMatrix))
	fmt.Println(sum(numberList))
}

func partTwo() {
	engineSchematicMatrix := [][]string {}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// load in engine schematic in an array of arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charList := []string {}
		text := scanner.Text()
		for _, ch := range text {
			charList = append(charList, string(ch))
		}
		engineSchematicMatrix = append(engineSchematicMatrix, charList)
	}
	printEngineSchematicMatrix(engineSchematicMatrix)

	gearCoordsToNumberMap := map[Coords][]int {}
	X := len(engineSchematicMatrix)
	Y := len(engineSchematicMatrix[0])

	numberString := ""
	adjacentToGearCoords := []Coords {}
	for rowPos, row := range engineSchematicMatrix {
		if numberString != "" && len(adjacentToGearCoords) > 0 {
			if intValue, err := strconv.Atoi(numberString); err == nil {
				for _, val := range adjacentToGearCoords {
					gearCoordsToNumberMap[val] = append(gearCoordsToNumberMap[val], intValue)
				}
			}
		}
		// reset numberString and adjacentToSymbol on new row
		numberString = ""
		adjacentToGearCoords = []Coords {}
		for charPos, char := range row {
			// if char is a number we append to numberString
			// and check adjacent entries
			if _, err := strconv.Atoi(char); err == nil {
				numberString += char

				if (charPos < X-1 && engineSchematicMatrix[rowPos][charPos+1] == "*") {
					coords := Coords{x:rowPos, y:charPos+1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (charPos > 0 && engineSchematicMatrix[rowPos][charPos-1] == "*") {
					coords := Coords{x:rowPos, y:charPos-1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos < X-1 && engineSchematicMatrix[rowPos+1][charPos] == "*") {
					coords := Coords{x:rowPos+1, y:charPos}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos < Y-1 && charPos < Y-1 && engineSchematicMatrix[rowPos+1][charPos+1] == "*") {
					coords := Coords{x:rowPos+1, y:charPos+1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos < X-1 && charPos > 0 && engineSchematicMatrix[rowPos+1][charPos-1] == "*") {
					coords := Coords{x:rowPos+1, y:charPos-1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos > 0 && engineSchematicMatrix[rowPos-1][charPos] == "*") {
					coords := Coords{x:rowPos-1, y:charPos}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos > 0 && charPos < Y-1 && engineSchematicMatrix[rowPos-1][charPos+1] == "*") {
					coords := Coords{x:rowPos-1, y:charPos+1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
				if (rowPos > 0 && charPos > 0 && engineSchematicMatrix[rowPos-1][charPos-1] == "*") {
					coords := Coords{x:rowPos-1, y:charPos-1}
					adjacentToGearCoords = append(adjacentToGearCoords, coords)
				}
			} else {
				// if char is not a number
				if numberString != "" && len(adjacentToGearCoords) > 0 {
					if intValue, err := strconv.Atoi(numberString); err == nil {
						for _, val := range adjacentToGearCoords {
							gearCoordsToNumberMap[val] = append(gearCoordsToNumberMap[val], intValue)
						}
					}
				}
				// reset numberString on non-number
				numberString = ""
				adjacentToGearCoords = []Coords {}
			}
		}
	}
	gearRatioSum := 0
	for _, part_number_list := range gearCoordsToNumberMap {
		unique_part_number_list := remove_duplicates(part_number_list)
		if len(unique_part_number_list) == 2 {
			gearRatioSum += part_number_list[0] * unique_part_number_list[1]
		}
	}
	fmt.Println(gearRatioSum)
}

func main() {
	//partOne()
	partTwo()
}