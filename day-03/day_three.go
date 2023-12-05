package main

import (
	"fmt"
	"os"
	"strings"
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

	for r, row := range strings.Split(string(data), "\n") {
		for c, col := range row {
			fmt.Println(c, r, string(col))
		}
	}
}
