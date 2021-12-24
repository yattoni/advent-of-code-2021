package day9

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToIntSlices(t *testing.T) {
	ints := ReadFileToIntSlices("prompt-input")
	fmt.Println(ints)
	assert.Equal(t, 2, ints[0][0])
	assert.Equal(t, 1, ints[0][1])
	assert.Equal(t, 3, ints[1][0])
	assert.Len(t, ints, 5)
	assert.Len(t, ints[4], 10)
}

func TestSolvePromptInput(t *testing.T) {
	ints := ReadFileToIntSlices("prompt-input")
	assert.Equal(t, 15, Solve(ints))
}

func TestSolve(t *testing.T) {
	ints := ReadFileToIntSlices("input")
	assert.Equal(t, 535, Solve(ints))
}

func TestGetBasinSize(t *testing.T) {
	ints := ReadFileToIntSlices("prompt-input")
	assert.Equal(t, 14, GetBasinSize(ints, Point{3, 3}))
}

func TestSolvePart2PromptInput(t *testing.T) {
	ints := ReadFileToIntSlices("prompt-input")
	assert.Equal(t, 1134, SolvePart2(ints))
}

func TestSolvePart2(t *testing.T) {
	ints := ReadFileToIntSlices("input")
	assert.Equal(t, 1122700, SolvePart2(ints))
}
