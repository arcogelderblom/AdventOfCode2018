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

func main() {
	LOCATION_NODE := 5
	LOCATION_BEFORE_NODE := 36
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

	for len(instructionOrder) != len(allInstructions) {
		sort.Strings(possibilities)

		start := 0
		current := possibilities[start]
		for !prerequisitesMet(instructionPrerequisitedNodes[current], instructionOrder) {
			start++
			current = possibilities[start]
		}

		for i := range instructionConnectedNodes[current] {
			if !contains(possibilities, instructionConnectedNodes[current][i]) && !contains(instructionOrder, instructionConnectedNodes[current][i]) {
				possibilities = append(possibilities, instructionConnectedNodes[current][i])
			}
		}

		instructionOrder = append(instructionOrder, current)

		index := getIndex(possibilities, current)
		possibilities = append(possibilities[:index], possibilities[index+1:]...)
	}

	fmt.Print("The order of instructions is: ")
	for i := range instructionOrder {
		fmt.Print(instructionOrder[i])
	}
}
