package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Start in the top left.
	position := 0

	// Count number of trees we shmack.
	numTrees := 0

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")

	for _, line := range lines {
		lineArr := strings.Split(line, "")
		currChar := lineArr[position]
		if currChar == "#" {
			// fmt.Println("All good")
			numTrees++
		}
		position = (position + 3) % len(lineArr)
	}
	fmt.Println(numTrees)
}
