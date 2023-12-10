package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCardInt(cardName string) int {
	cardStrIdSplit := strings.Fields(cardName)
	cardId, err := strconv.Atoi(cardStrIdSplit[1])
	check(err)
	return cardId
}

func parseNumbers(inputStr string) []int {
	var numbers []int
	numbersSplit := strings.Fields(inputStr)
	for _, numStr := range numbersSplit {
		num, err := strconv.Atoi(numStr)
		check(err)
		numbers = append(numbers, num)
	}
	return numbers
}

func getNumWin(numbersGot []int, numbersWin []int) int {
	numMatch := 0
	for _, g := range numbersGot {
		for _, w := range numbersWin {
			if g == w {
				numMatch += 1
			}
		}
	}
	return numMatch
}

func calcPoints(numMatch int) int {
	if numMatch == 0 {
		return 0
	} else if numMatch == 1 {
		return 1
	} else {
		return 2 * calcPoints(numMatch-1)
	}
}

func calcCardCopySum(cardsMap map[int]int) int {
	totalSum := 0
	for _, v := range cardsMap {
		totalSum += v
	}
	return totalSum
}

func printCardCopySum(cardsMap map[int]int) {
	for k, v := range cardsMap {
		fmt.Println(k, v)
	}
}

func main() {
	fmt.Println("Day4!!!")
	data, err := os.ReadFile("input.txt")
	check(err)
	dataLines := strings.Split(string(data), "\n")
	totalPoints := 0
	cardsMap := map[int]int{}

	// initial seeding of the cardNumberSplit
	for _, row := range dataLines {
		cardNumbersSplit := strings.Split(row, ":")
		cardsMap[getCardInt(cardNumbersSplit[0])] = 1
	}

	for _, row := range dataLines {
		cardNumbersSplit := strings.Split(row, ":")
		currentCardId := getCardInt(cardNumbersSplit[0])
		numbersSplit := strings.Split(cardNumbersSplit[1], "|")
		numbersGot := strings.TrimSpace(numbersSplit[0])
		numbersWin := strings.TrimSpace(numbersSplit[1])
		intGot := parseNumbers(numbersGot)
		intWin := parseNumbers(numbersWin)
		intNumWins := getNumWin(intGot, intWin)
		intPoints := calcPoints(intNumWins)

		increment := cardsMap[currentCardId]
		for i := currentCardId + 1; i <= currentCardId+intNumWins; i++ {
			if i <= len(dataLines) {
				cardsMap[i] += increment
			}
		}

		fmt.Println(getCardInt(cardNumbersSplit[0]), intNumWins, cardsMap)
		// fmt.Println(row, intNumWins, intPoints)
		totalPoints += intPoints
	}

	fmt.Println("Part One Result", totalPoints)
	// fmt.Println(cardsMap)
	// printCardCopySum(cardsMap)
	fmt.Println("Part Two Result", calcCardCopySum(cardsMap))
	// scoring
	// 1 double(1) = 2; double(double(double(1))) -

}
