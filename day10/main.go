package day10

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func ReadFileToLines(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	return strings.Fields(string(fileBytes))
}

func isOpenParen(char byte) bool {
	return char == '(' || char == '{' || char == '[' || char == '<'
}

func isClosingParen(char byte) bool {
	return char == ')' || char == '}' || char == ']' || char == '>'
}

func parensMatch(a, b byte) bool {
	if a == '(' {
		return b == ')'
	} else if a == '[' {
		return b == ']'
	} else if a == '{' {
		return b == '}'
	} else if a == '<' {
		return b == '>'
	} else {
		log.Fatal("Invalid open paran ", a)
	}
	return false
}

func getMatchingParen(a byte) byte {
	if a == '(' {
		return ')'
	} else if a == '[' {
		return ']'
	} else if a == '{' {
		return '}'
	} else if a == '<' {
		return '>'
	} else {
		log.Fatal("Invalid open paran ", a)
	}
	return '0'
}

func IsCorrupt(line string) string {
	for i := 0; i < len(line); i++ {
		if isOpenParen(line[i]) {
			fmt.Printf("%d is %s\n", i, string(line[i]))
			openCount := 0
			for j := i + 1; j < len(line); j++ {
				if isOpenParen(line[j]) {
					fmt.Printf("line[%d] is open %s\n", j, string(line[j]))
					openCount++
				} else if isClosingParen(line[j]) {
					fmt.Printf("line[%d] is closed %s\n", j, string(line[j]))
					if openCount != 0 {
						fmt.Printf("decrementing open count from %d to %d\n", openCount, openCount-1)
						openCount--
					} else {
						fmt.Printf("line[%d] should equal line[%d], %s %s\n", i, j, string(line[i]), string(line[j]))
						if !parensMatch(line[i], line[j]) {
							fmt.Printf("they don't match, returning %s\n", string(line[j]))
							return string(line[j])
						} else {
							break
						}
					}
				}
			}
		}
	}
	return "0"
}

func AutoComplete(line string) string {
	autoCompleteCharacters := ""
	for i := 0; i < len(line); i++ {
		if isOpenParen(line[i]) {
			fmt.Printf("%d is %s\n", i, string(line[i]))
			openCount := 0
			j := i + 1
			for ; j < len(line); j++ {
				if isOpenParen(line[j]) {
					// fmt.Printf("line[%d] is open %s\n", j, string(line[j]))
					openCount++
				} else if isClosingParen(line[j]) {
					// fmt.Printf("line[%d] is closed %s\n", j, string(line[j]))
					if openCount != 0 {
						// fmt.Printf("decrementing open count from %d to %d\n", openCount, openCount-1)
						openCount--
					} else {
						// fmt.Printf("line[%d] should equal line[%d], %s %s\n", i, j, string(line[i]), string(line[j]))
						if !parensMatch(line[i], line[j]) {
							// fmt.Printf("they don't match, returning %s\n", string(line[j]))
							return string(line[j])
						} else {
							break
						}
					}
				}
			}
			if j == len(line) {
				fmt.Printf("%d is %s and reached the end without a match\n", i, string(line[i]))
				autoCompleteCharacters = string(getMatchingParen(line[i])) + autoCompleteCharacters
			}
		}
	}
	return autoCompleteCharacters
}

func ScoreAutoComplete(str string) int {
	score := 0
	for i := 0; i < len(str); i++ {
		score *= 5
		c := str[i]
		if c == ')' {
			score += 1
		} else if c == ']' {
			score += 2
		} else if c == '}' {
			score += 3
		} else if c == '>' {
			score += 4
		}
	}
	return score
}

func Solve(lines []string) int {
	corruptChars := []string{}
	for _, line := range lines {
		corruptChar := IsCorrupt(line)
		if corruptChar != "0" {
			corruptChars = append(corruptChars, corruptChar)
		}
	}
	score := 0

	for _, c := range corruptChars {
		if c == ")" {
			score += 3
		} else if c == "]" {
			score += 57
		} else if c == "}" {
			score += 1197
		} else if c == ">" {
			score += 25137
		} else {
			log.Fatal("Unknown char", c)
		}
	}

	return score
}

func SolvePart2(lines []string) int {
	autoCompleteLines := []string{}
	for _, line := range lines {
		corruptChar := IsCorrupt(line)
		if corruptChar == "0" {
			autoCompleteLines = append(autoCompleteLines, AutoComplete(line))
		}
	}
	scores := []int{}
	for _, autoComplete := range autoCompleteLines {
		scores = append(scores, ScoreAutoComplete(autoComplete))
	}
	fmt.Printf("Scores: %v\n", scores)
	sort.SliceStable(scores, func(i, j int) bool { return scores[i] < scores[j] })
	fmt.Printf("Sorted scores: %v\n", scores)
	return scores[(len(scores) / 2)]
}
