package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Seed struct {
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func strToInt(inputStr string) int {
	num, err := strconv.Atoi(inputStr)
	check(err)
	return num
}

func parseInnerSplit(input []string) [][]int {
	innerSlice := make([][]int, len(input))
	for i, line := range input {
		lineInt := make([]int, 3)
		for ii, el := range strings.Fields(line) {
			lineInt[ii] = strToInt(el)
		}
		innerSlice[i] = lineInt
	}

	return innerSlice
}

func sortSlice(input [][]int) [][]int {

	return input
}

func main() {
	fmt.Println("Day5!!!")
	data, err := os.ReadFile("input.txt")
	check(err)
	dataLines := strings.Split(string(data), "\n")
	seedsStr := strings.Fields(strings.Split(string(dataLines[0]), ":")[1])
	// Initial seeding
	seeds := map[int]Seed{}
	for _, seed := range seedsStr {
		seedId := strToInt(seed)
		if seedId > 0 {
			seed := Seed{
				soil:        0,
				fertilizer:  0,
				water:       0,
				light:       0,
				temperature: 0,
				humidity:    0,
				location:    0,
			}
			seeds[seedId] = seed
		}
	}

	mapSplits := strings.Split(string(data), "\n\n")

	for _, mapping := range mapSplits {
		fmt.Println("--")
		innerMapSplit := strings.Split(mapping, "\n")
		mapName := innerMapSplit[0]
		fmt.Println(mapName)
		innerSlice := parseInnerSplit(innerMapSplit[1:])
		fmt.Println(innerSlice)
		// sort the slices

	}

}
