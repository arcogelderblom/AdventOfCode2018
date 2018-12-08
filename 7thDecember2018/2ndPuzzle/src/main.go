package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(list []string, element string) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}

func getIndex(list []string, element string) int {
	for i := range list {
		if element == list[i] {
			return i
		}
	}
	return -1
}


func prerequisitesMet(prerequisites []string, prerequisitesDone []string) bool {
	for i := range prerequisites {
		if !contains(prerequisitesDone, prerequisites[i]) {
			return false
		}
	}
	return true
}

func calcTime(letter string) int {
	BASE_TIME := 60
	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return BASE_TIME + (getIndex(alphabet, letter) + 1)
}

func main() {
	LOCATION_NODE := 5
	LOCATION_BEFORE_NODE := 36

	AMOUNT_OF_WORKERS := 5

	instructionOrder := []string{}

	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	instructionStrings := strings.Split(string(file), "\n")

	allInstructions := []string{}
	instructionConnectedNodes := make(map[string][]string)
	instructionPrerequisitedNodes := make(map[string][]string)

	for i := 0; i < len(instructionStrings); i++ {
		if !contains(allInstructions, string(instructionStrings[i][LOCATION_NODE])) {
			allInstructions = append(allInstructions, string(instructionStrings[i][LOCATION_NODE]))
		} else if !contains(allInstructions, string(instructionStrings[i][LOCATION_BEFORE_NODE])) {
			allInstructions = append(allInstructions, string(instructionStrings[i][LOCATION_BEFORE_NODE]))
		}
	}

	for i := range allInstructions {
		for j := 0; j < len(instructionStrings); j++ {
			name := string(allInstructions[i])
			if name == string(instructionStrings[j][LOCATION_NODE]) {
				instructionConnectedNodes[name] = append(instructionConnectedNodes[name], string(instructionStrings[j][LOCATION_BEFORE_NODE]))
			} else if name == string(instructionStrings[j][LOCATION_BEFORE_NODE]) {
				instructionPrerequisitedNodes[name] = append(instructionPrerequisitedNodes[name], string(instructionStrings[j][LOCATION_NODE]))
			}
		}
	}

	amountOfInstructions := len(allInstructions)

	// find first instruction
	possibilities := allInstructions
	copy(possibilities, allInstructions)
	for node := range instructionConnectedNodes {
		for i := range instructionConnectedNodes[node] {
			if contains(possibilities, instructionConnectedNodes[node][i]) {
				index := getIndex(possibilities, instructionConnectedNodes[node][i])
				possibilities = append(possibilities[:index], possibilities[index+1:]...)
			}
		}
	}

	timeElapsed := 0
	workers := make(map[string]int)
	busyWith := []string{}

	for len(instructionOrder) != amountOfInstructions {
		sort.Strings(possibilities)

		index := 0
		current := ""
		if len(possibilities) > 0 {
			current = possibilities[index]

			for len(workers) < AMOUNT_OF_WORKERS && index <= (len(possibilities)-1) {
				index = 0
				current = possibilities[index]
				for !prerequisitesMet(instructionPrerequisitedNodes[current], instructionOrder) && index <= (len(possibilities)-1) {
					current = possibilities[index]
					if !prerequisitesMet(instructionPrerequisitedNodes[current], instructionOrder) {
						index++
					}
				}
				if index <= len(possibilities) - 1 {
					workers[current] = calcTime(current)

					busyWith = append(busyWith, current)

					index := getIndex(possibilities, current)
					possibilities = append(possibilities[:index], possibilities[index+1:]...)
				}
			}
		}
		fmt.Println(timeElapsed, workers)
		for i := range workers {
			workers[i] -= 1
		}
		for i := range workers {
			if workers[i] == 0 { // remove tasks that are done
				instructionOrder = append(instructionOrder, i)


				index := getIndex(busyWith, i)
				busyWith = append(busyWith[:index], busyWith[index+1:]...)

				delete(workers, i)
				for j := range instructionConnectedNodes[i] {
					if !contains(possibilities, instructionConnectedNodes[i][j]) && !contains(instructionOrder, instructionConnectedNodes[i][j]) && !contains(busyWith, instructionConnectedNodes[i][j]) {
						possibilities = append(possibilities, instructionConnectedNodes[i][j])
					}
				}

			}
		}

		timeElapsed += 1
	}

	fmt.Print("The time elapsed is: ", timeElapsed)
}
