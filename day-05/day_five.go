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
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func convertStrToInt(inputStr string) int {
	num, err := strconv.Atoi(inputStr)
	check(err)
	return num
}

func main() {
	fmt.Println("Day4!!!")
	data, err := os.ReadFile("input.txt")
	check(err)
	dataLines := strings.Split(string(data), "\n")
	seedsStr := strings.Fields(strings.Split(string(dataLines[0]), ":")[1])
	var seeds []Seed
	for _, seed := range seedsStr {
		seedId := convertStrToInt(seed)
		if seedId > 0 {
			seed := Seed{
				seed:        seedId,
				soil:        0,
				fertilizer:  0,
				water:       0,
				light:       0,
				temperature: 0,
				humidity:    0,
				location:    0,
			}
			seeds = append(seeds, seed)
		}
	}
	fmt.Println(seeds)
}