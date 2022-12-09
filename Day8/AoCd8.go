package main

import (
	"fmt"
	"strconv"
	"strings"

	loadTextInput "github.com/WillCorrigan/AdventOfCode2022/LoadTextInput"
)

func main() {
	// Part1()
	Part2()
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

type ScenicTree struct {
	Value int
	North int
	East  int
	South int
	West  int
}

func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")
	lines := strings.Split(inputResult, "\r\n")

	var treeArr [][]ScenicTree
	var formattedArray [][]int

	for _, line := range lines {
		var tmpArr []int

		splitLine := strings.Split(line, "")

		for _, v := range splitLine {
			val, _ := strconv.Atoi(v)
			tmpArr = append(tmpArr, val)
		}
		formattedArray = append(formattedArray, tmpArr)
	}

	var currentMaxHeight int

	//iterate through input lines
	for rowIndex, row := range formattedArray {

		var lineScenicTreeArray []ScenicTree

		//iterate through line split values
		for itemIndex, itemValue := range row {
			currentMaxHeight = itemValue

			tree := ScenicTree{
				Value: itemValue,
				North: 0,
				East:  0,
				South: 0,
				West:  0,
			}

			//Check North
			for formattedArrayIndx := rowIndex - 1; formattedArrayIndx >= 0; formattedArrayIndx-- {
				if formattedArray[formattedArrayIndx][itemIndex] < currentMaxHeight {
					tree.North += 1
				}
				if formattedArray[formattedArrayIndx][itemIndex] >= currentMaxHeight {
					tree.North += 1
					break
				}
			}

			//Check East
			for eastIndex := (itemIndex + 1); eastIndex < len(row); eastIndex++ {
				if row[eastIndex] < currentMaxHeight {
					tree.East += 1
				}
				if row[eastIndex] >= currentMaxHeight {
					tree.East += 1
					break
				}
			}

			//Check South
			for formattedArrayIndx := rowIndex + 1; formattedArrayIndx < len(formattedArray); formattedArrayIndx++ {
				if formattedArray[formattedArrayIndx][itemIndex] < currentMaxHeight {
					tree.South += 1
				}
				if formattedArray[formattedArrayIndx][itemIndex] >= currentMaxHeight {
					tree.South += 1
					break
				}
			}

			//Check West
			for westIndex := (itemIndex - 1); westIndex >= 0; westIndex-- {
				if row[westIndex] < currentMaxHeight {
					tree.West += 1
				}
				if row[westIndex] >= currentMaxHeight {
					tree.West += 1
					break
				}
			}

			lineScenicTreeArray = append(lineScenicTreeArray, tree)
		}
		treeArr = append(treeArr, lineScenicTreeArray)
		currentMaxHeight = 0
	}

	var answerArr []int
	for _, row := range treeArr {
		for _, scenicTree := range row {
			answerArr = append(answerArr,
				(scenicTree.North *
					scenicTree.East *
					scenicTree.South *
					scenicTree.West))
		}
	}

	var total int
	for _, val := range answerArr {
		if val > total {
			total = val
		}
	}
	fmt.Printf("Part 2 Answer: %v", total)
}
