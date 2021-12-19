package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToUint16s(t *testing.T) {
	ints := []uint16{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}
	assert.Equal(t, ints, ReadFileToUint16s("prompt-input"))
}

func TestSolveGammaWithPromptInput(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")
	assert.Equal(t, uint16(0b10110), SolveGamma(input))
}

func TestSolveEpsilonWithPromptInput(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")
	assert.Equal(t, uint16(0b01001), SolveEpsilon(input, 5))
}

func TestSolveWithPromptInput(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")

	gamma := SolveGamma(input)
	epsilon := SolveEpsilon(input, 5)

	assert.Equal(t, uint16(22), gamma)
	assert.Equal(t, uint16(9), epsilon)

	assert.True(t, gamma&epsilon == 0)

	assert.Equal(t, uint16(198), gamma*epsilon)
}

func TestSolveTest(t *testing.T) {
	input := ReadFileToUint16s("test-input")

	gamma := SolveGamma(input)
	epsilon := SolveEpsilon(input, 12)

	assert.Equal(t, uint16(0b000111100010), gamma)   // 482
	assert.Equal(t, uint16(0b111000011101), epsilon) // 3613

	assert.True(t, gamma&epsilon == 0)

	assert.Equal(t, 482*3613, int(gamma)*int(epsilon))
}

func TestSolve(t *testing.T) {
	input := ReadFileToUint16s("input")

	gamma := SolveGamma(input)
	epsilon := SolveEpsilon(input, 12)

	assert.Equal(t, uint16(0x3af), gamma)   // 943
	assert.Equal(t, uint16(0xc50), epsilon) // 3152

	assert.True(t, gamma&epsilon == 0)

	assert.Equal(t, 2972336, int(gamma)*int(epsilon)) // 2972336
}

func TestSolveOxygenRatingWithPromptInput(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")

	assert.Equal(t, uint16(23), SolveOxygenRating(input, 5))
}

func TestSolveCO2RatingWithPromptInput(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")

	assert.Equal(t, uint16(10), SolveCO2Rating(input, 5))
}

func TestSolvePart2(t *testing.T) {
	input := ReadFileToUint16s("input")

	oxygen := SolveOxygenRating(input, 12)
	co2 := SolveCO2Rating(input, 12)

	assert.Equal(t, uint16(0x3a3), oxygen) // 943
	assert.Equal(t, uint16(0xe22), co2)    // 3152

	assert.Equal(t, 3368358, int(oxygen)*int(co2)) // 2972336

}

func TestFilterNumbersWithMostCommonBitSet(t *testing.T) {
	input := ReadFileToUint16s("prompt-input")

	ints := []uint16{
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b11100,
		0b10000,
		0b11001,
	}
	assert.Equal(t, ints, FilterNumbersWithMostCommonBitSet(input, uint16(0b10000)))
}
