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
	xCoordinate int
	yCoordinate int
}

func abs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func calcManhattanDistance(start coordinate, end coordinate) int {
	return abs(start.xCoordinate - end.xCoordinate) + abs(start.yCoordinate - end.yCoordinate)
}

func sumDistancesLessthen10000(point coordinate, coordinates []coordinate) (bool) {
	totalDistance := 0
	for i := range coordinates {
		totalDistance += calcManhattanDistance(point, coordinates[i])
	}

	if totalDistance < 10000 {
		return true
	}
	return false
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	coordinateStrings := strings.Split(string(file), "\n")
	coordinates := []coordinate{}
	for i := 0; i < len(coordinateStrings); i++ {
		point := strings.Split(strings.Replace(coordinateStrings[i], " ", "", 1), ",")
		xCoordinate, err := strconv.Atoi(point[0])
		checkError(err)
		yCoordinate, err := strconv.Atoi(point[1])
		checkError(err)

		coordinates = append(coordinates, coordinate{xCoordinate, yCoordinate})
	}

	greatestXCoordinate := 0
	greatestYCoordinate := 0
	for i := 0; i < len(coordinates); i++ {
		if coordinates[i].xCoordinate > greatestXCoordinate {
			greatestXCoordinate = coordinates[i].xCoordinate
		}
		if coordinates[i].yCoordinate > greatestYCoordinate {
			greatestYCoordinate = coordinates[i].yCoordinate
		}
	}

	regionSize := 0
	for x := 0; x <= greatestXCoordinate; x++ {
		for y := 0; y <= greatestYCoordinate; y++ {
			if sumDistancesLessthen10000(coordinate{x,y}, coordinates) {
				regionSize += 1
			}
		}
	}

	fmt.Println("The size of the region containing all locations which have a total distance to all given coordinates of less than 10000:", regionSize)
}
