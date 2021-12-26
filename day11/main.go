package day11

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
)

type Octopus struct {
	energyLevel    int
	hasFlashed     bool
	adjacentOctopi []*Octopus
}

type Point struct {
	x, y int
}

func ReadFileToOctopi(fileName string) [][]*Octopus {
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
	log.Printf("%v\n", ints)

	octopi := [][]*Octopus{}
	for y := 0; y < len(ints); y++ {
		octopi = append(octopi, []*Octopus{})
		for x := 0; x < len(ints[y]); x++ {
			octopus := &Octopus{energyLevel: ints[y][x], hasFlashed: false, adjacentOctopi: []*Octopus{}}
			octopi[y] = append(octopi[y], octopus)
		}
	}
	for y := 0; y < len(ints); y++ {
		for x := 0; x < len(ints[y]); x++ {
			adjacentPoints := getAdjacentPoints(ints, x, y)
			currentOctopus := octopi[y][x]
			for _, pt := range adjacentPoints {
				currentOctopus.adjacentOctopi = append(currentOctopus.adjacentOctopi, octopi[pt.y][pt.x])
			}
		}
	}

	return octopi
}

func IncrementEnergyLevels(octopi [][]*Octopus) {
	for y := 0; y < len(octopi); y++ {
		for x := 0; x < len(octopi[y]); x++ {
			octopi[y][x].energyLevel++
		}
	}
}

func FlashAndIncrementAdjacent(octopi [][]*Octopus) int {
	flashCount := 0
	for y := 0; y < len(octopi); y++ {
		for x := 0; x < len(octopi[y]); x++ {
			octopus := octopi[y][x]
			if octopus.energyLevel > 9 {
				if !octopus.hasFlashed {
					flashCount++
					octopus.hasFlashed = true
					for _, o := range octopus.adjacentOctopi {
						o.energyLevel++
					}
				}
			}
		}
	}
	return flashCount
}

func ResetFlashedOctopi(octopi [][]*Octopus) {
	for y := 0; y < len(octopi); y++ {
		for x := 0; x < len(octopi[y]); x++ {
			octopus := octopi[y][x]
			if octopus.hasFlashed {
				octopus.hasFlashed = false
				octopus.energyLevel = 0
			}
		}
	}
}

func RunOneStep(octopi [][]*Octopus) int {
	IncrementEnergyLevels(octopi)
	flashSum := 0
	flashCount := FlashAndIncrementAdjacent(octopi)
	flashSum += flashCount
	for flashCount > 0 {
		flashCount = FlashAndIncrementAdjacent(octopi)
		flashSum += flashCount
	}
	ResetFlashedOctopi(octopi)
	return flashSum
}

func Solve(octopi [][]*Octopus) int {
	flashSum := 0
	for i := 0; i < 100; i++ {
		flashCount := RunOneStep(octopi)
		flashSum += flashCount
		log.Printf("Step %d Flashes %d Total %d\n", i, flashCount, flashSum)
	}
	return flashSum
}

func SolvePart2(octopi [][]*Octopus) int {
	flashCount := 0
	stepCount := 0
	for flashCount != 100 {
		stepCount++
		flashCount = RunOneStep(octopi)
		log.Printf("Step %d Flashes %d\n", stepCount, flashCount)
	}
	return stepCount
}

/*
 * Copied from Day 9 but add getting diagonals
 */
func getAdjacentPoints(ints [][]int, x int, y int) []Point {
	adjacentPoints := []Point{}
	upRow, upRowErr := getSlice(ints, y-1)
	if upRowErr == nil {
		_, upErr := get(upRow, x)
		if upErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x, y - 1})
		}
		_, upRightErr := get(upRow, x+1)
		if upRightErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x + 1, y - 1})
		}
		_, upLeftErr := get(upRow, x-1)
		if upLeftErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x - 1, y - 1})
		}
	}

	downRow, downRowErr := getSlice(ints, y+1)
	if downRowErr == nil {
		_, downErr := get(downRow, x)
		if downErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x, y + 1})
		}
		_, downRightErr := get(downRow, x+1)
		if downRightErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x + 1, y + 1})
		}
		_, downLeftErr := get(downRow, x-1)
		if downLeftErr == nil {
			adjacentPoints = append(adjacentPoints, Point{x - 1, y + 1})
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
