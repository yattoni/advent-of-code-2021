package day1

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFileToIntSlice(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	// Split includes a trailing empty string since the file ends in a new line
	// intsAsStrings := strings.Split(fileString, "\n")
	intsAsStrings := strings.Fields(string(fileBytes))

	ints := []int{}

	for _, intAsString := range intsAsStrings {
		n, err := strconv.Atoi(intAsString)
		if err != nil {
			log.Fatal("Error converting ", intAsString, " to int", err)
		}
		ints = append(ints, n)
	}

	return ints
}

func Solve(data []int) int {
	depthIncreaseCount := 0

	// start on index 1 since 0 has no previous measurement
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			depthIncreaseCount++
		}
	}

	return depthIncreaseCount
}

func SplitIntoSlidingWindows(data []int, windowSize int) [][]int {
	windows := [][]int{}

	for i := 0; i < len(data)-windowSize+1; i++ {
		window := []int{data[i], data[i+1], data[i+2]}
		windows = append(windows, window)
	}

	return windows
}

func SumSlidingWindows(data [][]int) []int {
	sums := []int{}

	for i := 0; i < len(data); i++ {
		sum := 0
		for j := 0; j < len(data[i]); j++ {
			sum += data[i][j]
		}
		sums = append(sums, sum)
	}

	return sums
}
