package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(list []time.Time, item time.Time) bool {
	for i := range list {
		if item == list[i] {

			return true
		}
	}
	return false
}

func sortByTime(list []time.Time) []time.Time {
	sortedList := []time.Time{}

	for len(sortedList) != len(list) { // redo the sort list until everything is sorted
		for i := 0; i < len(list); i++ {
			lowest := list[i]
			for j := range list {
				if !contains(sortedList, list[j]) && list[j].Before(lowest) {
					lowest = list[j]
				}
			}
			if !contains(sortedList, lowest) {
				sortedList = append(sortedList, lowest)
			}

		}
	}

	return sortedList
}

func getGuardID(sentence string) int {
	separatedString := strings.Split(sentence, " ")
	guardID, err := strconv.Atoi(strings.Trim(separatedString[1], "#"))
	checkError(err)
	return guardID
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	logs := strings.Split(string(file), "\n")

	timeLogMap := make(map[time.Time]string)
	timeList := []time.Time{}
	for i := 0; i < len(logs); i++ {
		splittedLine := strings.Split(logs[i], "]")
		time, err := time.Parse("2006-01-02 15:04", strings.Trim(splittedLine[0], "[]"))
		checkError(err)
		timeList = append(timeList, time)
		timeLogMap[time] = strings.Trim(splittedLine[1], " ")
	}

	timeList = sortByTime(timeList)

	sleepMinutes := make(map[int]map[int]int)
	guardsSleeptime := make(map[int]int) // map with total amount of sleep per guard
	startMinute := 0
	endMinute := 0
	guardId := 0

	for i := 0; i < len(timeList); i++ {
		if timeLogMap[timeList[i]] != "falls asleep" && timeLogMap[timeList[i]] != "wakes up" {
			guardId = getGuardID(timeLogMap[timeList[i]])
		} else if timeLogMap[timeList[i]] == "falls asleep" {
			_, minute, _ := timeList[i].Clock() // omitted seconds and hours since they are irrelevant
			startMinute = minute
		} else if timeLogMap[timeList[i]] == "wakes up" {
			_, minute, _ := timeList[i].Clock() // omitted seconds and hours since they are irrelevant
			endMinute = minute
			guardsSleeptime[guardId] += endMinute - startMinute

			for j := startMinute; j < endMinute; j++ {
				if sleepMinutes[guardId] == nil {
					sleepMinutes[guardId] = map[int]int{}
				}
				sleepMinutes[guardId][j] += 1
			}
		}
	}

	idOfGuardHighestSleeptime := 0
	for i := range guardsSleeptime {
		highestSleeptime := guardsSleeptime[i]
		idOfGuardHighestSleeptime = i
		for j := range guardsSleeptime {
			if guardsSleeptime[j] > highestSleeptime {
				highestSleeptime = guardsSleeptime[j]
				idOfGuardHighestSleeptime = j
			}
		}
	}

	fmt.Println("Guard ID of the guard that is the most asleep:", idOfGuardHighestSleeptime)

	minuteMostAsleep := 0
	mostAsleepTime := 0
	for minute := range sleepMinutes[idOfGuardHighestSleeptime] {
		if sleepMinutes[idOfGuardHighestSleeptime][minute] > mostAsleepTime {
			minuteMostAsleep = minute
			mostAsleepTime = sleepMinutes[idOfGuardHighestSleeptime][minute]
		}
	}

	fmt.Println("The minute the guard is the most asleep:", minuteMostAsleep)

	fmt.Println("The ID of the guard multiplied by the minute equals:", (idOfGuardHighestSleeptime * minuteMostAsleep))
}
