package main_two

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

func onlyIntegers(inputArray []string) []string {
	numStringArray := make([]string, 0, len(inputArray)) // learning to make empty arrays
	numStrMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, line := range inputArray {
		onlyNumbersTmp := ""
		for i, char := range line {
			if unicode.IsNumber(char) {
				onlyNumbersTmp += string(char)
			}

			for j := 1; j < 6; j++ {
				if i+j >= len(line) {
					val, ok := numStrMap[line[i:]]
					if ok {
						onlyNumbersTmp += val
					}
				} else {
					val, ok := numStrMap[line[i:i+j]]
					if ok {
						onlyNumbersTmp += val
					}
				}

			}

		}
		fmt.Printf("Original string %s parsed digits %s \n", line, onlyNumbersTmp)

		if len(onlyNumbersTmp) == 1 {
			onlyNumbersTmp += onlyNumbersTmp
		}

		if len(onlyNumbersTmp) >= 2 {
			onlyNumbersTmp = string(onlyNumbersTmp[0]) + string(onlyNumbersTmp[len(onlyNumbersTmp)-1])
		}
		fmt.Printf("Parsed %s \n", onlyNumbersTmp)
		numStringArray = append(numStringArray, onlyNumbersTmp)
	}
	return numStringArray
}

func main_two() {
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
