package day8

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Line struct {
	wires  []string
	output []string
}

func ReadFile(fileName string) []Line {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.Fields(string(fileBytes))
	fmt.Println(fileFields)

	lines := []Line{}
	for i := 0; i < len(fileFields); i += 15 {
		wires := []string{
			fileFields[i],
			fileFields[i+1],
			fileFields[i+2],
			fileFields[i+3],
			fileFields[i+4],
			fileFields[i+5],
			fileFields[i+6],
			fileFields[i+7],
			fileFields[i+8],
			fileFields[i+9],
		}
		output := []string{
			fileFields[i+11],
			fileFields[i+12],
			fileFields[i+13],
			fileFields[i+14],
		}
		line := Line{wires, output}
		lines = append(lines, line)
	}
	return lines
}

func GetWiresForOne(wires []string) string {
	for _, wire := range wires {
		if len(wire) == 2 {
			return wire
		}
	}
	log.Fatalf("Could not find one in %v", wires)
	return ""
}

func GetWiresForFour(wires []string) string {
	for _, wire := range wires {
		if len(wire) == 4 {
			return wire
		}
	}
	log.Fatalf("Could not find four in %v", wires)
	return ""
}

func GetWiresForSeven(wires []string) string {
	for _, wire := range wires {
		if len(wire) == 3 {
			return wire
		}
	}
	log.Fatalf("Could not find seven in %v", wires)
	return ""
}

func GetWiresForEight(wires []string) string {
	for _, wire := range wires {
		if len(wire) == 7 {
			return wire
		}
	}
	log.Fatalf("Could not find eight in %v", wires)
	return ""
}

func GetWiresForThree(wires []string) string {
	wiresForOne := GetWiresForOne(wires)
	for _, wire := range wires {
		if len(wire) == 5 && strings.Contains(wire, string(wiresForOne[0])) && strings.Contains(wire, string(wiresForOne[1])) {
			return wire
		}
	}
	log.Fatalf("Could not find three in %v", wires)
	return ""
}

func GetWiresForFive(wires []string) string {
	wiresForOne := GetWiresForOne(wires)
	wiresForFour := GetWiresForFour(wires)

	log.Printf("Wires in one: %s, wires in four: %s", wiresForOne, wiresForFour)

	wiresInFourNotOne := []string{}
	for _, w := range strings.Split(wiresForFour, "") {
		if !strings.Contains(wiresForOne, w) {
			wiresInFourNotOne = append(wiresInFourNotOne, w)
		}
	}

	log.Printf("Wires in four not one: %v", wiresInFourNotOne)

	for _, wire := range wires {
		if len(wire) == 5 && strings.Contains(wire, string(wiresInFourNotOne[0])) && strings.Contains(wire, string(wiresInFourNotOne[1])) {
			log.Printf("Wires in five: %v", wire)
			return wire
		}
	}
	log.Fatalf("Could not find five in %v", wires)
	return ""
}

func GetWiresForTwo(wires []string) string {
	wiresForFive := GetWiresForFive(wires)
	wiresForThree := GetWiresForThree(wires)
	for _, wire := range wires {
		if len(wire) == 5 && wire != wiresForFive && wire != wiresForThree {
			return wire
		}
	}
	log.Fatalf("Could not find two in %v", wires)
	return ""
}

func GetWiresForSix(wires []string) string {
	wiresForOne := GetWiresForOne(wires)
	wiresForFour := GetWiresForFour(wires)
	wiresForNine := GetWiresForNine(wires)

	log.Printf("six: Wires in one: %s, wires in four: %s", wiresForOne, wiresForFour)

	wiresInFourNotOne := []string{}
	for _, w := range strings.Split(wiresForFour, "") {
		if !strings.Contains(wiresForOne, w) {
			wiresInFourNotOne = append(wiresInFourNotOne, w)
		}
	}

	log.Printf("six: Wires in four not one: %v", wiresInFourNotOne)

	for _, wire := range wires {
		if len(wire) == 6 &&
			strings.Contains(wire, string(wiresInFourNotOne[0])) &&
			strings.Contains(wire, string(wiresInFourNotOne[1])) &&
			wire != wiresForNine {
			log.Printf("six: Wires in six: %v", wire)
			return wire
		}
	}
	log.Fatalf("Could not find six in %v", wires)
	return ""
}

func GetWiresForNine(wires []string) string {
	wiresForFour := GetWiresForFour(wires)
	for _, wire := range wires {
		if len(wire) == 6 &&
			strings.Contains(wire, string(wiresForFour[0])) &&
			strings.Contains(wire, string(wiresForFour[1])) &&
			strings.Contains(wire, string(wiresForFour[2])) &&
			strings.Contains(wire, string(wiresForFour[3])) {
			return wire
		}
	}
	log.Fatalf("Could not find nine in %v", wires)
	return ""
}

func GetWiresForZero(wires []string) string {
	wiresForSix := GetWiresForSix(wires)
	wiresForNine := GetWiresForNine(wires)
	for _, wire := range wires {
		if len(wire) == 6 && wire != wiresForSix && wire != wiresForNine {
			return wire
		}
	}
	log.Fatalf("Could not find zero in %v", wires)
	return ""
}

func DecodeLine(line Line) int {
	zero := sortString(GetWiresForZero(line.wires))
	one := sortString(GetWiresForOne(line.wires))
	two := sortString(GetWiresForTwo(line.wires))
	three := sortString(GetWiresForThree(line.wires))
	four := sortString(GetWiresForFour(line.wires))
	five := sortString(GetWiresForFive(line.wires))
	six := sortString(GetWiresForSix(line.wires))
	seven := sortString(GetWiresForSeven(line.wires))
	eight := sortString(GetWiresForEight(line.wires))
	nine := sortString(GetWiresForNine(line.wires))

	result := []string{}

	for _, output := range line.output {
		sortedOutput := sortString(output)
		if sortedOutput == zero {
			result = append(result, "0")
		} else if sortedOutput == one {
			result = append(result, "1")
		} else if sortedOutput == two {
			result = append(result, "2")
		} else if sortedOutput == three {
			result = append(result, "3")
		} else if sortedOutput == four {
			result = append(result, "4")
		} else if sortedOutput == five {
			result = append(result, "5")
		} else if sortedOutput == six {
			result = append(result, "6")
		} else if sortedOutput == seven {
			result = append(result, "7")
		} else if sortedOutput == eight {
			result = append(result, "8")
		} else if sortedOutput == nine {
			result = append(result, "9")
		}
	}
	resultNum, _ := strconv.Atoi(strings.Join(result, ""))
	log.Printf("Decoded wires: %v\n", []string{zero, one, two, three, four, five, six, seven, eight, nine})
	log.Printf("%v = %d\n", line, resultNum)
	return resultNum
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Solve(lines []Line) int {
	sum := 0
	for _, line := range lines {
		for _, output := range line.output {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				sum++
			}
		}
	}
	return sum
}

func SolvePart2(lines []Line) int {
	sum := 0
	for _, line := range lines {
		sum += DecodeLine(line)
	}
	return sum
}
