package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Line struct {
	start Coordinate
	end   Coordinate
}

type Coordinate struct {
	x int
	y int
}

func ReadFileToLines(fileName string) []Line {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.FieldsFunc(string(fileBytes), func(r rune) bool { return !unicode.IsDigit(r) })
	fmt.Println(fileFields)
	fmt.Println(len(fileFields))

	fileFieldNumbers := []int{}

	for _, field := range fileFields {
		n, err := strconv.Atoi(field)
		if err == nil {
			fileFieldNumbers = append(fileFieldNumbers, n)
		}
	}

	lines := []Line{}
	fmt.Println(fileFieldNumbers)
	for i := 0; i < len(fileFieldNumbers); i += 4 {
		lines = append(lines, Line{Coordinate{fileFieldNumbers[i], fileFieldNumbers[i+1]}, Coordinate{fileFieldNumbers[i+2], fileFieldNumbers[i+3]}})
	}

	return lines
}

func FilterStraightLines(lines []Line) []Line {
	straightLines := []Line{}
	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			straightLines = append(straightLines, line)
		}
	}
	return straightLines
}

// Flip line coordinates where the end is farther left or up than the start so draw can traverse each line in the same direction, right and down
// Assuming 0,0 is top left
func SortLines(lines []Line) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line.start.x > line.end.x || line.start.y > line.end.y {
			lines[i] = Line{Coordinate{line.end.x, line.end.y}, Coordinate{line.start.x, line.start.y}}
		}
	}
}

func DrawLines(lines []Line) [][]int {
	maxX := findMaxX(lines)
	maxY := findMaxY(lines)

	lineMap := [][]int{}

	for i := 0; i < maxY; i++ {
		lineMap = append(lineMap, []int{})
		for j := 0; j < maxX; j++ {
			lineMap[i] = append(lineMap[i], 0)
		}
	}

	for _, line := range lines {
		fmt.Printf("line: %v\n", line)
		for i := line.start.x; i <= line.end.x; i++ {
			for j := line.start.y; j <= line.end.y; j++ {
				fmt.Printf("i = %d j = %d\n", i, j)
				lineMap[j][i] += 1
			}
		}
	}

	return lineMap
}

func findMaxX(lines []Line) int {
	max := 0
	for _, line := range lines {
		if line.start.x > max {
			max = line.start.x
		}
		if line.end.x > max {
			max = line.end.x
		}
	}
	return max + 1
}

func findMaxY(lines []Line) int {
	max := 0
	for _, line := range lines {
		if line.start.y > max {
			max = line.start.y
		}
		if line.end.y > max {
			max = line.end.y
		}
	}
	return max + 1
}

func Solve(lines []Line) int {
	straightLines := FilterStraightLines(lines)
	SortLines(straightLines)
	drawn := DrawLines(straightLines)

	pointsGreaterThan2 := 0

	for i := 0; i < len(drawn); i++ {
		for j := 0; j < len(drawn[i]); j++ {
			if drawn[i][j] >= 2 {
				pointsGreaterThan2++
			}
		}
	}
	return pointsGreaterThan2
}
