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

func GetFormattedData() ([][]string, []string) {
	inputResult := loadTextInput.LoadInput("input.txt")

	parse := strings.Split(inputResult, "\r\n")

	var indxArr []int

	//get indexes of char locations
	for i := 1; i <= 9; i++ {
		indx := strings.Index(parse[8], strconv.Itoa(i))
		indxArr = append(indxArr, indx)
	}

	//transpose text
	var formattedArr [][]string
	for _, v := range indxArr {
		var tmpArr []string
		for x := 7; x >= 0; x-- {
			parsedItem := string(parse[x][v])
			if parsedItem != " " {
				tmpArr = append(tmpArr, parsedItem)
			}
		}
		formattedArr = append(formattedArr, tmpArr)
		tmpArr = nil
	}

	return formattedArr, parse
}

func Part1() {
	formattedArray, parse := GetFormattedData()

	for _, v := range parse[10:] {
		splitThings := strings.Split(v, " ")
		amountToPop, _ := strconv.Atoi(splitThings[1])
		arrayIndexToPopFrom, _ := strconv.Atoi(splitThings[3])
		arrayIndexToPopFrom -= 1
		arrayIndexToPushTo, _ := strconv.Atoi(splitThings[5])
		arrayIndexToPushTo -= 1
		arrayPointerToPopFrom := &formattedArray[arrayIndexToPopFrom]
		toPopLen := len(*arrayPointerToPopFrom) - amountToPop
		items := (*arrayPointerToPopFrom)[toPopLen:len(*arrayPointerToPopFrom)]
		var itemsReversed []string
		for i := len(items) - 1; i >= 0; i-- {
			itemsReversed = append(itemsReversed, items[i])
		}
		*arrayPointerToPopFrom = (*arrayPointerToPopFrom)[:toPopLen]
		formattedArray[arrayIndexToPushTo] = append(formattedArray[arrayIndexToPushTo], itemsReversed...)
	}

	for _, v := range formattedArray {
		fmt.Print(v[len(v)-1])
	}
	fmt.Print("\n")
}

func Part2() {
	formattedArray, parse := GetFormattedData()

	for _, v := range parse[10:] {
		splitThings := strings.Split(v, " ")
		amountToPop, _ := strconv.Atoi(splitThings[1])
		arrayIndexToPopFrom, _ := strconv.Atoi(splitThings[3])
		arrayIndexToPopFrom -= 1
		arrayIndexToPushTo, _ := strconv.Atoi(splitThings[5])
		arrayIndexToPushTo -= 1
		arrayPointerToPopFrom := &formattedArray[arrayIndexToPopFrom]
		toPopLen := len(*arrayPointerToPopFrom) - amountToPop
		items := (*arrayPointerToPopFrom)[toPopLen:len(*arrayPointerToPopFrom)]
		*arrayPointerToPopFrom = (*arrayPointerToPopFrom)[:toPopLen]
		formattedArray[arrayIndexToPushTo] = append(formattedArray[arrayIndexToPushTo], items...)
	}

	for _, v := range formattedArray {
		fmt.Print(v[len(v)-1])
	}
}
