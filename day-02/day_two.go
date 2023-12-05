package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// max contained only 12 red cubes, 13 green cubes, and 14 blue cubes

func checkCondition(gameMap map[string]int) bool {
	passed := true

	if gameMap["red"] > 12 {
		passed = false
	}

	if gameMap["green"] > 13 {
		passed = false
	}

	if gameMap["blue"] > 14 {
		passed = false
	}

	return passed
}

func findSumGames(inputTextArray []string) int {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sumIds := 0

	for _, line := range inputTextArray {
		lineSplit := strings.Split(line, ":")
		gameIdStr := strings.Split(lineSplit[0], " ")[1]
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			panic(err)
		}

		cubeSets := strings.Split(lineSplit[1], ";")
		// gameMap := map[string]int{
		// 	"gameId": gameId,
		// }
		valid := true
		for _, sets := range cubeSets {
			splitCubes := strings.Split(sets, ",")
			for _, cube := range splitCubes {
				cubeSplit := strings.Split(strings.TrimSpace(cube), " ")
				numCubesStr := strings.TrimSpace(cubeSplit[0])
				numCubes, err := strconv.Atoi(numCubesStr)
				if err != nil {
					panic(err)
				}
				cubeColor := cubeSplit[1]
				if numCubes > maxCubes[cubeColor] {
					valid = false
				}

			}
		}
		if valid {
			sumIds += gameId
		}

		// fmt.Println(gameMap, checkCondition(gameMap))

	}
	return sumIds
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputTextArray := strings.Split(inputText, "\n")
	sumIds := findSumGames(inputTextArray)
	fmt.Println(sumIds)

}
