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

func addr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] + registersCopy[instruction[2]]
	return registersCopy
}

func addi(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] + instruction[2]
	return registersCopy
}

func mulr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] * registers[instruction[2]]
	return registersCopy
}

func muli(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] * instruction[2]
	return registersCopy
}

func banr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] & registers[instruction[2]]
	return registersCopy
}

func bani(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] & instruction[2]
	return registersCopy
}

func borr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] | registers[instruction[2]]
	return registersCopy
}

func bori(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]] | instruction[2]
	return registersCopy
}

func setr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = registersCopy[instruction[1]]
	return registersCopy
}

func seti(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	registersCopy[instruction[3]] = instruction[1]
	return registersCopy
}

func gtir(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if instruction[1] > registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func gtri(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] > instruction[2] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func gtrr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] > registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func eqir(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if instruction[1] == registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func eqri(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] == instruction[2] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func eqrr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] == registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	}
	registersCopy[instruction[3]] = 0
	return registersCopy
}

func stringListToIntList(stringList []string) []int {
	var returnList = make([]int, len(stringList))
	for i := range stringList {
		returnList[i], _ = strconv.Atoi(stringList[i])
	}
	return returnList
}

func equals(int1 []int, int2 []int) bool {
	for i := range int1 {
		if int1[i] != int2[i] {
			return false
		}
	}

	return true
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	item := strings.Split(strings.Split(string(file), "\n\n\n")[0], "\n\n")

	count := 0
	functions := []interface{}{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	for i := range item {
		correctResult := 0
		splittedItem := strings.Split(item[i], "\n")
		before := stringListToIntList(strings.Split(strings.Trim(splittedItem[0], "Before: []"), ", "))
		instruction := stringListToIntList(strings.Split(splittedItem[1]," "))
		after := stringListToIntList(strings.Split(strings.Trim(splittedItem[2], "After: []"), ", "))

		for j := range functions {
			if equals(functions[j].(func([]int, []int)([]int))(before, instruction), after) {
				correctResult += 1
			}
		}

		if correctResult >= 3 {
			count += 1
		}
	}

	fmt.Println("The amount of samples which could have 3 or more instructions is:", count)
}