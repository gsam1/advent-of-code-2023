package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// 544433
// 76314915

type Part struct {
	row    int
	column int
	symbol string
	value  []string
}

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

	condition := unicode.IsNumber(rune(sourceNum))
	for condition {
		numberStr += string(sourceNum)
		column := c + lr*(1+increment)
		if column < 0 || column > len(inputDataGrid[r])-1 {
			condition = false
		} else {
			sourceNum = inputDataGrid[r][c+lr*(1+increment)]
			condition = unicode.IsNumber(rune(sourceNum))
			increment += 1
		}
	}

	if direction == "left" {
		numberStr = reverseStr(numberStr)
	}

	return numberStr
}

func scanGridStr(inputDataGrid []string, r int, c int, mode string) []string {
	var numbers []string

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
		numbers = append(numbers, checkLeftRight(inputDataGrid, c-1, r, "left"))
	}
	// check right
	if unicode.IsNumber(rune(right)) {
		numbers = append(numbers, checkLeftRight(inputDataGrid, c+1, r, "right"))

	}

	if mode == "normal" || mode == "first" {
		if unicode.IsNumber(rune(below)) {
			numbers = append(numbers, checkLeftRight(inputDataGrid, c-1, r+1, "left")+string(below)+checkLeftRight(inputDataGrid, c+1, r+1, "right"))
		} else {
			if unicode.IsNumber(rune(dlbelow)) {
				numbers = append(numbers, checkLeftRight(inputDataGrid, c-1, r+1, "left"))
			}
			if unicode.IsNumber(rune(drbelow)) {
				numbers = append(numbers, checkLeftRight(inputDataGrid, c+1, r+1, "right"))
			}
		}
	}

	if mode == "normal" || mode == "last" {
		// check up
		if unicode.IsNumber(rune(up)) {
			numbers = append(numbers, checkLeftRight(inputDataGrid, c-1, r-1, "left")+string(up)+checkLeftRight(inputDataGrid, c+1, r-1, "right"))
		} else {
			if unicode.IsNumber(rune(dlup)) {
				numbers = append(numbers, checkLeftRight(inputDataGrid, c-1, r-1, "left"))
			}
			if unicode.IsNumber(rune(drup)) {
				numbers = append(numbers, checkLeftRight(inputDataGrid, c+1, r-1, "right"))
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
	var numbers []Part
	// fmt.Println(string(initialDataGrid[73][1]))

	fmt.Println("----Check")
	for r, row := range initialDataGrid {
		var partNum [][]string
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
				newPart := Part{row: r, column: c, symbol: string(col), value: scanGridStr(initialDataGrid, r, c, mode)}
				partNum = append(partNum, newPart.value)
				numbers = append(numbers, newPart)
			}
		}
		fmt.Println(row, partNum)
	}

	totalSum := 0
	ratioSum := 0
	for _, v := range numbers {
		tmpSum := 0
		tmpRatio := 0
		if len(v.value) == 2 {
			a, err := strconv.Atoi(v.value[0])
			check(err)
			b, err := strconv.Atoi(v.value[1])
			check(err)
			tmpRatio += a * b
		}
		for _, vv := range v.value {
			val, err := strconv.Atoi(vv)
			check(err)
			tmpSum += val
		}
		// fmt.Println(v, tmpSum)
		ratioSum += tmpRatio
		totalSum += tmpSum

	}
	fmt.Println("----Results")
	fmt.Println("Part One: ", totalSum)
	fmt.Println("Part Two: ", ratioSum)
}
