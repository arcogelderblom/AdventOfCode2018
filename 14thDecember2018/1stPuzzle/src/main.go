package main

import (
	"fmt"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func updateElfIndex(recipes []string, currentIndex int) int {
	movePlaces, err := strconv.Atoi(recipes[currentIndex])
	checkError(err)
	movePlaces += 1
	index := currentIndex
	for i := movePlaces; i > 0; i-- {
		index += 1
		if index == len(recipes) {
			index = 0
		}
	}
	return index
}

func addNumberToStringList(number int, list []string) []string {
	numberAsString := strconv.Itoa(number)
	for i := range numberAsString {
		list = append(list, string(numberAsString[i]))
	}
	return list
}

func getScoresNextTen(recipes []string, tries int) string {
	elf1Index := 0
	elf2Index := 1

	for len(recipes) < tries + 10 {
		elf1Value, err := strconv.Atoi(recipes[elf1Index])
		checkError(err)
		elf2Value, err := strconv.Atoi(recipes[elf2Index])
		checkError(err)
		value := elf1Value + elf2Value
		recipes = addNumberToStringList(value, recipes)
		elf1Index = updateElfIndex(recipes, elf1Index)
		elf2Index = updateElfIndex(recipes, elf2Index)
	}

	returnString := ""
	for i := range recipes[tries:] {
		returnString += recipes[tries+i]
	}

	return returnString
}

func main() {
	recipes := []string{"3", "7"}
	input := 702831
	fmt.Println("Scores of the next ten recipes is:", getScoresNextTen(recipes, input))
}