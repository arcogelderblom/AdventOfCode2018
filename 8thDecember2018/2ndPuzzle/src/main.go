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

type Node struct {
	index int
	childrenNum int
	children []Node
	metadataNum int
	metadata []int
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

func getNode(index int, split []int) Node {
	node := Node{index: index, childrenNum: split[index], metadataNum: split[index+1]}

	offset := node.index + 2

	for i := 0; i < node.childrenNum ; i++ {
		childNode := getNode(offset, split)
		node.children = append(node.children, childNode)
		offset = offset + getLength(childNode)
	}

	for i := 0; i < node.metadataNum ; i++ {
		node.metadata = append(node.metadata,split[offset + i])
	}

	return node
}

func getLength(node Node) int {
	length := 2
	for i := 0; i < node.childrenNum ; i++ {
		length = length + getLength(node.children[i])
	}
	length = length + node.metadataNum
	return length
}

func sum(node Node) int {
	total := 0
	for _,v := range node.children {
		total = total + sum(v)
	}
	for _,v := range node.metadata {
		total = total + v
	}
	return total
}

func sumInts(input []int) int {
	sum := 0
	for index := range input {
		sum += input[index]
	}
	return sum
}

func getNodeValue(node Node) int {
	if node.childrenNum == 0 {
		return sumInts(node.metadata)
	} else {
		total := 0
		for i := 0; i < node.metadataNum; i ++ {
			if node.metadata[i] > node.childrenNum {
				total += 0
			} else {
				total += getNodeValue(node.children[node.metadata[i]-1])
			}
		}
		return total
	}
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	data := strings.Split(string(file), " ")
	dataInInt := convertToInt(data)

	nodes := getNode(0, dataInInt) // remember to split multiple child nodes

	fmt.Println("The sum of all metadata of the root node is", getNodeValue(nodes))
}
