package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println(polymer)
}
