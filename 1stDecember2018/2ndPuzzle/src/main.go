package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	Function which calculates a result based on a int and a string
	which holds the number and the operator (like "+1")
 */
func calculate(startNumber int, calculation string) int {
	number, err := strconv.Atoi(calculation[1:])
	checkError(err)

	if calculation[0] == '-' {
		startNumber -= number
	} else if calculation[0] == '+' {
		startNumber += number
	}
	return startNumber
}

func contains(array []int, number int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == number {
			return true
		}
	}
	return false
}

func main() {
	startFreq := 0

	visitedNumbers := []int {}

	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	freqChanges := strings.Split(string(file), "\n")

	index := 0
	cur := startFreq
	for true {
		cur = calculate(cur, freqChanges[index])

		if contains(visitedNumbers, cur) {
			break
		} else {
			visitedNumbers = append(visitedNumbers, cur)
		}


		if index == (len(freqChanges)-1) {
			index = 0
		} else {
			index++
		}
	}

	fmt.Println("The first redundant frequency is:", cur)
}
