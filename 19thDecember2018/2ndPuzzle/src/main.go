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

func getDividers(number int) (dividers []int) {
	for i:= 1; i <= number; i ++ {
		if number%i == 0 {
			fmt.Println( "10551398%",i,"=",10551398%i)
			dividers = append(dividers, i)
		}
	}
	return
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	input := strings.Split(string(file), "\n")

	instructionPointerReg, err := strconv.Atoi(strings.Trim(input[0], "#ip "))
	checkError(err)

	instructions := input[1:]

	registers := []int{1, 0, 0, 0, 0, 0}
	instructionReg := []int{0, 0, 0, 0}
	dividersList := getDividers(10551398)
	dividersListIndex := 0
	fmt.Println(dividersList)
	sum := 0
	for _, value := range dividersList {
		sum += value
	}
	fmt.Println("The correct answer is:",sum) // this yields the correct answer since this is wat the assembly code does, but the loop should be fixed too actually
	for true {
		if registers[instructionPointerReg] >= len(instructions) || registers[instructionPointerReg] < 0 {
			break
		} else if registers[instructionPointerReg] == 3 { // wrote the routine like this to speed it somewhat up (not the fastest solution)

			for registers[4] < registers[1] {
				for registers[1] != registers[4]*registers[2] {
					if registers[2] > registers[1] {
						if dividersListIndex < len(dividersList) {
							registers[4] = dividersList[dividersListIndex]
							dividersListIndex += 1
						} else {
							registers[4] += 1
						}
						if registers[4] > registers[1] {
							registers[5] = 16
							break
						} else {
							registers[2] = 1
						}
					} else {
						registers[5] = 3
					}
					registers[2] += 1
				}
				if registers[1] == registers[4]*registers[2] {
					registers[0] += registers[4]
					registers[2] += 1
					if registers[2] > registers[1] {
						if dividersListIndex < len(dividersList) {
							registers[4] = dividersList[dividersListIndex]
							dividersListIndex += 1
						} else {
							registers[4] += 1
						}
						if registers[4] > registers[1] {
							registers[5] = 16
							break
						} else {
							registers[2] = 1
						}
					} else {
						registers[5] = 3
					}
				}
				fmt.Println(registers)
			}
		} else {
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
		fmt.Println(registers)
	}

	fmt.Println("The value in register 0 is:", registers[0])
}