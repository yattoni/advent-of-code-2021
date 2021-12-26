package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToOctopi(t *testing.T) {
	octopi := ReadFileToOctopi("prompt-input")
	assert.Len(t, octopi, 10)
	assert.Len(t, octopi[0], 10)

	assert.Equal(t, 5, octopi[0][0].energyLevel)
	assert.Equal(t, 7, octopi[1][1].energyLevel)
	assert.Equal(t, 4, octopi[0][1].energyLevel)
	assert.Equal(t, 2, octopi[1][0].energyLevel)

	assert.Len(t, octopi[0][0].adjacentOctopi, 3)
	assert.Len(t, octopi[1][1].adjacentOctopi, 8)
	assert.Len(t, octopi[0][1].adjacentOctopi, 5)
	assert.Len(t, octopi[1][0].adjacentOctopi, 5)
}

func TestSolvePromptInput(t *testing.T) {
	octopi := ReadFileToOctopi("prompt-input")
	assert.Equal(t, 1656, Solve(octopi))
}

func TestSolve(t *testing.T) {
	octopi := ReadFileToOctopi("input")
	assert.Equal(t, 1741, Solve(octopi))
}

func TestSolvePart2PromptInput(t *testing.T) {
	octopi := ReadFileToOctopi("prompt-input")
	assert.Equal(t, 195, SolvePart2(octopi))
}

func TestSolvePart2Input(t *testing.T) {
	octopi := ReadFileToOctopi("input")
	assert.Equal(t, 440, SolvePart2(octopi))
}
