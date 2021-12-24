package day9

import (
	"errors"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func ReadFileToIntSlices(fileName string) [][]int {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	ints := [][]int{}
	currentRow := []int{}
	for _, fileByte := range fileBytes {
		byteAsStr, _ := strconv.Atoi(string(fileByte))
		if string(fileByte) == "\n" {
			ints = append(ints, currentRow)
			currentRow = []int{}
		} else {
			currentRow = append(currentRow, byteAsStr)
		}

	}
	return ints
}

func Solve(ints [][]int) int {
	lowPoints := getLowPoints(ints)
	sum := 0
	for _, pt := range lowPoints {
		sum += ints[pt.y][pt.x] + 1
	}
	return sum
}

func SolvePart2(ints [][]int) int {
	lowPoints := getLowPoints(ints)
	basinSizes := []int{}
	for _, pt := range lowPoints {
		basinSizes = append(basinSizes, GetBasinSize(ints, pt))
	}
	sort.SliceStable(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	log.Printf("Largest basins: %d, %d, %d\n", basinSizes[0], basinSizes[1], basinSizes[2])
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

type Point struct {
	x, y int
}

func GetBasinSize(ints [][]int, point Point) int {
	basin := [][]bool{}
	for i := 0; i < len(ints); i++ {
		basin = append(basin, []bool{})
		for j := 0; j < len(ints[i]); j++ {
			basin[i] = append(basin[i], false)
		}
	}
	pointsToExplore := []Point{point}
	for len(pointsToExplore) > 0 {
		current := pointsToExplore[0]
		pointsToExplore = pointsToExplore[1:]
		if ints[current.y][current.x] != 9 && !basin[current.y][current.x] {
			basin[current.y][current.x] = true
			pointsToExplore = append(pointsToExplore, getAdjacentPoints(ints, current.x, current.y)...)
		}
	}
	basinSize := 0
	for y := 0; y < len(basin); y++ {
		for x := 0; x < len(basin[y]); x++ {
			if basin[y][x] {
				log.Printf("basin: (%d, %d)\n", x, y)
				basinSize++
			}
		}
	}
	log.Printf("%d\n", basinSize)
	return basinSize
}

func getLowPoints(ints [][]int) []Point {
	lowPoints := []Point{}
	for y := 0; y < len(ints); y++ {
		for x := 0; x < len(ints[y]); x++ {
			current := ints[y][x]
			adjacentPoints := getAdjacentPoints(ints, x, y)
			isLow := true
			for _, pt := range adjacentPoints {
				if current >= ints[pt.y][pt.x] {
					isLow = false
				}
			}
			if isLow {
				lowPoints = append(lowPoints, Point{x, y})
			}
		}
	}
	return lowPoints
}

func getAdjacentPoints(ints [][]int, x int, y int) []Point {
	adjacentPoints := []Point{}
	upRow, upRowErr := getSlice(ints, y-1)
	if upRowErr == nil {
		_, upErr := get(upRow, x)
		if upErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x, y - 1})
		}
	}

	downRow, downRowErr := getSlice(ints, y+1)
	if downRowErr == nil {
		_, downErr := get(downRow, x)
		if downErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x, y + 1})
		}
	}

	_, rightErr := get(ints[y], x+1)
	if rightErr == nil {
		adjacentPoints = append(adjacentPoints, Point{x + 1, y})
	}

	_, leftErr := get(ints[y], x-1)
	if leftErr == nil {
		adjacentPoints = append(adjacentPoints, Point{x - 1, y})
	}

	return adjacentPoints
}

// get number and return error instead of panicing if out of bounds
func get(ints []int, index int) (value int, err error) {
	value = -1
	err = errors.New("out of bounds")
	defer handleOutOfBounds()
	value = ints[index]
	return value, nil
}

func getSlice(ints [][]int, index int) (value []int, err error) {
	value = []int{}
	err = errors.New("out of bounds")
	defer handleOutOfBounds()
	value = ints[index]
	return value, nil
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		log.Println("Recovering from panic:", r)
	}
}
