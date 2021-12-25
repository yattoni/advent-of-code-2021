package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToLines(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Len(t, lines, 10)
	assert.Equal(t, "[({(<(())[]>[[{[]{<()<>>", lines[0])
	assert.Equal(t, "<{([{{}}[<[[[<>{}]]]>[]]", lines[9])
}

func TestIsCorrupt(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Equal(t, "0", IsCorrupt(lines[0]))
	assert.Equal(t, "0", IsCorrupt(lines[1]))
	assert.Equal(t, "}", IsCorrupt(lines[2]))
	assert.Equal(t, "0", IsCorrupt(lines[3]))
	assert.Equal(t, ")", IsCorrupt(lines[4]))
	assert.Equal(t, "]", IsCorrupt(lines[5]))
	assert.Equal(t, "0", IsCorrupt(lines[6]))
	assert.Equal(t, ")", IsCorrupt(lines[7]))
	assert.Equal(t, ">", IsCorrupt(lines[8]))
	assert.Equal(t, "0", IsCorrupt(lines[9]))
}

func TestSolvePromptInput(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Equal(t, 26397, Solve(lines))
}

func TestSolve(t *testing.T) {
	lines := ReadFileToLines("input")
	assert.Equal(t, 341823, Solve(lines))
}

func TestAutoComplete(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Equal(t, "}}]])})]", AutoComplete(lines[0]))
	assert.Equal(t, ")}>]})", AutoComplete(lines[1]))
	assert.Equal(t, "}}>}>))))", AutoComplete(lines[3]))
	assert.Equal(t, "]]}}]}]}>", AutoComplete(lines[6]))
	assert.Equal(t, "])}>", AutoComplete(lines[9]))
}

func TestScoreAutoComplete(t *testing.T) {
	assert.Equal(t, 288957, ScoreAutoComplete("}}]])})]"))
	assert.Equal(t, 5566, ScoreAutoComplete(")}>]})"))
	assert.Equal(t, 1480781, ScoreAutoComplete("}}>}>))))"))
	assert.Equal(t, 995444, ScoreAutoComplete("]]}}]}]}>"))
	assert.Equal(t, 294, ScoreAutoComplete("])}>"))
}

func TestSolvePart2PromptInput(t *testing.T) {
	lines := ReadFileToLines("prompt-input")
	assert.Equal(t, 288957, SolvePart2(lines))
}

func TestSolvePart2(t *testing.T) {
	lines := ReadFileToLines("input")
	assert.Equal(t, 2801302861, SolvePart2(lines))
}
