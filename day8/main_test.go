package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	lines := ReadFile("prompt-input")
	assert.Len(t, lines, 10)
}

func TestSolvePromptInput(t *testing.T) {
	lines := ReadFile("prompt-input")
	assert.Equal(t, 26, Solve(lines))
}

func TestSolve(t *testing.T) {
	lines := ReadFile("input")
	assert.Equal(t, 493, Solve(lines))
}

func TestGetWiresForOne(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "ab", GetWiresForOne(wires))
}

func TestGetWiresForFour(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "eafb", GetWiresForFour(wires))
}

func TestGetWiresForSeven(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "dab", GetWiresForSeven(wires))
}

func TestGetWiresForEight(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "acedgfb", GetWiresForEight(wires))
}

func TestGetWiresForThree(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "fbcad", GetWiresForThree(wires))
}

func TestGetWiresForFive(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "cdfbe", GetWiresForFive(wires))
}

func TestGetWiresForTwo(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "gcdfa", GetWiresForTwo(wires))
}

func TestGetWiresForSix(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "cdfgeb", GetWiresForSix(wires))
}

func TestGetWiresForNine(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "cefabd", GetWiresForNine(wires))
}

func TestGetWiresForZero(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	assert.Equal(t, "cagedb", GetWiresForZero(wires))
}

func TestDecodeLine(t *testing.T) {
	wires := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	output := []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}
	assert.Equal(t, 5353, DecodeLine(Line{wires, output}))
}

func TestDecodeDebug(t *testing.T) {
	lines := ReadFile("prompt-input")
	assert.Equal(t, 8394, DecodeLine(lines[0]))
}

func TestSolvePart2PromptInput(t *testing.T) {
	lines := ReadFile("prompt-input")
	assert.Equal(t, 61229, SolvePart2(lines))
}

func TestSolvePart2(t *testing.T) {
	lines := ReadFile("input")
	assert.Equal(t, 1010460, SolvePart2(lines))
}
