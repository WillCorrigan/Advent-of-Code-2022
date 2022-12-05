package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	Part2()
}

func findIntersection(input []string) rune {
	hash := make(map[rune]bool)
	var runeToReturn rune

	for _, v := range input[0] {
		hash[v] = true
	}

	for _, v := range input[1] {
		if hash[v] {
			runeToReturn = v
		}
	}

	return runeToReturn
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	temp := strings.Split(strings.TrimSpace(inputResult), "\r\n")

	var total int32

	for _, v := range temp {
		r := findIntersection([]string{v[:len(v)/2], v[len(v)/2:]})
		if unicode.IsUpper(r) {
			total += r - 38
		} else {
			total += r - 96
		}
	}

	fmt.Printf("Part 1 answer: %v\n", total)
}

func findBadge(input []string) rune {
	hash := make(map[rune][]bool)
	var runeToReturn rune

	for _, v := range input[0] {
		hash[v] = []bool{true, false, false}
	}

	for _, v := range input[1] {
		if _, found := hash[v]; found {
			hash[v][1] = true
		}
	}

	for _, v := range input[2] {
		if _, found := hash[v]; found {
			hash[v][2] = true
		}
	}

	for k, v := range hash {
		if reflect.DeepEqual(v, []bool{true, true, true}) {
			runeToReturn = k
		}
	}

	return runeToReturn
}

func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")

	temp := strings.Split(strings.TrimSpace(inputResult), "\r\n")

	var total int32

	for i := 0; i < len(temp); i += 3 {
		r := findBadge([]string{temp[i], temp[i+1], temp[i+2]})
		if unicode.IsUpper(r) {
			total += r - 38
		} else {
			total += r - 96
		}
	}

	fmt.Printf("Part 2 answer: %v\n", total)
}
