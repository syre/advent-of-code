package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	pointSum := 0
	for scanner.Scan() {
		text := scanner.Text()

		cardString := strings.Split(text, ":")[1]
		cardStringSplitted := strings.Split(cardString, "|")
		winningNumbersStrings := strings.Split(cardStringSplitted[0], " ")
		myNumbersStrings := strings.Split(cardStringSplitted[1], " ")
		cardPoints := 0

		// we use maps because we want uniques
		winningNumbersMap := map[int]bool {}
		myNumbersMap := map[int]bool {}
		// populate winning numbers and my numbers
		for _, winningNumberString := range winningNumbersStrings {
			winningNumber, err := strconv.Atoi(strings.Trim(winningNumberString, " "))
			if err == nil {
				winningNumbersMap[winningNumber] = true
			}
		}
		for _, myNumber := range myNumbersStrings {
			myNumber, err := strconv.Atoi(strings.Trim(myNumber, " "))
			if err == nil {
				myNumbersMap[myNumber] = true
			}
		}
		// find which are winning numbers

		for winningNumber, _ := range winningNumbersMap {
			for myNumber, _  := range myNumbersMap {
				if myNumber == winningNumber  {
					if cardPoints == 0 {
						cardPoints = 1
					} else {
						cardPoints = cardPoints*2
					}
				}
			}
		}
		pointSum += cardPoints
	}

		fmt.Println(pointSum)
}

func partTwo() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// map of card index -> number of cards
	scratchCardsMap := map[int]int {}
	cardIndex := 0
	for scanner.Scan() {
		text := scanner.Text()

		cardString := strings.Split(text, ":")[1]
		cardStringSplitted := strings.Split(cardString, "|")
		winningNumbersStrings := strings.Split(cardStringSplitted[0], " ")
		myNumbersStrings := strings.Split(cardStringSplitted[1], " ")

		// we use maps because we want uniques
		winningNumbersMap := map[int]bool {}
		myNumbersMap := map[int]bool {}
		// populate winning numbers and my numbers
		for _, winningNumberString := range winningNumbersStrings {
			winningNumber, err := strconv.Atoi(strings.Trim(winningNumberString, " "))
			if err == nil {
				winningNumbersMap[winningNumber] = true
			}
		}
		for _, myNumber := range myNumbersStrings {
			myNumber, err := strconv.Atoi(strings.Trim(myNumber, " "))
			if err == nil {
				myNumbersMap[myNumber] = true
			}
		}

		// initialize scratchCardsMap to 1 if not set
		_, ok := scratchCardsMap[cardIndex]
		if !ok {
			scratchCardsMap[cardIndex] = 1
		}
		// find which are winning numbers
		for i := 0; i < scratchCardsMap[cardIndex]; i++ {
			numMatchingNumbers := 0
			for winningNumber, _ := range winningNumbersMap {
				for myNumber, _  := range myNumbersMap {
					if myNumber == winningNumber  {
						numMatchingNumbers += 1
					}
				}
			}
			for j := 1; j < numMatchingNumbers + 1; j++ {
				// initialize scratchCardsMap on index cardIndex+j to 1 if not set
				_, ok := scratchCardsMap[cardIndex+j]
				if !ok {
					scratchCardsMap[cardIndex+j] = 1
				}
				scratchCardsMap[cardIndex+j] += 1
			}
		}
		cardIndex += 1
	}

	fmt.Println(scratchCardsMap)

	cardSum := 0
	for _, numberOfCards := range scratchCardsMap {
		cardSum += numberOfCards
	}
	fmt.Println(cardSum)
}

func main() {
	//partOne()
	partTwo()
}