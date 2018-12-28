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

type coordinate struct {
	x int
	y int
}

func getClayPiecesFromString(clayPiece string) []coordinate {
	returnCoordinates := []coordinate{}
	if clayPiece[0] == 'x' {
		coordinates := strings.Split(clayPiece, ", ")
		x, err := strconv.Atoi(coordinates[0][2:])
		checkError(err)
		yPositions := strings.Split(strings.Trim(coordinates[1], "y="), "..")
		yStart, err := strconv.Atoi(yPositions[0])
		checkError(err)
		yEnd, err := strconv.Atoi(yPositions[1])
		checkError(err)
		for y := yStart; y <= yEnd; y++ {
			returnCoordinates = append(returnCoordinates, coordinate{x,y})
		}
	} else if clayPiece[0] == 'y' {
		coordinates := strings.Split(clayPiece, ", ")
		y, err := strconv.Atoi(coordinates[0][2:])
		checkError(err)
		xPositions := strings.Split(strings.Trim(coordinates[1], "x="), "..")
		xStart, err := strconv.Atoi(xPositions[0])
		checkError(err)
		xEnd, err := strconv.Atoi(xPositions[1])
		checkError(err)
		for x := xStart; x <= xEnd; x++ {
			returnCoordinates = append(returnCoordinates, coordinate{x,y})
		}
	}
	return returnCoordinates
}

func getMaxValues(coordinates []coordinate) (lowestY int, highestY int) {
	lowestY = coordinates[0].y
	highestY = coordinates[0].y
	for i := range coordinates {
		if coordinates[i].y < lowestY {
			lowestY = coordinates[i].y
		} else if coordinates[i].y > highestY {
			highestY = coordinates[i].y
		}
	}
	return
}

func contains(element coordinate, list []coordinate) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}

func fillWithWater(waterOrigin coordinate, clayCoordinates []coordinate, maxY int) (water []coordinate) {
	startingPoints := []coordinate{{waterOrigin.x,waterOrigin.y}}
	fallingDownwards := true
	for i := 0; i >= 0; i++ {
		if i == len(startingPoints) {
			break
		}

		x := startingPoints[i].x
		y := startingPoints[i].y

		for y < maxY {
			if !contains(coordinate{x, y}, clayCoordinates) && !contains(coordinate{x, y}, water) {
				water = append(water, coordinate{x, y})
				fallingDownwards = true
			} else if !contains(coordinate{x, y + 1}, clayCoordinates) && !contains(coordinate{x, y + 1}, water) {
				water = append(water, coordinate{x, y + 1})
				y += 1
				fallingDownwards = true
			} else if fallingDownwards &&  contains(coordinate{x, y + 1}, water) {
				break
			} else {
				fallingDownwards = false
				waterBuffer, newPoints := fillWithWaterHorizontally(coordinate{x, y}, clayCoordinates, water)
				startingPoints = append(startingPoints, newPoints...)
				water = append(water, waterBuffer...)

				if len(newPoints) == 0 {
					y -= 1
				} else {
					break
				}
			}
		}
	}
	return
}

func fillWithWaterHorizontally(waterStart coordinate, clayCoordinates []coordinate, water []coordinate) (addedWater []coordinate, newStartPoints []coordinate) {
	x := waterStart.x
	y := waterStart.y

	for !contains(coordinate{x-1, y}, clayCoordinates) && !contains(coordinate{x-1, y}, water) {
		addedWater = append(addedWater, coordinate{x-1, y})
		if !(contains(coordinate{x-1, y+1}, water) || contains(coordinate{x-1, y+1}, clayCoordinates)) {
			newStartPoints = append(newStartPoints, coordinate{x-1,y})
			break
		}
		x -= 1
	}

	x = waterStart.x
	y = waterStart.y
	for !contains(coordinate{x+1, y}, clayCoordinates) && !contains(coordinate{x+1, y}, water) {
		addedWater = append(addedWater, coordinate{x+1, y})
		if !(contains(coordinate{x+1, y+1}, water) || contains(coordinate{x+1, y+1}, clayCoordinates)) {
			newStartPoints = append(newStartPoints, coordinate{x+1,y})
			break
		}
		x += 1
	}

	return
}

func getAmountReached(water []coordinate, lowestY int, highestY int) (amount int) {
	for i := range water {
		if water[i].y >= lowestY && water[i].y <= highestY {
			amount += 1
		}
	}
	return
}

func printall(water []coordinate, clay []coordinate) {
	for y := 0; y < 2000; y++ {
		for x := 350; x < 650; x++ {
			if contains(coordinate{x,y}, water) {
				fmt.Print("~")
			} else if contains(coordinate{x,y}, clay) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	clayPiecesString := strings.Split(string(file), "\n")

	clayPieces := []coordinate{}
	for i := range clayPiecesString {
		clayPieces = append(clayPieces, getClayPiecesFromString(clayPiecesString[i])...)
	}

	lowestY, maxY := getMaxValues(clayPieces)

	water := fillWithWater(coordinate{500,0}, clayPieces, maxY)
	printall(water, clayPieces)
	fmt.Println("The amount of tiles the water can reach is:", getAmountReached(water, lowestY, maxY)) // minus one to get rid of the starting point
}