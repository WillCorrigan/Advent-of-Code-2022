package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	lines := strings.Split(inputResult, "\r\n")

	xIncreases := make(map[int]int)

	x := 1
	cycle := 1
	for _, v := range lines {
		if cycle == 19 || cycle == 20 ||
			cycle == 60 || cycle == 59 ||
			cycle == 100 || cycle == 99 ||
			cycle == 140 || cycle == 139 ||
			cycle == 180 || cycle == 179 ||
			cycle == 220 || cycle == 219 {
			xIncreases[cycle] = x
		}
		if len(strings.Split(v, " ")) == 2 {
			val, _ := strconv.Atoi(strings.Split(v, " ")[1])
			x += val
		}

		if strings.Split(v, " ")[0] == "addx" {
			cycle += 2
		} else {
			cycle += 1
		}
	}
	answer := 0
	for i, v := range xIncreases {
		switch i {
		case 19:
			answer += v * 20
		case 60:
			answer += v * 60
		case 99:
			answer += v * 100
		case 139:
			answer += v * 140
		case 179:
			answer += v * 180
		case 220:
			answer += v * 220
		}
	}
	fmt.Println(xIncreases)
	fmt.Printf("Part 1 Answer: %v\n", answer)
}

func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")
	lines := strings.Split(inputResult, "\r\n")

	var part2Answer []string

	cycle := 1
	currentString := ""
	xPosition := 1

	for _, v := range lines {
		cycleCheck := func(cycleThing int) {
			if (cycleThing)%40 == 0 {
				cycle = 1
				part2Answer = append(part2Answer, currentString)
				currentString = ""
			}
		}

		if strings.Split(v, " ")[0] == "addx" {
			currentString += Draw(cycle-1, xPosition)
			cycle += 1
			cycleCheck(cycle - 1)
			currentString += Draw(cycle-1, xPosition)
			val, _ := strconv.Atoi(strings.Split(v, " ")[1])
			xPosition += val
			cycle += 1
			cycleCheck(cycle - 1)
		} else {
			currentString += Draw(cycle-1, xPosition)
			cycle += 1
			cycleCheck(cycle - 1)
		}
	}

	if err := writeLines(part2Answer, "part2Answer.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func Draw(cycle int, xPosition int) string {
	if cycle == xPosition || cycle == xPosition-1 || cycle == xPosition+1 {
		return "# "
	} else {
		return ". "
	}
}
