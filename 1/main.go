package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Find the two nums that add up to 2020.
	total := 2020

	res := make([]int, 3)

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")

	nums, err := strArrToInt(lines)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	res[0], res[1], err = sumOfTwo(nums, total)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:")
	fmt.Println(res[0] * res[1])

	// Part 2
	res[0], res[1], res[2], err = sumOfThree(nums, total)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2:")
	fmt.Println(res[0] * res[1] * res[2])
}

func strArrToInt(lines []string) ([]int, error) {
	var err error
	nums := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		nums[i], err = strconv.Atoi(lines[i])
		if err != nil {
			return nil, err
		}
	}
	return nums, nil
}

func sumOfTwo(nums []int, total int) (int, int, error) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == total {
				return nums[i], nums[j], nil
			}
		}
	}
	return 0, 0, fmt.Errorf("Failed to find two numbers that sum to: %d", total)
}

func sumOfThree(nums []int, total int) (int, int, int, error) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == total {
					return nums[i], nums[j], nums[k], nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("Failed to find two numbers that sum to: %d", total)
}
