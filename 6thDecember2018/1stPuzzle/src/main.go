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

func getClosestPoint(point coordinate, coordinates []coordinate) (coordinate, bool) {
	distanceMap := make(map[coordinate]int)
	smallestDistance := 10000
	for i := range coordinates {
		distance := calcManhattanDistance(point, coordinates[i])
		distanceMap[coordinates[i]] = distance
		if distance < smallestDistance {
			smallestDistance = distance
		}
	}

	occurences := 0
	entrySmallestDistance := coordinate{}
	for entry := range distanceMap {
		if distanceMap[entry] == smallestDistance {
			entrySmallestDistance = entry
			occurences += 1
		}
	}

	if occurences == 1 {
		return entrySmallestDistance, true
	} else {
		return coordinate{}, false
	}
}

func contains(list []coordinate, element coordinate) bool {
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

	exclude := []coordinate{}
	pointsPerCoordinate := make(map[coordinate]int)
	for x := 0; x <= greatestXCoordinate; x++ {
		for y := 0; y <= greatestYCoordinate; y++ {
			closest, onePoint := getClosestPoint(coordinate{x,y}, coordinates)

			if onePoint {
				if x == 0 || x == greatestXCoordinate || y == 0 || y == greatestYCoordinate {
					if !contains(exclude, closest) {
						exclude = append(exclude, closest)
					}
				}
				pointsPerCoordinate[closest] += 1

			}
		}
	}

	highest := 0
	for point := range pointsPerCoordinate {
		if !contains(exclude, point) {
			if pointsPerCoordinate[point] > highest {
				highest = pointsPerCoordinate[point]
			}
		}
	}

	fmt.Println("The size of the largest area that is not infinite:", highest)
}
