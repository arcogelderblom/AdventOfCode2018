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

type node struct {
	children int
	metadata int
	data int
}

func convertToInt(list []string) []int {
	returnArray := []int{}

	for index := range list {
		number, err := strconv.Atoi(list[index])
		checkError(err)

		returnArray = append(returnArray, number)
	}
	return returnArray
}
func sum(input []int) int {
	sum := 0
	for index := range input {
		sum += input[index]
	}
	return sum
}

func createNodes(input []int) []node {
	indexChildren := 0
	indexMetadata := 1
	headerLen := 2

	nodeList := []node{}

	for len(input) > 0 {
		fmt.Println(input)
		if input[indexChildren] > 0 {
			length := len(input)

			node := node{input[indexChildren], input[indexMetadata], sum(input[(length-input[indexMetadata]):length])}
			nodeList = append(nodeList, node)
			fmt.Println(node)

			input = input[headerLen:(length - input[indexMetadata])]
		} else {
			node := node{input[indexChildren], input[indexMetadata], sum(input[headerLen:headerLen+input[indexMetadata]])}
			nodeList = append(nodeList, node)
			fmt.Println(node)
			input = input[headerLen+input[indexMetadata]:]
		}
	}
	return nodeList
}

func calcTotalMetadata(nodes []node) int {
	sum := 0
	for node := range nodes {
		sum += nodes[node].data
	}
	return sum
}

func main() {
	//file, err := ioutil.ReadFile("../input.txt")
	//file, err := ioutil.ReadFile("../testinput.txt") // this goes well

	/*
	file testinputwitherror.txt
	2 3 1 3 0 1 11 10 11 12 1 1 0 1 99 2 1 1 2
	A-----------------------------------------
		B------------------ C-----------
			D-----  			 E----
	*/

	file, err := ioutil.ReadFile("../testinputwitherror.txt") // this fails
	checkError(err)
	data := strings.Split(string(file), " ")
	dataInts := convertToInt(data)

	nodes := createNodes(dataInts)

	fmt.Println("The sum of all metadata is", calcTotalMetadata(nodes))
}
