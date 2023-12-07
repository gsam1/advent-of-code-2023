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

func checkLeftRight(inputDataGrid []string, r int, c int, direction string) string {
	numberStr := ""
	sourceNum := inputDataGrid[r][c]
	increment := 0

	lr := 1
	if direction == "left" {
		lr = -1
	}

	for unicode.IsNumber(rune(sourceNum)) {
		numberStr += string(sourceNum)
		sourceNum = inputDataGrid[r][c+lr*(1+increment)]
		increment += 1
	}

	return numberStr
}

func checkBottomTop(inputDataGrid []string, r int, c int, direction string) string {
	numberStr := ""
	sourceNum := inputDataGrid[r][c]
	increment := 0

	lc := 1
	if direction == "left" {
		lc = -1
	}

	for unicode.IsNumber(rune(sourceNum)) {
		numberStr += string(sourceNum)
		sourceNum = inputDataGrid[r+lc*(1+increment)][c]
		increment += 1
	}

	return numberStr
}

func scanGrid(inputDataGrid []string, r int, c int) int {
	left := inputDataGrid[r][c-1]
	right := inputDataGrid[r][c+1]
	numbers := map[string]string{}

	if r == 1 {
		bottom := inputDataGrid[r+1][c]
		dlbottom := inputDataGrid[r+1][c-1]
		drbottom := inputDataGrid[r+1][c+1]

		// check left
		if unicode.IsNumber(rune(left)) {
			numbers[strconv.Itoa(r)+strconv.Itoa(c-1)] = checkLeftRight(inputDataGrid, c, r, "left")
		}
		// check right
		if unicode.IsNumber(rune(right)) {
			numbers[strconv.Itoa(r)+strconv.Itoa(c-1)] = checkLeftRight(inputDataGrid, c, r, "right")
		}
		// check r + 1
		if unicode.IsNumber(rune(bottom)) {
			numbers[strconv.Itoa(r+1)+strconv.Itoa(c+1)] = numberStr
		}

	} else if r == len(inputDataGrid) {
		// check left
		// check right
		// check r - 1
	} else {
		// check left
		// check right
		// check r - 1
		// check r + 1

	}

}

// grid stuff
func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	initialDataGrid := strings.Split(string(data), "\n")
	fmt.Printf("Rows: %d Columns %d \n", len(initialDataGrid), len(initialDataGrid[0]))
	// fmt.Println(string(initialDataGrid[73][1]))

	for r, row := range initialDataGrid {
		for c, col := range row {
			if !unicode.IsNumber(col) && string(col) != "." {
				scanGrid(initialDataGrid, r, c)
			}
		}
	}
}
