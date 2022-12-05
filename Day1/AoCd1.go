package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
		Part1()
		Part2()
	}

func Part1() {
	inputResult := LoadInput("input.txt")

	temp := strings.Split(inputResult, "\n")
	max := 0
	curr := 0

	for _, v := range temp {

		if len(v) != 0 {
			intVar, _ := strconv.Atoi(v)
			curr += intVar
		}
	
		if len(v) == 0 {
			if max <= curr {
				max = curr
			}
			curr = 0
		}
	}
	fmt.Println(max)
}

func Part2() {
	inputResult := LoadInput("input.txt")

	temp := strings.Split(inputResult, "\n")
	var listOfMax []int
	curr := 0

	for _, v := range temp {

		if len(v) != 0 {
			intVar, _ := strconv.Atoi(v)
			curr += intVar
		}
	
		if len(v) == 0 {
			listOfMax = append(listOfMax, curr)
			curr = 0
		}
	}

	sort.Slice(listOfMax, func(i, j int) bool {
    return listOfMax[i] > listOfMax[j]
})

result := 0

for i := 0; i < len(listOfMax[0:3]); i++ {
	result += listOfMax[i]
}

	fmt.Println(result)
}

func LoadInput(path string) string {
    input, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    return string(input)
}