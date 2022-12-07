package main

import (
	"fmt"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	Part2()
}

func SlidingWindow(inputResult string, setSize int) {
	indx := setSize
	for i := range inputResult {
		isUnique := make(map[rune]bool)
		slidingWindow := inputResult[i : i+setSize]
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

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	SlidingWindow(inputResult, 4)
}

// it's just part 1 but with 14 lol??
func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")

	SlidingWindow(inputResult, 14)
}
