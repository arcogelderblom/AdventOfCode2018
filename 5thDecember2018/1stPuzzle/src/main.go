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
	
	for i := 0; i < len(polymer)-1; i++ {
		if polymer[i] == (polymer[i+1]+32) || polymer[i] == (polymer[i+1]-32) {
			stringToDelete := string(polymer[i]) + string(polymer[i+1])
			polymer = strings.Replace(polymer, stringToDelete, "", -1)
			i = 0 // start over with checking since something is deleted
		}
	}
	fmt.Println("Amount of units left in polymer:", len(polymer))
}
