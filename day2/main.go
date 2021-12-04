package day2

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type Instruction struct {
	direction Direction
	distance  int
}

func ReadFileToInstructions(fileName string) []Instruction {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.Fields(string(fileBytes))

	instructions := []Instruction{}

	for i := 0; i < len(fileFields); i += 2 {
		direction := directionFromString(fileFields[i])

		distance, err := strconv.Atoi(fileFields[i+1])
		if err != nil {
			log.Fatal("Error converting ", fileFields[i+1], " to int", err)
		}
		instructions = append(instructions, Instruction{direction, distance})
	}

	return instructions
}

func Solve(instructions []Instruction) int {
	horizontalPosition := 0
	depth := 0

	for _, instruction := range instructions {
		if instruction.direction == Forward {
			horizontalPosition += instruction.distance
		} else if instruction.direction == Down {
			depth += instruction.distance
		} else if instruction.direction == Up {
			depth -= instruction.distance
		} else {
			log.Fatal("Unknown direction")
		}
	}

	return horizontalPosition * depth
}

func SolveWithAim(instructions []Instruction) int {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, instruction := range instructions {
		if instruction.direction == Forward {
			horizontalPosition += instruction.distance
			depth += aim * instruction.distance
		} else if instruction.direction == Down {
			aim += instruction.distance
		} else if instruction.direction == Up {
			aim -= instruction.distance
		} else {
			log.Fatal("Unknown direction")
		}
	}

	return horizontalPosition * depth
}

func directionFromString(str string) Direction {
	switch str {
	case "forward":
		return Forward
	case "up":
		return Up
	case "down":
		return Down
	default:
		return -1
	}
}
