package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	treeCounts := make([]int, 0)

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		treeCounts = append(treeCounts, doSlope(lines, slope[0], slope[1]))
	}

	total := treeCounts[0]
	for i := 1; i < len(treeCounts); i++ {
		total = total * treeCounts[i]
	}

	fmt.Println(total)
}

// Return how many trees get bonked by following the slope.
func doSlope(lines []string, right, down int) int {
	// Start in the top left.
	position := 0

	// Count number of trees we shmack.
	trees := 0

	for i := 0; i < len(lines); i++ {
		lineArr := strings.Split(lines[i], "")
		currChar := lineArr[position]
		if currChar == "#" {
			trees++
		}

		position = (position + right) % len(lineArr)

		if down > 1 {
			i = i + (down - 1)
		}
	}
	return trees
}
