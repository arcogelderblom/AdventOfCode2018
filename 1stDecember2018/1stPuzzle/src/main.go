package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	//"strings"
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

func main() {
	startFreq := 0

	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	freqChanges := strings.Split(string(file), "\n")

	for i := 0; i < len(freqChanges); i++ {
		startFreq = calculate(startFreq, freqChanges[i])

	}

	fmt.Println("This results in the frequency:", startFreq)
}
