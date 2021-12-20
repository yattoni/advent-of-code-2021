package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToFish(t *testing.T) {
	fish := ReadFileToFish("prompt-input")
	assert.Equal(t, []int{3, 4, 3, 1, 2}, fish)
}

func TestRunOneDay(t *testing.T) {
	fish := ReadFileToFish("prompt-input")

	day1 := RunOneDay(fish)
	assert.Equal(t, []int{2, 3, 2, 0, 1}, day1)

	day2 := RunOneDay(day1)
	assert.Equal(t, []int{1, 2, 1, 6, 0, 8}, day2)

	day3 := RunOneDay(day2)
	assert.Equal(t, []int{0, 1, 0, 5, 6, 7, 8}, day3)
}

func TestSolvePromptInput(t *testing.T) {
	fish := ReadFileToFish("prompt-input")

	assert.Equal(t, 5934, Solve(fish, 80, RunOneDay, CountFish))
}

func TestSolve(t *testing.T) {
	fish := ReadFileToFish("input")

	assert.Equal(t, 373378, Solve(fish, 80, RunOneDay, CountFish))
}

func TestSolvePart1Efficient(t *testing.T) {
	fish := ReadFileToEfficientFish("input")

	assert.Equal(t, 373378, Solve(fish, 80, RunOneEfficientDay, CountEfficientFish))
}

/**
 * Running part 2 with my part1 algorithm takes >30 seconds and was eating up 20 GB of memory when I killed it
 * On day 164 the slice of fish was 554529914 int's long which is 4436239312 bytes which is 4.4 GB.
 * If it could make it to day 256, the length would be 1682576647495 8 byte ints which would be 13.46 TB
 */
// func TestSolvePart2Iniital(t *testing.T) {
// 	fish := ReadFileToFish("input")

// 	assert.Equal(t, 1682576647495, Solve(fish, 256, RunOneDay, CountFish))
// }

func TestSolvePart2PromptInput(t *testing.T) {
	fish := ReadFileToEfficientFish("prompt-input")

	assert.Equal(t, 26984457539, Solve(fish, 256, RunOneEfficientDay, CountEfficientFish))
}

func TestSolvePart2(t *testing.T) {
	fish := ReadFileToEfficientFish("input")

	assert.Equal(t, 1682576647495, Solve(fish, 256, RunOneEfficientDay, CountEfficientFish))
}

func TestSolvePart2MoreEfficient(t *testing.T) {
	fish := ReadFileToEfficientFish("input")

	assert.Equal(t, 1682576647495, Solve(fish, 256, RunOneMoreEfficientDay, CountEfficientFish))
}
