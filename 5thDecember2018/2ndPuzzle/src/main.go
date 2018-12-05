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
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	polymer := string(file)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	results := make(map[string]int)

	for i := 0; i < len(alphabet); i++{
		result := strings.Replace(polymer, string(alphabet[i]), "", -1)
		result = strings.Replace(result, string(alphabet[i] - 32), "", -1)
		for i := 0; i < len(result)-1; i++ {
			if result[i] == (result[i+1]+32) || result[i] == (result[i+1]-32) {
				stringToDelete := string(result[i]) + string(result[i+1])
				result = strings.Replace(result, stringToDelete, "", -1)
				i = 0 // start over with checking since something is deleted
			}
		}
		results[string(alphabet[i])] = len(result)
	}

	lowestCount := len(polymer)
	for letter := range results {
		if results[letter] < lowestCount {
			lowestCount = results[letter]
		}
	}

	fmt.Println("Amount of units left in polymer:", lowestCount)
}
