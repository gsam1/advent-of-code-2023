package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
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
				if r == 1 {
					fmt.Println("Top row", c, col)
				} else if r == len(initialDataGrid) {
					fmt.Println("Bottom row", c, col)
				}

				// fmt.Println(r, c, string(col))
			}
			// fmt.Println(c, r, string(col))
		}
	}
}
