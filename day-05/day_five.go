package main

import (
	"fmt"
	"os"
	"sort"
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

func sortSlice(input [][]int, sortedIndex int) [][]int {
	newArray := make([][]int, len(input))
	copy(newArray, input)
	sort.Slice(newArray[:], func(i, j int) bool {
		return newArray[i][sortedIndex] < newArray[j][sortedIndex]
	})
	return newArray
}

func printMapContinue(input [][]int, index int) {
	for i, el := range input {
		if i+1 >= len(input) {
			break
		}
		currentMapEnd := el[index] + el[2]
		nextSliceStart := input[i+1][index]
		fmt.Println(el[index], currentMapEnd, nextSliceStart, currentMapEnd == nextSliceStart)
		// fmt.Println("Start range", el[1], "End of range", el[1]+el[2], "Next range start", innerSlice[i][1])
	}
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
		// sort the slices
		sortedSliceBySrc := sortSlice(innerSlice, 1)
		sortedSliceByDest := sortSlice(innerSlice, 0)
		// fmt.Println(sortedSliceBySrc)
		fmt.Println("Source")
		printMapContinue(sortedSliceBySrc, 1)
		fmt.Println("Destination")
		printMapContinue(sortedSliceByDest, 0)
	}

}
