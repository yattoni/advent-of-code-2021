package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type CostFinderFunc func([]int, int) int

func ReadFileToCrabPositions(fileName string) []int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.FieldsFunc(string(fileBytes), func(r rune) bool { return !unicode.IsDigit(r) })
	fmt.Println(fileFields)

	positions := []int{}
	for _, field := range fileFields {
		num, err := strconv.Atoi(field)
		if err == nil {
			positions = append(positions, num)
		}
	}
	return positions
}

func SumFuelCostToTravelToPosition(positions []int, target int) int {
	sum := 0
	for _, n := range positions {
		sum += int(math.Abs(float64(n) - float64(target)))
	}
	return sum
}

/*
1,1
2,3
3,6
4,10
5,15
6,21
7,28
8,36
9,45
10,55
11,66

= n(n+1) / 2
*/
func SumFuelCostToTravelToPositionPart2(positions []int, target int) int {
	sum := 0
	for _, n := range positions {
		difference := int(math.Abs(float64(n) - float64(target)))
		sum += (difference * (difference + 1)) / 2
	}
	return sum
}

func Solve(positions []int, costFinder CostFinderFunc) int {
	min := findSliceMin(positions)
	max := findSliceMax(positions)

	leastFuel := math.MaxInt

	for i := min; i <= max; i++ {
		fuel := costFinder(positions, i)
		if fuel < leastFuel {
			leastFuel = fuel
		}
	}

	return leastFuel
}

func findSliceMax(slice []int) int {
	max := 0
	for _, n := range slice {
		if n > max {
			max = n
		}
	}
	return max
}

func findSliceMin(slice []int) int {
	min := 0
	for _, n := range slice {
		if n < min {
			min = n
		}
	}
	return min
}
