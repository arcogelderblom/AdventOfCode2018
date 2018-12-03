package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	xCoordinate int
	yCoordinate int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getInfo(claim string) (int, point, string){
	id, err := strconv.Atoi(strings.Trim(strings.Fields(claim)[0], "#"))
	checkError(err)

	coordinates := strings.Split(strings.Trim(strings.Fields(claim)[2], ":"), ",")
	xCoordinate, err := strconv.Atoi(coordinates[0])
	checkError(err)

	yCoordinate, err := strconv.Atoi(coordinates[1])
	checkError(err)

	startPoint := point{xCoordinate,yCoordinate}

	size := strings.Fields(claim)[3]

	return id, startPoint, size
}

func calculatePositions(startPoint point, size string) []point {
	allPoints := []point{}

	axis := strings.Split(size, "x")
	xAxis, err := strconv.Atoi(axis[0])
	checkError(err)
	yAxis, err := strconv.Atoi(axis[1])
	checkError(err)

	for x := 0; x < xAxis; x++ {
		for y := 0; y < yAxis; y++ {
			allPoints = append(allPoints, point{startPoint.xCoordinate + x, startPoint.yCoordinate + y})
		}
	}

	return allPoints
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	claims := strings.Split(string(file), "\n")

	// Put everything in a map of all points belonging to a ID
	claimsMap := make(map[int][]point)
	for i := 0; i < len(claims); i++ {
		id, startPoint, size := getInfo(claims[i])
		claimsMap[id] = calculatePositions(startPoint, size)
	}

	// Check which points overlap
	overlapMap := make(map[point]int)
	for claim := range claimsMap {
		for i := 0; i < len(claimsMap[claim]); i++ {
			overlapMap[claimsMap[claim][i]] += 1
		}
	}

	overlapArray := []point{}
	for point := range overlapMap {
		if overlapMap[point] > 1 {
			overlapArray = append(overlapArray, point)
		}
	}

	fmt.Println("Total amount of square inches of overlap is:", len(overlapArray))
}
