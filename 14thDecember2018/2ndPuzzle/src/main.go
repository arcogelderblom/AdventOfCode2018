package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func calcPowerLevelGrid(startPoint coordinate, serialNumber int, gridSizeX int, gridSizeY int) int {
	totalPowerLevel := 0
	for x:= startPoint.x; x < startPoint.x + gridSizeX; x++ {
		for y:= startPoint.y; y < startPoint.y + gridSizeY; y++ {
			point := coordinate{x,y }
			totalPowerLevel += calcPowerLevel(point, serialNumber)
		}
	}
	return totalPowerLevel
}

func calcPowerLevel(point coordinate, serialNumber int) int {
	rackId := point.x + 10
	powerLevel := rackId * point.y
	powerLevel += serialNumber
	powerLevel *= rackId
	if powerLevel / 100 > 1 {
		powerLevel = (powerLevel / 100) % 10
	} else {
		powerLevel = 0
	}
	powerLevel -= 5
	return powerLevel
}

func main() {
	input := 5177 // serial number

	largestValuePoint := coordinate{}
	largestValueGridSize := 0
	largestValue := 0
	for gridSize := 1; gridSize < 300; gridSize++ {
		fmt.Println(gridSize)
		for x := 1; x <= 300-gridSize-1; x++ {
			for y := 1; y <= 300-gridSize-1; y++ {
				point := coordinate{x, y}
				curValue := calcPowerLevelGrid(point, input, gridSize, gridSize)
				if curValue > largestValue {
					largestValueGridSize = gridSize
					largestValuePoint = point
					largestValue = curValue
				}
			}
		}
	}

	fmt.Println("Point, grid size with the largest value is:", largestValuePoint, largestValueGridSize)
}
