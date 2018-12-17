package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type note struct {
	left2 bool
	left1 bool
	plant bool
	right1 bool
	right2 bool
}

type pot struct {
	index int
	left2 bool
	left1 bool
	containsPlant bool
	right1 bool
	right2 bool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getBoolValue(char uint8) bool {
	if char == '#' {
		return true
	}
	return false
}

func getNoteMap(noteStrings []string) map[note]bool {
	noteMap := make(map[note]bool)
	for i := range noteStrings {
		key, value := getNoteBoolFromString(noteStrings[i])
		noteMap[key] = value
	}
	return noteMap
}

func getNoteBoolFromString(noteString string) (note, bool) {
	key, value := strings.Split(noteString, " => ")[0], strings.Split(noteString, " => ")[1]
	return note{getBoolValue(key[0]), getBoolValue(key[1]), getBoolValue(key[2]), getBoolValue(key[3]), getBoolValue(key[4])}, getBoolValue(value[0])
}

func getPotListFromString(potString string) []pot {
	var potList = make([]pot, len(potString)+4)

	for i := range potList {
		value := false

		switch i {
		case 0:
			potList[i] = pot{i-2, false, false, value, false, false}
		case 1:
			potList[i] = pot{i-2, false, potList[i-1].containsPlant, value, false, false}
			potList[i-1].right1 = value
		case len(potList) - 2:
			potList[i] = pot{i-2, potList[i-2].containsPlant, potList[i-1].containsPlant, value, false, false}
			potList[i-1].right1 = value
			potList[i-2].right2 = value
		case len(potList) - 1:
			potList[i] = pot{i-2, potList[i-2].containsPlant, potList[i-1].containsPlant, value, false, false}
			potList[i-1].right1 = value
			potList[i-2].right2 = value
		default:
			value = getBoolValue(potString[i-2])
			potList[i] = pot{i-2, potList[i-2].containsPlant, potList[i-1].containsPlant, value, false, false}
			potList[i-1].right1 = value
			potList[i-2].right2 = value
		}
	}

	return potList
}

func printPots(pots []pot) string {
	printString := ""
	for i := range pots {
		if pots[i].containsPlant {
			printString += "#"
		} else {
			printString += "."
		}
	}
	return printString
}

func normalizeGeneration(generation []pot) ([]pot, int, int) {
	startIndex := 0
	endIndex := 0
	// normalize the front
	if !generation[0].left2 && !generation[0].left1 && !generation[0].containsPlant && !generation[0].right1 && generation[0].right2 {
		startIndex = generation[0].index
	} else {
		if generation[0].containsPlant {
			secondPot := pot{generation[0].index-1, false, false, false, generation[0].containsPlant, generation[1].containsPlant}
			firstPot := pot{generation[0].index-2, false, false, false, false, generation[0].containsPlant}
			generation = append([]pot{firstPot, secondPot}, generation...)
			startIndex = generation[0].index
		} else if !generation[0].left2 && !generation[0].left1 && !generation[0].containsPlant && generation[0].right1 {
			firstPot := pot{generation[1].index-2, false, false, false, false, generation[1].containsPlant}
			generation = append([]pot{firstPot}, generation...)
			startIndex = generation[0].index
		} else {
			for i := range generation {
				if !generation[i].left2 && !generation[i].left1 && !generation[i].containsPlant && !generation[i].right1 && generation[i].right2 {
					startIndex = generation[i].index
					generation = generation[i:]
					break
				}
			}
		}
	}

	// normalize the back
	if generation[len(generation)-1].left2 && !generation[len(generation)-1].left1 && !generation[len(generation)-1].containsPlant && !generation[len(generation)-1].right1 && !generation[len(generation)-1].right2 {
		endIndex = generation[len(generation)-1].index
	} else {
		if generation[len(generation)-1].containsPlant {
			secondToLastPot := pot{generation[len(generation)-1].index+1, generation[len(generation)-2].containsPlant, generation[len(generation)-1].containsPlant, false, false, false}
			lastPot := pot{generation[len(generation)-1].index+2, generation[len(generation)-1].containsPlant, false, false, false, false}
			generation = append(generation, []pot{secondToLastPot, lastPot}...)
			endIndex = generation[len(generation)-1].index
		} else if generation[len(generation)-1].left1 && !generation[len(generation)-1].containsPlant && !generation[len(generation)-1].right1 && !generation[len(generation)-1].right2 {
			lastPot := pot{generation[len(generation)-1].index+1, generation[len(generation)-2].containsPlant, generation[len(generation)-1].containsPlant, false, false, false}
			generation = append(generation, lastPot)
			endIndex = generation[len(generation)-1].index
		} else {
			for i := len(generation) - 1; i > 0; i-- {
				if !generation[i].left2 && !generation[i].left1 && !generation[i].containsPlant && !generation[i].right1 && generation[i].right2 {
					startIndex = generation[i].index
					generation = generation[i:]
				}
			}
			for i := range generation {
				if generation[i].left2 && !generation[i].left1 && !generation[i].containsPlant && !generation[i].right1 && !generation[i].right2 {
					startIndex = generation[i].index
					generation = generation[:i]
				}
			}
		}
	}

	return generation, startIndex, endIndex
}

func evolve(oldGeneration []pot, startIndex int, endIndex int, noteMap map[note]bool) ([]pot, int, int) {
	newGeneration := []pot{}
	newStartIndex := startIndex
	newEndIndex := endIndex

	// calculate new generation
	for i := startIndex; i <= endIndex; i++ {
		index := i - startIndex
		oldPot := oldGeneration[index]
		noteForNoteMap := note{oldPot.left2, oldPot.left1, oldPot.containsPlant, oldPot.right1, oldPot.right2}
		switch index {
		case 0:
			newGeneration = append(newGeneration, pot{i, false, false, noteMap[noteForNoteMap], false, false})
		case 1:
			newGeneration = append(newGeneration, pot{i, false, newGeneration[index-1].containsPlant, noteMap[noteForNoteMap], false, false})
			newGeneration[index-1].right1 = noteMap[noteForNoteMap]
		default:
			newGeneration = append(newGeneration, pot{i, newGeneration[index-2].containsPlant, newGeneration[index-1].containsPlant, noteMap[noteForNoteMap], false, false})
			newGeneration[index-2].right2 = noteMap[noteForNoteMap]
			newGeneration[index-1].right1 = noteMap[noteForNoteMap]
		}
	}

	newGeneration, newStartIndex, newEndIndex = normalizeGeneration(newGeneration)

	return newGeneration, newStartIndex, newEndIndex
}

func findRecurringSequence(generation []pot, noteMap map[note]bool) int {
	startIndex := -2
	endIndex := len(generation) - 3

	generationNum := 0
	oldGeneration := []pot{}
	for printPots(oldGeneration) != printPots(generation) {
		oldGeneration = generation
		generation, startIndex, endIndex = evolve(generation, startIndex, endIndex, noteMap)
		generationNum += 1
	}
	fmt.Println(generationNum-1, "is the first recurring sequence.")
	fmt.Println(printPots(oldGeneration))
	return generationNum - 1
}

func getGeneration(generation []pot, noteMap map[note]bool, generationNum int) []pot {
	startIndex := -2
	endIndex := len(generation) - 3

	recurringNum := findRecurringSequence(generation, noteMap)
	if generationNum < recurringNum {
		for i := 1; i <= generationNum; i ++ {
			generation, startIndex, endIndex = evolve(generation, startIndex, endIndex, noteMap)
		}
	} else {
		for i := 1; i <= recurringNum; i ++ {
			generation, startIndex, endIndex = evolve(generation, startIndex, endIndex, noteMap)
		}
		for i := range generation {
			generation[i].index += generationNum - recurringNum
		}
	}
	fmt.Println("Generation:", generationNum)
	fmt.Println(printPots(generation))
	fmt.Println()
	return generation
}

func getSum(plants []pot) int {
	sum := 0
	for i := range plants {
		if plants[i].containsPlant {
			sum += plants[i].index
		}
	}
	return sum
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)

	initialStateString := strings.Split(string(file), "\n")[0][15:]
	noteStrings := strings.Split(string(file), "\n")[2:]

	fmt.Println("Total sum of the pots' numbers which contain plants:", getSum(getGeneration(getPotListFromString(initialStateString), getNoteMap(noteStrings), 50000000000)))
}
