package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func calcPowerLevelGrid(startPoint coordinate, serialNumber int, gridSizeX int, gridSizeY int) int {
	totalPowerLevel := 0
	for x:= startPoint.x; x < startPoint.x +  gridSizeX; x++ {
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
	largestValue := 0
	for x:= 1; x <= 300-2; x++ {
		for y:= 1; y <= 300-2; y++ {
			point := coordinate{x,y }
			curValue := calcPowerLevelGrid(point, input, 3, 3)
			if curValue > largestValue {
				largestValuePoint = point
				largestValue = curValue
			}
		}
	}

	fmt.Println("Point with the largest value is:", largestValuePoint)
}
