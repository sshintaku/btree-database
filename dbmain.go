package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Animal struct {
	Species  string `json:"Species"`
	Location string `json:"Location"`
}

type Node struct {
	RightNode *Node
	LeftNode  *Node
	Data      Animal
}

func main() {
	file, _ := ioutil.ReadFile("./data1.json")
	var animalList []Animal
	json.Unmarshal(file, &animalList)
	sort.SliceStable(animalList, func(i, j int) bool { return animalList[i].Species < animalList[j].Species })
	tree := CreateTree(animalList)
	WriteDatabase(tree)
	rTree := ReadDatabase()
	fmt.Printf("Read first node data: %s", rTree.Data)

}

func ReadDatabase() Node {
	db, err := ioutil.ReadFile("db.bin")
	checkError(err)
	var node Node
	json.Unmarshal(db, &node)
	return node
}

func WriteDatabase(data Node) {
	dBytes, _ := json.Marshal(&data)
	err := ioutil.WriteFile("db.bin", dBytes, 0644)
	checkError(err)

}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetMiddleIndex(animalList []Animal) *int {
	length := len(animalList)
	answer := length%2 == 0
	if !answer {
		length++
	}
	value := length / 2
	if value < 1 {
		return nil
	}
	return &value
}

func SplitArray(list []Animal) ([]Animal, []Animal) {
	middleIndex := GetMiddleIndex(list)
	rightSplit := list[0 : *middleIndex-1]
	leftSplit := list[*middleIndex+1 : len(list)]
	return rightSplit, leftSplit
}

func CreateTree(animalList []Animal) Node {
	var rootNode Node
	var rightNode Node
	var leftNode Node
	right, left := SplitArray(animalList)
	rootNode.Data = animalList[*GetMiddleIndex(animalList)]
	rootNode.RightNode = &rightNode
	rootNode.LeftNode = &leftNode
	firstChild := CreateSubTree(*rootNode.RightNode, right)
	rootNode.RightNode = &firstChild
	secondChild := CreateSubTree(*rootNode.LeftNode, left)
	rootNode.LeftNode = &secondChild
	return rootNode

}

func CreateSubTree(node Node, list []Animal) Node {
	if len(list) == 1 {
		node.Data = list[0]
		return node
	}

	if len(list) == 2 {
		node.Data = list[1]
		var rightNode Node
		rightNode.Data = list[0]
		node.RightNode = &rightNode
		return node
	}

	node.Data = list[*GetMiddleIndex(list)]
	if node.LeftNode == nil {
		var child Node
		node.LeftNode = &child
	}
	if node.RightNode == nil {
		var child Node
		node.RightNode = &child
	}
	right, left := SplitArray(list)
	firstChild := CreateSubTree(*node.RightNode, right)
	node.RightNode = &firstChild
	secondChild := CreateSubTree(*node.LeftNode, left)
	node.LeftNode = &secondChild

	return node

}
