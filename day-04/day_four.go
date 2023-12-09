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

func main() {
	fmt.Println("Day4!!!")
	data, err := os.ReadFile("input.txt")
	check(err)
	dataLines := strings.Split(string(data), "\n")
	totalPoints := 0
	for _, row := range dataLines {
		cardNumbersSplit := strings.Split(row, ":")
		numbersSplit := strings.Split(cardNumbersSplit[1], "|")
		numbersGot := strings.TrimSpace(numbersSplit[0])
		numbersWin := strings.TrimSpace(numbersSplit[1])
		intGot := parseNumbers(numbersGot)
		intWin := parseNumbers(numbersWin)
		intNumWins := getNumWin(intGot, intWin)
		intPoints := calcPoints(intNumWins)
		// fmt.Println(row, intNumWins, intPoints)
		totalPoints += intPoints
	}
	fmt.Println("Result", totalPoints)
	// scoring
	// 1 double(1) = 2; double(double(double(1))) -

}
