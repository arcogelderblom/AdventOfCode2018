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

func main() {
	twoTimes := 0 // How many times there was a occurence of 2 of the same
	threeTimes := 0 // How many times there was a occurence of 2 of the same
	checkedThreeTimes := false
	checkedTwoTimes := false

	file, err := ioutil.ReadFile("../input.txt")
	codes := strings.Split(string(file), "\n")
	checkError(err)

	for i := 0; i < len(codes); i++ {
		checkedThreeTimes = false
		checkedTwoTimes = false
		counter := make(map[string]int)

		for j := 0; j < len(codes[i]); j++ {
			counter[string(codes[i][j])] += 1
		}
		fmt.Println(counter)
		for _, value := range counter {
			if value == 3 && checkedThreeTimes == false {
				threeTimes += 1
				checkedThreeTimes = true
			} else if value == 2 && checkedTwoTimes == false {
				twoTimes += 1
				checkedTwoTimes = true
			}
		}
	}

	fmt.Println("Two times:", twoTimes, "Three times:", threeTimes)
	fmt.Println("Sum:", twoTimes, "*", threeTimes)
	result := twoTimes * threeTimes
	fmt.Println(result)
}
