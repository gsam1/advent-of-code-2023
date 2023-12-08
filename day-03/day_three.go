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

func reverseStr(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func checkLeftRight(inputDataGrid []string, c int, r int, direction string) string {
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

	if direction == "left" {
		numberStr = reverseStr(numberStr)
	}

	return numberStr
}

func scanGrid(inputDataGrid []string, r int, c int, mode string) map[string]string {
	numbers := map[string]string{}

	if c > 0 || c < len(inputDataGrid[r])-1 {
		left := inputDataGrid[r][c-1]
		right := inputDataGrid[r][c+1]

		below := inputDataGrid[r+1][c]
		up := inputDataGrid[r-1][c]
		dlbelow := inputDataGrid[r+1][c-1]
		drbelow := inputDataGrid[r+1][c+1]
		dlup := inputDataGrid[r-1][c-1]
		drup := inputDataGrid[r-1][c+1]

		// check left
		if unicode.IsNumber(rune(left)) {
			index := "r" + strconv.FormatInt(int64(r), 10) + "c" + strconv.FormatInt(int64(c-1), 10)
			// fmt.Println("Left is number.", index)
			numbers[index] = checkLeftRight(inputDataGrid, c-1, r, "left")
		}
		// check right
		if unicode.IsNumber(rune(right)) {
			index := "r" + strconv.FormatInt(int64(r), 10) + "c" + strconv.FormatInt(int64(c+1), 10)
			// fmt.Println("Right is number.", index)
			numbers[index] = checkLeftRight(inputDataGrid, c+1, r, "right")

		}

		if mode == "normal" || mode == "first" {
			if unicode.IsNumber(rune(below)) {
				index := "r" + strconv.FormatInt(int64(r), 10) + "c" + strconv.FormatInt(int64(c), 10)
				// fmt.Println("Below is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c-1, r+1, "left") + string(below) + checkLeftRight(inputDataGrid, c+1, r+1, "right")
			} else if unicode.IsNumber(rune(dlbelow)) {
				index := "r" + strconv.FormatInt(int64(r), 10) + "c" + strconv.FormatInt(int64(c-1), 10)
				// fmt.Println("Diagonal Left Below is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c-1, r+1, "left")
			} else if unicode.IsNumber(rune(drbelow)) {
				index := "r" + strconv.FormatInt(int64(r), 10) + "c" + strconv.FormatInt(int64(c+1), 10)
				// fmt.Println("Diagonal Right Below is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c+1, r+1, "right")
			}
		}

		if mode == "normal" || mode == "last" {
			// check up
			if unicode.IsNumber(rune(up)) {
				index := "r" + strconv.FormatInt(int64(r-1), 10) + "c" + strconv.FormatInt(int64(c-1), 10)
				// fmt.Println("Up is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c-1, r-1, "left") + string(up) + checkLeftRight(inputDataGrid, c+1, r-1, "right")
			} else if unicode.IsNumber(rune(dlup)) {
				index := "r" + strconv.FormatInt(int64(r-1), 10) + "c" + strconv.FormatInt(int64(c-1), 10)
				// fmt.Println("Diagonal Up Left is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c-1, r-1, "left")
			} else if unicode.IsNumber(rune(drup)) {
				index := "r" + strconv.FormatInt(int64(r-1), 10) + "c" + strconv.FormatInt(int64(c+1), 10)
				// fmt.Println("Diagonal Up Right is number.", index)
				numbers[index] = checkLeftRight(inputDataGrid, c+1, r-1, "right")
			}
		}
	}

	return numbers

}

// grid stuff
func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	initialDataGrid := strings.Split(string(data), "\n")
	fmt.Printf("Rows: %d Columns %d \n", len(initialDataGrid), len(initialDataGrid[0]))
	numbers := map[string]string{}
	// fmt.Println(string(initialDataGrid[73][1]))

	for r, row := range initialDataGrid {
		for c, col := range row {
			if !unicode.IsNumber(col) && string(col) != "." {
				// fmt.Printf("!!!!! Condition met: %s r:%d c:%d \n", string(col), r, c)
				mode := "normal"
				if r == 0 {
					mode = "first"
				}
				if r == len(initialDataGrid)-1 {
					mode = "last"
				}
				res_scan := scanGrid(initialDataGrid, r, c, mode)
				for k, v := range res_scan {
					numbers[k] = v
				}
			}
		}
	}

	fmt.Println("----Results")
	for k, v := range numbers {
		fmt.Println(k, v)
	}

}
