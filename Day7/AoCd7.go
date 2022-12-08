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

type Directory struct {
	Name              string
	ParentDirectory   *Directory
	Files             []int
	NestedDirectories []*Directory
}

func RecursionThroughFolders(folderStructure map[string]*Directory, pwd *Directory) int {
	totalToReturn := 0
	totalToReturn += Sum(pwd.Files)
	for _, dir := range pwd.NestedDirectories {
		totalToReturn += RecursionThroughFolders(folderStructure, folderStructure[dir.Name])
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

func Part1() {
	inputResult := loadTextInput.LoadInput("input.txt")

	parse := strings.Split(inputResult, "\r\n")

	currentWorkingDirectory := &Directory{
		Name:              "/",
		Files:             nil,
		NestedDirectories: nil,
		ParentDirectory:   nil}

	mapOfFolders := make(map[string]*Directory)

	mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory

	for i, v := range parse {
		if strings.HasPrefix(v, "$ ls") {
			continue
		} else if strings.HasPrefix(v, "dir") {
			if _, exists := mapOfFolders[fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:])]; exists {
				currentWorkingDirectory.NestedDirectories =
					append(currentWorkingDirectory.NestedDirectories,
						mapOfFolders[fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:])])
			} else {
				newDirectory := &Directory{
					Name:              fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:]),
					Files:             nil,
					NestedDirectories: nil,
					ParentDirectory:   currentWorkingDirectory}
				currentWorkingDirectory.NestedDirectories =
					append(currentWorkingDirectory.NestedDirectories, newDirectory)
				mapOfFolders[newDirectory.Name] = newDirectory
			}
		} else if strings.HasPrefix(v, "$ cd") {
			mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory
			if strings.HasSuffix(v, "..") {
				currentWorkingDirectory = mapOfFolders[currentWorkingDirectory.ParentDirectory.Name]
			} else if strings.HasSuffix(v, "/") {
				currentWorkingDirectory = mapOfFolders["/"]
			} else {
				p := fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[5:])
				currentWorkingDirectory = mapOfFolders[p]
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

func Part2() {
	inputResult := loadTextInput.LoadInput("input.txt")

	parse := strings.Split(inputResult, "\r\n")

	currentWorkingDirectory := &Directory{
		Name:              "/",
		Files:             nil,
		NestedDirectories: nil,
		ParentDirectory:   nil}

	mapOfFolders := make(map[string]*Directory)

	mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory

	for i, v := range parse {
		if strings.HasPrefix(v, "$ ls") {
			continue
		} else if strings.HasPrefix(v, "dir") {
			if _, exists := mapOfFolders[fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:])]; exists {
				currentWorkingDirectory.NestedDirectories =
					append(currentWorkingDirectory.NestedDirectories,
						mapOfFolders[fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:])])
			} else {
				newDirectory := &Directory{
					Name:              fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[4:]),
					Files:             nil,
					NestedDirectories: nil,
					ParentDirectory:   currentWorkingDirectory}
				currentWorkingDirectory.NestedDirectories =
					append(currentWorkingDirectory.NestedDirectories, newDirectory)
				mapOfFolders[newDirectory.Name] = newDirectory
			}
		} else if strings.HasPrefix(v, "$ cd") {
			mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory
			if strings.HasSuffix(v, "..") {
				currentWorkingDirectory = mapOfFolders[currentWorkingDirectory.ParentDirectory.Name]
			} else if strings.HasSuffix(v, "/") {
				currentWorkingDirectory = mapOfFolders["/"]
			} else {
				p := fmt.Sprintf("%v/%v", currentWorkingDirectory.Name, v[5:])
				currentWorkingDirectory = mapOfFolders[p]
			}
		} else {
			x, _ := strconv.Atoi(strings.Split(v, " ")[0])
			currentWorkingDirectory.Files = append(currentWorkingDirectory.Files, x)
			if i == len(parse)-1 {
				mapOfFolders[currentWorkingDirectory.Name] = currentWorkingDirectory
			}
		}
	}

	var fileSizeTotal int
	var spaceToFree int
	root := mapOfFolders["/"]
	fileSizeTotal += RecursionThroughRoot(mapOfFolders, root)
	fmt.Printf("Total file size: %v\n", fileSizeTotal)
	spaceToFree = 30000000 - (70000000 - fileSizeTotal)
	fmt.Printf("Total space needed to free: %v\n", spaceToFree)

	folderSizes := make(map[string]int)
	for _, folder := range mapOfFolders {
		fileSizeTotal := 0
		fileSizeTotal += RecursionThroughFolders(mapOfFolders, folder)
		folderSizes[folder.Name] = fileSizeTotal
	}
	minSizeOfFolderToDelete := fileSizeTotal
	for _, val := range folderSizes {
		if val >= spaceToFree && val < minSizeOfFolderToDelete {
			minSizeOfFolderToDelete = val
		}
	}
	fmt.Printf("Minimum size of folder to delete (part 2): %v\n", minSizeOfFolderToDelete)
}

func RecursionThroughRoot(mapOfFolders map[string]*Directory, directory *Directory) int {
	totalToReturn := 0
	totalToReturn += Sum(directory.Files)
	for _, dir := range directory.NestedDirectories {
		totalToReturn += RecursionThroughRoot(mapOfFolders, mapOfFolders[dir.Name])
	}
	return totalToReturn
}
