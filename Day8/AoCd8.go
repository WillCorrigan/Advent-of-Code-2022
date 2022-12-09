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

	var treeArr [][]Tree

	for lineIndx, line := range lines {

		var lineTreeArray []Tree

		splitLine := strings.Split(line, "")
		currentMaxHeight, _ := strconv.Atoi(splitLine[0])

		//left to right
		for i, v := range splitLine {
			val, _ := strconv.Atoi(v)

			if lineIndx == 0 || lineIndx == len(lines)-1 || i == 0 || i == len(splitLine)-1 || val > currentMaxHeight {
				tree := Tree{
					Value:   val,
					Visible: true}

				lineTreeArray = append(lineTreeArray, tree)
			} else {
				tree := Tree{
					Value:   val,
					Visible: false}

				lineTreeArray = append(lineTreeArray, tree)
			}
			if val > currentMaxHeight {
				currentMaxHeight = val
			}
		}
		treeArr = append(treeArr, lineTreeArray)
	}

	var currentMaxHeight int
	//right to left
	for i := len(treeArr) - 1; i >= 0; i-- {
		for j := len(treeArr[i]) - 1; j >= 0; j-- {
			if treeArr[i][j].Value > currentMaxHeight {
				treeArr[i][j].Visible = true
				currentMaxHeight = treeArr[i][j].Value
			}
		}
		currentMaxHeight = 0
	}

	transposedArr := transposeTreeArr(treeArr)

	//down to up
	for i := len(transposedArr) - 1; i >= 0; i-- {
		for j := len(transposedArr[i]) - 1; j >= 0; j-- {
			if transposedArr[i][j].Value > currentMaxHeight {
				transposedArr[i][j].Visible = true
				currentMaxHeight = transposedArr[i][j].Value
			}
		}
		currentMaxHeight = 0
	}

	//up to down
	for i := 0; i < len(transposedArr); i++ {
		for j := 0; j < len(transposedArr); j++ {
			if transposedArr[i][j].Value > currentMaxHeight {
				transposedArr[i][j].Visible = true
				currentMaxHeight = transposedArr[i][j].Value
			}
		}
		currentMaxHeight = 0
	}

	var total int
	for _, v := range transposedArr {
		for _, vInner := range v {
			if vInner.Visible {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func transposeTreeArr(arr [][]Tree) [][]Tree {
	xl := len(arr[0])
	yl := len(arr)
	transposedArr := make([][]Tree, xl)
	for i := range transposedArr {
		transposedArr[i] = make([]Tree, yl)
	}
	for i := 0; i < len(arr[0]); i++ {
		for x := 0; x < len(arr); x++ {
			transposedArr[i][x] = arr[x][i]
		}
	}
	return transposedArr
}

// func Part2() {

// }
