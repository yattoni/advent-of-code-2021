package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToIntSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, ReadFileToIntSlice("test-input"))
}

func TestSolveWithPromptInput(t *testing.T) {
	input := ReadFileToIntSlice("prompt-input")
	assert.Equal(t, 7, Solve(input))
}

func TestSolve(t *testing.T) {
	input := ReadFileToIntSlice("input")
	assert.Equal(t, 1228, Solve(input))
}

func TestSplitIntoSlidingWindows(t *testing.T) {
	input := ReadFileToIntSlice("test-input")
	windows := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5, 6},
	}
	assert.Equal(t, windows, SplitIntoSlidingWindows(input, 3))
}

func TestSumSlidingWindows(t *testing.T) {
	input := ReadFileToIntSlice("test-input")
	windows := SplitIntoSlidingWindows(input, 3)
	assert.Equal(t, []int{6, 9, 12, 15}, SumSlidingWindows(windows))
}

func TestSolvePart2(t *testing.T) {
	input := ReadFileToIntSlice("input")
	windows := SplitIntoSlidingWindows(input, 3)
	sums := SumSlidingWindows(windows)
	assert.Equal(t, 1257, Solve(sums))
}
