package main

import (
	"fmt"
	"strings"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	Part1Map()
	Part2Map()

	//A|X == rock == 1
	//B|Y == paper == 2
	//C|Z == scissors == 3
	//lose == 0
	//draw == 3
	//win == 6
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	test := strings.Split(inputResult, "\n")

	var newArr [][]string

	for _, v := range test {
		newArr = append(newArr, strings.Split(v, " "))
	}

	var total int32

	for _, v := range newArr {
		switch v[1] {
		case "X":
			total += 1
			switch v[0] {
			case "A":
				total += 3
			case "B":
				total += 0
			case "C":
				total += 6
			}
		case "Y":
			total += 2
			switch v[0] {
			case "A":
				total += 6
			case "B":
				total += 3
			case "C":
				total += 0
			}
		case "Z":
			total += 3
			switch v[0] {
			case "A":
				total += 0
			case "B":
				total += 6
			case "C":
				total += 3
			}
		}
	}

	fmt.Printf("Part 1: %v\n", total)
}

func Part1Map() {

	inputResult := loadTextInput.LoadInput("input.txt")

	test := strings.Split(inputResult, "\n")

	m := make(map[string]int32)

	m["A X"] = 1 + 3
	m["B X"] = 1 + 0
	m["C X"] = 1 + 6
	m["A Y"] = 2 + 6
	m["B Y"] = 2 + 3
	m["C Y"] = 2 + 0
	m["A Z"] = 3 + 0
	m["B Z"] = 3 + 6
	m["C Z"] = 3 + 3

	var total int32

	for _, v := range test {
		total += m[v]
	}

	fmt.Printf("Part 1 Maps: %v\n", total)
}

func Part2Map() {

	//X == lose
	//Y == draw
	//Z == win

	inputResult := loadTextInput.LoadInput("input.txt")

	test := strings.Split(inputResult, "\n")

	m := make(map[string]int32)

	m["A X"] = 3 + 0
	m["A Y"] = 1 + 3
	m["A Z"] = 2 + 6
	m["B X"] = 1 + 0
	m["B Y"] = 2 + 3
	m["B Z"] = 3 + 6
	m["C X"] = 2 + 0
	m["C Y"] = 3 + 3
	m["C Z"] = 1 + 6

	var total int32

	for _, v := range test {
		total += m[v]
	}

	fmt.Printf("Part 2: %v\n", total)
}
