package main

import (
	"fmt"
	"strconv"
	"strings"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	Part1()
	// Part2()
}

type Coordinate struct {
	X int
	Y int
}

type Head struct {
	X int
	Y int
}

type Tail struct {
	X int
	Y int
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")
	lines := strings.Split(inputResult, "\r\n")

	visitedCoords := make(map[Coordinate]bool)

	head := &Head{
		X: 0,
		Y: 0,
	}

	tail := &Tail{
		X: 0,
		Y: 0,
	}

	visitedCoords[Coordinate{X: 0, Y: 0}] = true

	for _, v := range lines {

		split := strings.Split(v, " ")

		direction, strAmount := split[0], split[1]

		amount, _ := strconv.Atoi(strAmount)

		for i := 1; i <= amount; i++ {
			switch direction {
			case "U":
				head.Y += 1
				//up shift
				if head.Y == tail.Y {
					continue
				}
				if head.Y-tail.Y == 2 {
					//straight up
					tail.Y += 1
					//check if diagonal right
					if head.X-tail.X == 1 {
						tail.X += 1
						//check if diagonal left
					} else if head.X-tail.X == -1 {
						tail.X -= 1
					}
				}
				visitedCoords[Coordinate{
					X: tail.X,
					Y: tail.Y,
				}] = true
			case "R":
				head.X += 1
				if head.X-tail.X == 0 {
					continue
				}
				if head.X-tail.X == 2 {
					tail.X += 1
					if head.Y-tail.Y == 1 {
						tail.Y += 1
					} else if head.Y-tail.Y == -1 {
						tail.Y -= 1
					}
				}
				visitedCoords[Coordinate{
					X: tail.X,
					Y: tail.Y,
				}] = true
			case "D":
				head.Y -= 1
				if head.Y == tail.Y {
					continue
				}
				if head.Y-tail.Y == -2 {
					//straight down
					tail.Y -= 1
					//check if diagonal right
					if head.X-tail.X == 1 {
						tail.X += 1
						//check if diagonal left
					} else if head.X-tail.X == -1 {
						tail.X -= 1
					}
				}
				visitedCoords[Coordinate{X: tail.X, Y: tail.Y}] = true
			case "L":
				head.X -= 1
				if head.X-tail.X == 0 {
					continue
				}
				if head.X-tail.X == -2 {
					tail.X -= 1
					if head.Y-tail.Y == 1 {
						tail.Y += 1
					} else if head.Y-tail.Y == -1 {
						tail.Y -= 1
					}
				}
				visitedCoords[Coordinate{
					X: tail.X,
					Y: tail.Y,
				}] = true
			}
		}
	}

	fmt.Printf("Part 1 Answer: %v", len(visitedCoords))
}
