package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func FilterDiagonalLines(lines []Line) []Line {
	diagonalLines := []Line{}
	for _, line := range lines {
		if math.Abs(float64(line.start.x-line.end.x)) == math.Abs(float64(line.start.y-line.end.y)) {
			fmt.Printf("Line is diagonal: %v\n", line)
			diagonalLines = append(diagonalLines, line)
		}
	}
	return diagonalLines
}

// Flip line coordinates where the end is farther left or up than the start so draw can traverse each line in the same direction, right and down
// Assuming 0,0 is top left
func SortLines(lines []Line) {
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line.start.x > line.end.x || line.start.y > line.end.y {
			fmt.Printf("Flipping line %v\n", line)
			lines[i] = Line{Coordinate{line.end.x, line.end.y}, Coordinate{line.start.x, line.start.y}}
		}
	}
}

func DrawLines(straigntLines, diagonalLines []Line) [][]int {
	maxXStraight := findMaxX(straigntLines)
	maxYStraight := findMaxY(straigntLines)
	maxXDiagonal := findMaxX(diagonalLines)
	maxYDiagonal := findMaxY(diagonalLines)

	maxX := 0
	if maxXStraight > maxXDiagonal {
		maxX = maxXStraight
	} else {
		maxX = maxXDiagonal
	}

	maxY := 0
	if maxYStraight > maxYDiagonal {
		maxY = maxYStraight
	} else {
		maxY = maxYDiagonal
	}

	lineMap := [][]int{}

	for i := 0; i < maxY; i++ {
		lineMap = append(lineMap, []int{})
		for j := 0; j < maxX; j++ {
			lineMap[i] = append(lineMap[i], 0)
		}
	}

	for _, line := range straigntLines {
		fmt.Printf("line: %v\n", line)
		for i := line.start.x; i <= line.end.x; i++ {
			for j := line.start.y; j <= line.end.y; j++ {
				fmt.Printf("i = %d j = %d\n", i, j)
				lineMap[j][i] += 1
			}
		}
	}

	for _, line := range diagonalLines {
		fmt.Printf("diagonal line: %v\n", line)
		j := line.start.y
		for i := line.start.x; i <= line.end.x; i++ {
			fmt.Printf("i = %d j = %d\n", i, j)
			lineMap[j][i] += 1
			if line.start.y < line.end.y {
				j++
			} else {
				j--
			}
		}
		j = line.start.y
		for i := line.start.x; i >= line.end.x; i-- {
			fmt.Printf("i = %d j = %d\n", i, j)
			lineMap[j][i] += 1
			if line.start.y < line.end.y {
				j++
			} else {
				j--
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
	drawn := DrawLines(straightLines, []Line{})

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

func SolvePart2(lines []Line) int {
	straightLines := FilterStraightLines(lines)
	diagonalLines := FilterDiagonalLines(lines)

	SortLines(straightLines)
	SortLines(diagonalLines)

	drawn := DrawLines(straightLines, diagonalLines)

	for _, d := range drawn {
		fmt.Println(d)
	}

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
