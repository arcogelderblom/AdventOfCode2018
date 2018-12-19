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
	registersCopy[instruction[3]] = registersCopy[instruction[1]] * registersCopy[instruction[2]]
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
	registersCopy[instruction[3]] = registersCopy[instruction[1]] & registersCopy[instruction[2]]
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
	registersCopy[instruction[3]] = registersCopy[instruction[1]] | registersCopy[instruction[2]]
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
	} else {
		registersCopy[instruction[3]] = 0
	}
	return registersCopy
}

func gtri(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] > instruction[2] {
		registersCopy[instruction[3]] = 1
	} else {
		registersCopy[instruction[3]] = 0
	}
	return registersCopy
}

func gtrr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] > registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	} else {
		registersCopy[instruction[3]] = 0
	}
	return registersCopy
}

func eqir(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if instruction[1] == registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	} else {
		registersCopy[instruction[3]] = 0
	}
	return registersCopy
}

func eqri(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] == instruction[2] {
		registersCopy[instruction[3]] = 1
	} else {
		registersCopy[instruction[3]] = 0
	}
	return registersCopy
}

func eqrr(registers []int, instruction []int) []int {
	var registersCopy = make([]int, len(registers))
	copy(registersCopy, registers)
	if registersCopy[instruction[1]] == registersCopy[instruction[2]] {
		registersCopy[instruction[3]] = 1
	} else {
		registersCopy[instruction[3]] = 0
	}
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

func contains(list []string, element string) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	item := strings.Split(strings.Split(string(file), "\n\n\n")[0], "\n\n")

	functions := []interface{}{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	options := []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	optionNumName := make(map[int]string)
	for len(optionNumName) < 16 {
		for i := range item {
			tmpPossibilities := []string{}
			correctResult := 0
			splittedItem := strings.Split(item[i], "\n")
			before := stringListToIntList(strings.Split(strings.Trim(splittedItem[0], "Before: []"), ", "))
			instruction := stringListToIntList(strings.Split(splittedItem[1], " "))
			after := stringListToIntList(strings.Split(strings.Trim(splittedItem[2], "After: []"), ", "))

			if optionNumName[instruction[0]] == "" {
				for j := range functions {
					if equals(functions[j].(func([]int, []int) ([]int))(before, instruction), after) {
						if options[j] != "" {
							correctResult += 1
							tmpPossibilities = append(tmpPossibilities, options[j])
						}
					}
				}
			}

			if correctResult == 1 {
				optionNumName[instruction[0]] = tmpPossibilities[0]
				for j := range options {
					if options[j] == tmpPossibilities[0] {
						options[j] = ""
					}
				}
			}
		}
	}

	options = []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"}
	fmt.Println(optionNumName)
	register := []int{0,0,0,0}
	opCodes := strings.Split(strings.Split(string(file), "\n\n\n\n")[1], "\n")
	for i := range opCodes {
		opCode := stringListToIntList(strings.Split(opCodes[i], " "))
		for j := range options {
			if options[j] == optionNumName[opCode[0]] {
				register = functions[j].(func([]int, []int) ([]int))(register, opCode)
			}
		}
	}

	fmt.Println("The value in register 0 is:", register[0])
}