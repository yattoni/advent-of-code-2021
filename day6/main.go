package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type DayRunnerFunc func([]int) []int

type FishCounterFunc func([]int) int

func ReadFileToFish(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.FieldsFunc(string(fileBytes), func(r rune) bool { return !unicode.IsDigit(r) })
	fmt.Println(fileFields)

	fish := []int{}
	for _, field := range fileFields {
		num, err := strconv.Atoi(field)
		if err == nil {
			fish = append(fish, num)
		}
	}
	return fish
}

func ReadFileToEfficientFish(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.FieldsFunc(string(fileBytes), func(r rune) bool { return !unicode.IsDigit(r) })
	fmt.Println(fileFields)

	fish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, field := range fileFields {
		num, err := strconv.Atoi(field)
		if err == nil {
			fish[num]++
		}
	}
	return fish
}

func RunOneDay(fish []int) []int {
	newFishToAdd := 0

	for i := 0; i < len(fish); i++ {
		if fish[i] == 0 {
			fish[i] = 6
			newFishToAdd++
		} else {
			fish[i]--
		}
	}

	for i := 0; i < newFishToAdd; i++ {
		fish = append(fish, 8)
	}
	return fish
}

func RunOneEfficientDay(fish []int) []int {
	newFish := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	newFishToAdd := fish[0]

	for i := 1; i < len(fish); i++ {
		newFish[i-1] = fish[i]
	}

	newFish[6] += newFishToAdd
	newFish[8] += newFishToAdd

	return newFish
}

func Solve(initialFish []int, days int, runDay DayRunnerFunc, counter FishCounterFunc) int {
	currentFish := initialFish
	for i := 0; i < days; i++ {
		log.Printf("Running day: %d. Current length of fish: %d\n", i, len(currentFish))
		currentFish = runDay(currentFish)
	}
	// log.Printf("Final Fish: %v\n", currentFish)
	return counter(currentFish)
}

func CountFish(fish []int) int {
	return len(fish)
}

func CountEfficientFish(fish []int) int {
	numFish := 0

	for i := 0; i < len(fish); i++ {
		numFish += fish[i]
	}

	return numFish
}
