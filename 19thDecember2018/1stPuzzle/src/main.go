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

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	input := strings.Split(string(file), "\n")

	instructionPointerReg, err := strconv.Atoi(strings.Trim(input[0], "#ip "))
	checkError(err)

	instructions := input[1:]

	registers := []int{0, 0, 0, 0, 0, 0}
	instructionReg := []int{0, 0, 0, 0}

	for true {
		if registers[instructionPointerReg] >= len(instructions) || registers[instructionPointerReg] < 0 {
			break
		}
		//fmt.Println(registers)
		curInstruction := strings.Split(instructions[registers[instructionPointerReg]], " ")
		instruction := curInstruction[0]
		instructionReg[1], _ = strconv.Atoi(curInstruction[1])
		instructionReg[2], _ = strconv.Atoi(curInstruction[2])
		instructionReg[3], _ = strconv.Atoi(curInstruction[3])
		switch instruction {
		case "addr":
			registers = addr(registers, instructionReg)
		case "addi":
			registers = addi(registers, instructionReg)
		case "mulr":
			registers = mulr(registers, instructionReg)
		case "muli":
			registers = muli(registers, instructionReg)
		case "banr":
			registers = banr(registers, instructionReg)
		case "bani":
			registers = bani(registers, instructionReg)
		case "borr":
			registers = borr(registers, instructionReg)
		case "bori":
			registers = bori(registers, instructionReg)
		case "setr":
			registers = setr(registers, instructionReg)
		case "seti":
			registers = seti(registers, instructionReg)
		case "gtir":
			registers = gtir(registers, instructionReg)
		case "gtri":
			registers = gtri(registers, instructionReg)
		case "gtrr":
			registers = gtrr(registers, instructionReg)
		case "eqir":
			registers = eqir(registers, instructionReg)
		case "eqri":
			registers = eqri(registers, instructionReg)
		case "eqrr":
			registers = eqrr(registers, instructionReg)
		}

		registers[instructionPointerReg] += 1
	}
	fmt.Println("The value in register 0 is:", registers[0])
}