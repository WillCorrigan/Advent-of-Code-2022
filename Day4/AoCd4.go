package main

import (
	"fmt"
	"strconv"
	"strings"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	Part2()
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	temp := strings.Split(inputResult, "\r\n")

	var total int32

	for _, v := range temp {
		t := strings.Split(v, ",")
		split1 := strings.Split(t[0], "-")
		start1, _ := strconv.Atoi(split1[0])
		end1, _ := strconv.Atoi(split1[1])
		split2 := strings.Split(t[1], "-")
		start2, _ := strconv.Atoi(split2[0])
		end2, _ := strconv.Atoi(split2[1])

		if (start2 >= start1 && end2 <= end1) || (start1 >= start2 && end1 <= end2) {
			total += 1
		}
	}

	fmt.Printf("Part 1 answer is: %v\n", total)
}

func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")

	temp := strings.Split(inputResult, "\r\n")

	var total int32

	for _, v := range temp {
		t := strings.Split(v, ",")
		split1 := strings.Split(t[0], "-")
		start1, _ := strconv.Atoi(split1[0])
		end1, _ := strconv.Atoi(split1[1])
		split2 := strings.Split(t[1], "-")
		start2, _ := strconv.Atoi(split2[0])
		end2, _ := strconv.Atoi(split2[1])

		if (start1 <= start2 && end1 >= start2) || (start2 <= start1 && end2 >= start1) {
			total += 1
		}
	}

	fmt.Printf("Part 2 answer is: %v\n", total)
}
