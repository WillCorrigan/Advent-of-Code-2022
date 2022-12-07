package main

import (
	"fmt"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	Part2()
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")
	indx := 4
	for i, _ := range inputResult {
		isUnique := make(map[rune]bool)
		slidingWindow := inputResult[i : i+4]
		for _, slidingWindowValue := range slidingWindow {
			if _, ok := isUnique[slidingWindowValue]; ok {
				isUnique[slidingWindowValue] = false
			} else {
				isUnique[slidingWindowValue] = true
			}
		}

		allUnique := true
		for _, value := range isUnique {
			if !value {
				allUnique = false
				break
			}
		}

		if allUnique {
			fmt.Println(indx)
			break
		}

		indx += 1
	}
}

// it's just part 1 but with 14 lol??
func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")
	indx := 14
	for i, _ := range inputResult {
		isUnique := make(map[rune]bool)
		slidingWindow := inputResult[i : i+14]
		for _, slidingWindowValue := range slidingWindow {
			if _, ok := isUnique[slidingWindowValue]; ok {
				isUnique[slidingWindowValue] = false
			} else {
				isUnique[slidingWindowValue] = true
			}
		}

		allUnique := true
		for _, value := range isUnique {
			if !value {
				allUnique = false
				break
			}
		}

		if allUnique {
			fmt.Println(indx)
			break
		}

		indx += 1
	}
}
