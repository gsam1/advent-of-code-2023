package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	fmt.Println(e)
	if e != nil {
		panic(e)
	}
}

func isNumber(input string) bool {
	isNum := true
	for _, char := range input {
		fmt.Println(char)
		if !unicode.IsNumber(char) {
			isNum = false
		}
	}
	return isNum
}

func parseNumbers(inputStr string) []int {
	var numbers []int
	numbersSplit := strings.Split(inputStr, " ")
	for _, numStr := range numbersSplit {
		if isNumber(numStr) {
			num, err := strconv.Atoi(numStr)
			check(err)
			numbers = append(numbers, num)
		}
	}
	fmt.Println(inputStr, numbers)
	return numbers
}

func main() {
	fmt.Println("Day4!!!")
	data, err := os.ReadFile("input.txt")
	check(err)
	dataLines := strings.Split(string(data), "\n")
	for _, row := range dataLines {
		cardNumbersSplit := strings.Split(row, ":")
		numbersSplit := strings.Split(cardNumbersSplit[1], "|")
		numbersGot := numbersSplit[0]
		// numbersWin := numbersSplit[1]
		res := parseNumbers(numbersGot)
		fmt.Println(res)
		// fmt.Println(res, numbersWin)
	}
	// scoring
	// 1 double(1) = 2; double(double(double(1))) -

}
