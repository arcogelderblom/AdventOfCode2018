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

func insert(list []int, insertion int, index int) []int {
	var slice1 = make([]int, len(list[:index]))
	var slice2 = make([]int, len(list[index:]))
	copy(slice1, list[:index])
	copy(slice2, list[index:])
	return append(append(slice1, insertion), slice2...)
}

func addMarble(currentList []int, currentIndex int, marble int) ([]int, int) {
	if marble == 1 {
		currentList = append(currentList, marble)
		currentIndex += 1
	} else {
		if currentIndex == len(currentList) - 1 {
			currentIndex = 1
			currentList = insert(currentList, marble, currentIndex)
			currentIndex = 1
		} else if currentIndex + 1 ==  len(currentList) - 1 {
			currentList = append(currentList, marble)
			currentIndex += 2
		} else {
			currentIndex += 2
			currentList = insert(currentList, marble, currentIndex)
		}
	}
	return currentList, currentIndex
}

func countScore(gameList []int, playerList []int, playersTurn int, marble int, currentIndex int) ([]int, []int, int) {
	currentIndex -= 7
	if currentIndex < 0 {
		currentIndex = len(gameList) + currentIndex
	}
	playerList[playersTurn-1] += marble + gameList[currentIndex]
	gameList = append(gameList[:currentIndex], gameList[currentIndex+1:]...)
	return gameList, playerList, currentIndex
}

func playGame(amountOfPlayers int, lastMarbleVal int) []int {
	var playerList = make([]int, amountOfPlayers) // list of players where index is the playernumber and the value is the score
	gameField := []int{0}

	playerTurn := 1
	index := 0
	for i := 1; i <= lastMarbleVal; i++ {
		if i % 23 == 0 {
			gameField, playerList, index = countScore(gameField, playerList, playerTurn, i, index)
		} else {
			gameField, index = addMarble(gameField, index, i)
		}
		playerTurn += 1
		if playerTurn > amountOfPlayers {
			playerTurn = 1
		}
	}
	return playerList
}

func getMaxScore(scores []int) (int, int){
	maxScore := 0
	bestPlayer := 0
	for i := range scores {
		if scores[i] > maxScore {
			maxScore = scores[i]
			bestPlayer = i + 1
		}
	}
	return maxScore, bestPlayer
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	amountOfPlayers, err := strconv.Atoi(strings.Split(string(file), " ")[0])
	checkError(err)
	lastMarble, err := strconv.Atoi(strings.Split(string(file), " ")[6])
	checkError(err)

	lastMarble *= 100 // marble 100 times larger
	players := playGame(amountOfPlayers, lastMarble)
	maxScore, bestPlayer := getMaxScore(players)

	fmt.Println("The best player is", bestPlayer, "with a score of", maxScore)
}
