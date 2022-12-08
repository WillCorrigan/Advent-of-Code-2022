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

type Folder struct {
	Name            string
	ParentDirectory string
	Files           []int
	NestedFolders   []string
}

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	parse := strings.Split(inputResult, "\r\n")

	currentWorkingDirectory := Folder{
		Name:            "/",
		Files:           nil,
		NestedFolders:   nil,
		ParentDirectory: ""}

	mapOfFolders := make(map[string]Folder)

	for i, v := range parse {
		if v[0:4] == "$ ls" {
			continue
		} else if v[0:3] == "dir" {
			currentWorkingDirectory.NestedFolders = append(currentWorkingDirectory.NestedFolders, v[4:])
		} else if v[0:4] == "$ cd" {
			mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory
			if v[5:] == ".." {
				currentWorkingDirectory = mapOfFolders[currentWorkingDirectory.ParentDirectory]
			} else {
				currentWorkingDirectory =
					Folder{
						Name:            currentWorkingDirectory.Name + v[5:],
						Files:           nil,
						NestedFolders:   nil,
						ParentDirectory: currentWorkingDirectory.Name}
			}
		} else {
			x, _ := strconv.Atoi(strings.Split(v, " ")[0])
			currentWorkingDirectory.Files = append(currentWorkingDirectory.Files, x)
			if i == len(parse)-1 {
				mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory
			}
		}
	}

	var total int
	for _, folder := range mapOfFolders {
		fileSizeTotal := 0
		fileSizeTotal += RecursionThroughFolders(mapOfFolders, folder)
		if fileSizeTotal < 100000 {
			total += fileSizeTotal
		}
	}
	fmt.Println(total)
}

func RecursionThroughFolders(folderStructure map[string]Folder, pwd Folder) int {
	totalToReturn := 0
	totalToReturn += Sum(pwd.Files)
	for _, dir := range pwd.NestedFolders {
		totalToReturn += RecursionThroughFolders(folderStructure, folderStructure[dir])
	}
	return totalToReturn
}

func Sum(arr []int) int {
	total := 0
	if len(arr) == 0 {
		return 0
	}
	for _, v := range arr {
		total += v
	}
	return total
}

// func Part2() {

// }
