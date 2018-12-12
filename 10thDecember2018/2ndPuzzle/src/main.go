package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coordinate struct {
	left int
	right int
}

type point struct {
	position coordinate
	velocity coordinate
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getPoints(pointListStrings []string) map[int]point {
	pointMap := make(map[int]point)
	for i := range pointListStrings {
		star := getPointFromString(pointListStrings[i])
		pointMap[i] = star
	}

	return pointMap
}

func getPointFromString(pointString string) point {
	tmp := strings.Split(pointString, "> ")
	positionList := strings.Split(strings.Replace(strings.Trim(tmp[0], "position=< "), " ", "", -1), ",")
	velocityList := strings.Split(strings.Replace(strings.Trim(tmp[1], "velocity=< >"), " ", "", -1), ",")

	position := coordinate{}
	velocity := coordinate{}
	position.left, _ = strconv.Atoi(positionList[0])
	position.right, _ = strconv.Atoi(positionList[1])
	velocity.left, _ = strconv.Atoi(velocityList[0])
	velocity.right, _ = strconv.Atoi(velocityList[1])
	return point{position, velocity}
}

func getDimensions(points map[int]point) (int, int, int, int) {
	lowestLeft := 0
	highestRight := 0
	lowestUp := 0
	highestDown := 0
	for i := range points {
		if points[i].position.left < lowestLeft {
			lowestLeft = points[i].position.left
		}
		if points[i].position.left > highestRight {
			highestRight = points[i].position.left
		}
		if points[i].position.right < lowestUp {
			lowestUp = points[i].position.right
		}
		if points[i].position.right > highestDown {
			highestDown =points[i].position.right
		}
	}
	return lowestLeft, highestRight, lowestUp, highestDown
}

func printPoints(points map[int]point, startx int, endx int, starty int, endy int) string {
	stringMap := make(map[int]string)
	line := ""
	for x := startx; x <= endx; x++ {
		line += "."
	}
	line += "\n"

	for y := starty; y <= endy; y++ {
		stringMap[y] = line
	}

	for i := range points {
		out := []rune(stringMap[points[i].position.right])
		out[points[i].position.left + (startx * -1)] = '#'
		stringMap[points[i].position.right] = string(out)
	}

	returnString := ""
	for y := starty; y <= endy; y++ {
		returnString += stringMap[y]
	}
	return returnString
}

func calcNewPoints(points map[int]point) map[int]point {
	newPoints := make(map[int]point)
	for i := range points {
		velocity := points[i].velocity
		left := points[i].position.left + points[i].velocity.left
		right := points[i].position.right + points[i].velocity.right
		newPoints[i] = point{coordinate{left, right}, velocity}
	}
	return newPoints
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)

	pointsStrings := strings.Split(string(file), "\n")
	points := getPoints(pointsStrings)

	startx, endx, starty, endy := getDimensions(points)
	lastStartX, lastEndX,lastStartY,lastEndY := startx, endx, starty, endy

	i := 0
	oldPoints := make(map[int]point)
	for startx >= lastStartX && endx <= lastEndX && starty >= lastStartY && endy <= lastEndY {
		oldPoints = points
		points = calcNewPoints(points)
		lastStartX, lastEndX, lastStartY, lastEndY = startx, endx, starty, endy
		startx, endx, starty, endy = getDimensions(points)
		i++
	}

	fmt.Println("After", i - 1, "seconds, the message spells:")
	fmt.Println(printPoints(oldPoints, startx, endx, starty, endy))
}
