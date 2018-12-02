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
	inconsistency := 0
	correct1 := ""
	correct2 := ""

	file, err := ioutil.ReadFile("../input.txt")
	codes := strings.Split(string(file), "\n")
	checkError(err)

	for i := 0; i < len(codes); i++ {
		for j := i+1; j < len(codes); j++ {
			inconsistency = 0
			for k := 0; k < len(codes[i]); k++ {
				if codes[i][k] != codes[j][k] {
					inconsistency += 1
				}
			}
			if inconsistency == 1 { // if inconsistency of only 1 is found it is the correct pair of codes
				correct1 = codes[i]
				correct2 = codes[j]
				break
			}
		}
	}

	// find common letters
	common := ""
	for i := 0; i < len(correct1); i++ {
		if correct1[i] == correct2[i] {
			common += string(correct1[i])
		}
	}

	fmt.Println(correct1, correct2)
	fmt.Println(common)
}
