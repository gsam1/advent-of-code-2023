package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func replaceStringNumbers(inputString string) string {
	numStrMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "zero": "0"}
	for spelledStr, numStr := range numStrMap {
		inputString = strings.Replace(inputString, spelledStr, numStr, 100) // lazy solution
	}
	// inputString = strings.Replace(inputString, "one", "1", 20) // lazy solution
	return inputString
}

func onlyIntegers(inputArray []string) []string {
	numStringArray := make([]string, 0, len(inputArray)) // learning to make empty arrays

	for _, line := range inputArray {
		line = replaceStringNumbers(line)
		onlyNumbersTmp := ""
		for _, char := range line {
			if unicode.IsNumber(char) {
				onlyNumbersTmp += string(char)
			}
		}

		if len(onlyNumbersTmp) == 1 {
			onlyNumbersTmp += onlyNumbersTmp
		}

		if len(onlyNumbersTmp) > 2 {
			onlyNumbersTmp = string(onlyNumbersTmp[0]) + string(onlyNumbersTmp[len(onlyNumbersTmp)-1])
		}

		numStringArray = append(numStringArray, onlyNumbersTmp)
	}
	return numStringArray
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	input_text := string(dat)
	input_array := strings.Split(input_text, "\n")
	// fmt.Println(input_array)
	integer_string_array := onlyIntegers(input_array)
	fmt.Println(integer_string_array)
	total_sum := 0
	for _, numStr := range integer_string_array {
		numInt, err := strconv.Atoi(numStr)
		check(err)
		total_sum += numInt
	}
	fmt.Println(total_sum)
	// map all line to a list without chars

}
