package day5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToLines(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Len(t, lines, 10)
	assert.Equal(t, Line{Coordinate{0, 9}, Coordinate{5, 9}}, lines[0])
	assert.Equal(t, Line{Coordinate{5, 5}, Coordinate{8, 2}}, lines[9])
}

func TestFilterStraightLines(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	straightLines := FilterStraightLines(lines)
	assert.Len(t, straightLines, 6)
}

func TestDrawLines(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	straightLines := FilterStraightLines(lines)
	SortLines(straightLines)
	drawn := DrawLines(straightLines)
	for _, d := range drawn {
		fmt.Println(d)
	}
	// assert.Fail(t, "")
}

func TestSolvePromptInput(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Equal(t, 5, Solve(lines))
}

func TestSolveInput(t *testing.T) {
	lines := ReadFileToLines("input")
	assert.Equal(t, 5147, Solve(lines))
}
