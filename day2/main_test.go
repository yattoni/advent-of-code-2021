package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToDirections(t *testing.T) {
	instructions := []Instruction{
		{Forward, 5},
		{Down, 5},
		{Forward, 8},
		{Up, 3},
		{Down, 8},
		{Forward, 2},
	}
	assert.Equal(t, instructions, ReadFileToInstructions("prompt-input"))
}

func TestSolveWithPromptInput(t *testing.T) {
	instructions := ReadFileToInstructions("prompt-input")
	assert.Equal(t, 150, Solve(instructions))
}

func TestSolve(t *testing.T) {
	instructions := ReadFileToInstructions("input")
	assert.Equal(t, 1868935, Solve(instructions))
}

func TestSolveWithAimWithPromptInput(t *testing.T) {
	instructions := ReadFileToInstructions("prompt-input")
	assert.Equal(t, 900, SolveWithAim(instructions))
}

func TestSolveWithAim(t *testing.T) {
	instructions := ReadFileToInstructions("input")
	assert.Equal(t, 1965970888, SolveWithAim(instructions))
}
