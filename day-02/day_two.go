package main

import (
	"fmt"
	"os"
	"strings"
)

// max contained only 12 red cubes, 13 green cubes, and 14 blue cubes

func main() {
	fmt.Println("hello")
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputTextArray := strings.Split(inputText, "\n")

	for _, line := range inputTextArray {
		lineSplit := strings.Split(line, ":")
		gameId := strings.Split(lineSplit[0], " ")[1]
		cubeSets := strings.Split(lineSplit[1], ";")
		for _, sets := range cubeSets {
			fmt.Printf("%s: %s \n", gameId, string(sets))

		}
	}
}
