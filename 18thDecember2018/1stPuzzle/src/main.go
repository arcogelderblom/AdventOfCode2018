package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getNewLand(land [][]string) [][]string {
	newLand := make([][]string, len(land))
	for i := range newLand {
		newLand[i] = make([]string, len(land[i]))
	}

	for y := range land {
		for x := range land[y] {
			newLand[y][x] = getNewValue(land[y][x], x, y, land)
		}
	}

	return newLand
}

func getNewValue(value string, xValue int, yValue int, land [][]string) string {
	surroundings := []string{}

	if yValue == 0 { // top row
		if xValue == 0 { // top left
			surroundings = append(surroundings, land[yValue+1][xValue])
			surroundings = append(surroundings, land[yValue+1][xValue+1])
			surroundings = append(surroundings, land[yValue][xValue+1])
		} else if xValue == len(land[0]) - 1 { // top right
			surroundings = append(surroundings, land[yValue+1][xValue-1])
			surroundings = append(surroundings, land[yValue][xValue-1])
			surroundings = append(surroundings, land[yValue+1][xValue])
		} else {
			surroundings = append(surroundings, land[yValue][xValue-1])
			surroundings = append(surroundings, land[yValue+1][xValue-1])
			surroundings = append(surroundings, land[yValue+1][xValue])
			surroundings = append(surroundings, land[yValue+1][xValue+1])
			surroundings = append(surroundings, land[yValue][xValue+1])
		}
	} else if yValue == len(land) - 1 { // bottom row
		if xValue == 0 { // bottom left
			surroundings = append(surroundings, land[yValue-1][xValue])
			surroundings = append(surroundings, land[yValue-1][xValue+1])
			surroundings = append(surroundings, land[yValue][xValue+1])
		} else if xValue == len(land[0]) - 1 { // bottom right
			surroundings = append(surroundings, land[yValue-1][xValue-1])
			surroundings = append(surroundings, land[yValue][xValue-1])
			surroundings = append(surroundings, land[yValue-1][xValue])
		} else {
			surroundings = append(surroundings, land[yValue][xValue-1])
			surroundings = append(surroundings, land[yValue-1][xValue-1])
			surroundings = append(surroundings, land[yValue-1][xValue])
			surroundings = append(surroundings, land[yValue-1][xValue+1])
			surroundings = append(surroundings, land[yValue][xValue+1])
		}
	} else if xValue == 0 {
		surroundings = append(surroundings, land[yValue-1][xValue])
		surroundings = append(surroundings, land[yValue-1][xValue+1])
		surroundings = append(surroundings, land[yValue][xValue+1])
		surroundings = append(surroundings, land[yValue+1][xValue+1])
		surroundings = append(surroundings, land[yValue+1][xValue])
	} else if xValue == len(land[0]) - 1 {
		surroundings = append(surroundings, land[yValue-1][xValue])
		surroundings = append(surroundings, land[yValue-1][xValue-1])
		surroundings = append(surroundings, land[yValue][xValue-1])
		surroundings = append(surroundings, land[yValue+1][xValue-1])
		surroundings = append(surroundings, land[yValue+1][xValue])
	} else { // everything in between
		surroundings = append(surroundings, land[yValue-1][xValue-1])
		surroundings = append(surroundings, land[yValue][xValue-1])
		surroundings = append(surroundings, land[yValue+1][xValue-1])
		surroundings = append(surroundings, land[yValue-1][xValue])
		surroundings = append(surroundings, land[yValue+1][xValue])
		surroundings = append(surroundings, land[yValue-1][xValue+1])
		surroundings = append(surroundings, land[yValue][xValue+1])
		surroundings = append(surroundings, land[yValue+1][xValue+1])
	}

	/*
		Determine new value
		Rules:
		An open acre will become filled with trees if three or more adjacent acres contained trees. Otherwise, nothing happens.
		An acre filled with trees will become a lumberyard if three or more adjacent acres were lumberyards. Otherwise, nothing happens.
		An acre containing a lumberyard will remain a lumberyard if it was adjacent to at least one other lumberyard and at least one acre containing trees. Otherwise, it becomes open.
	*/
	openAcre := 0
	woodedAcre := 0
	lumberyards := 0
	for i := range surroundings {
		if surroundings[i] == "." {
			openAcre += 1
		} else if surroundings[i] == "|" {
			woodedAcre += 1
		} else {
			lumberyards += 1
		}
	}

	if value == "." && woodedAcre >= 3 {
		value = "|"
	} else if value == "|" && lumberyards >= 3 {
		value = "#"
	} else if value == "#" && (lumberyards < 1 || woodedAcre < 1) {
		value = "."
	}

	return value
}


func getResourceValue(land [][]string) int {
	lumberyards := 0
	woodedAcres := 0

	for y := range land {
		for x := range land[y] {
			if land[y][x] == "|" {
				woodedAcres += 1
			} else if land[y][x] == "#" {
				lumberyards += 1
			}
		}
	}
	return lumberyards * woodedAcres
}

func main() {
	minutes := 10

	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	splittedLand := strings.Split(string(file), "\n")
	var land = make([][]string, len(splittedLand))
	for y := range splittedLand {
		for x := range splittedLand[y] {
			land[y] = append(land[y], string(splittedLand[y][x]))
		}
	}

	for minute := 0; minute < minutes; minute++ {
		for y := range land {
			for x := range land[y] {
				fmt.Print(land[y][x])
			}
			fmt.Print("\n")
		}
		fmt.Println()
		land = getNewLand(land)
	}

	fmt.Println("The total resource value after", minutes, "minutes is:", getResourceValue(land))
}