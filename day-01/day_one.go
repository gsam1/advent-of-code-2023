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

func OnlyIntegers(inputArray []string) []string {
	numStringArray := make([]string, 0, len(inputArray)) // learning to make empty arrays
	for _, line := range inputArray {
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
	integer_string_array := OnlyIntegers(input_array)
	// fmt.Println(integer_string_array)
	total_sum := 0
	for _, numStr := range integer_string_array {
		numInt, err := strconv.Atoi(numStr)
		check(err)
		total_sum += numInt
	}
	fmt.Println(total_sum)
	// map all line to a list without chars

}
