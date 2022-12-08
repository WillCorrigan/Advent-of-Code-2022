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

type Tree struct {
	Value   int
	Visible bool
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")
	lines := strings.Split(inputResult, "\r\n")

	var inputArr [][]*Tree

	for _, line := range lines {
		var tmpArr []*Tree
		splt := strings.Split(line, "")
		for _, v := range splt {
			val, err := strconv.Atoi(v)
			if err == nil {
				test := &Tree{
					Value:   val,
					Visible: false}

				tmpArr = append(tmpArr, test)
			}
		}
		inputArr = append(inputArr, tmpArr)
	}

	for _, v := range inputArr[0] {
		v.Visible = true
	}
	for _, v := range inputArr[len(inputArr)-1] {
		v.Visible = true
	}
	for i := 0; i < len(inputArr); i++ {
		inputArr[i][0].Visible = true
	}

	for i := 0; i < len(inputArr); i++ {
		inputArr[i][len(inputArr)-1].Visible = true
	}

	fmt.Println(inputArr)
}

// func Part2() {

// }
