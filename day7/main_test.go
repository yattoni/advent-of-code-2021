package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToCrabPositions(t *testing.T) {
	positions := ReadFileToCrabPositions("prompt-input")
	assert.Equal(t, []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, positions)
}

func TestSumFuelCostToTravelToPosition(t *testing.T) {
	positions := ReadFileToCrabPositions("prompt-input")
	fuel := SumFuelCostToTravelToPosition(positions, 2)
	assert.Equal(t, 37, fuel)
}

func TestSumFuelCostToTravelToPositionPart2(t *testing.T) {
	positions := ReadFileToCrabPositions("prompt-input")
	fuel := SumFuelCostToTravelToPositionPart2(positions, 2)
	assert.Equal(t, 206, fuel)
}

func TestSolvePromptInput(t *testing.T) {
	positions := ReadFileToCrabPositions("prompt-input")
	assert.Equal(t, 37, Solve(positions, SumFuelCostToTravelToPosition))
}

func TestSolve(t *testing.T) {
	positions := ReadFileToCrabPositions("input")
	assert.Equal(t, 352707, Solve(positions, SumFuelCostToTravelToPosition))
}

func TestSolvePromptInputPart2(t *testing.T) {
	positions := ReadFileToCrabPositions("prompt-input")
	assert.Equal(t, 168, Solve(positions, SumFuelCostToTravelToPositionPart2))
}

func TestSolvePart2(t *testing.T) {
	positions := ReadFileToCrabPositions("input")
	assert.Equal(t, 95519693, Solve(positions, SumFuelCostToTravelToPositionPart2))
}
