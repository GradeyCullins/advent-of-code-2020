package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Number of valid passwords.
	validCount := 0

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")
	for _, line := range lines {
		bounds, target, str, err := parseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		if valid := checkStr(bounds, target, str); valid {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func checkStr(bounds [2]int, target, str string) bool {
	count := strings.Count(str, target)
	return count >= bounds[0] && count <= bounds[1]
}

func parseLine(line string) ([2]int, string, string, error) {
	parts := strings.Split(line, " ")
	rule := parts[0]
	target := strings.Replace(parts[1], ":", "", 1)
	str := parts[2]
	bounds, err := parseRange(rule)
	if err != nil {
		return [2]int{}, "", "", err
	}
	return bounds, target, str, nil
}

func parseRange(str string) ([2]int, error) {
	var res [2]int
	var err error
	parts := strings.Split(str, "-")
	if len(parts) == 1 {
		return [2]int{}, fmt.Errorf("%s does not contain a \"-\"", str)
	}
	for i := 0; i < len(res); i++ {
		res[i], err = strconv.Atoi(parts[i])
		if err != nil {
			return [2]int{}, err
		}
	}
	return res, nil
}
