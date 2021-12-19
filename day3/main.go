package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var testBits = []uint16{
	0b1,
	0b10,
	0b100,
	0b1000,
	0b10000,
	0b100000,
	0b1000000,
	0b10000000,
	0b100000000,
	0b1000000000,
	0b10000000000,
	0b100000000000,
	0b1000000000000,
	0b10000000000000,
	0b100000000000000,
	0b1000000000000000,
}

type compareFloatsFunc func(float64, float64) bool

func ReadFileToUint16s(fileName string) []uint16 {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	intsAsStrings := strings.Fields(string(fileBytes))

	ints := []uint16{}

	for _, intAsString := range intsAsStrings {
		n, err := strconv.ParseInt(intAsString, 2, 16)
		if err != nil {
			log.Fatal("Error converting ", intAsString, " to int", err)
		}
		ints = append(ints, uint16(n))
	}

	return ints
}

func SolveGamma(data []uint16) uint16 {
	gamma := 0

	for _, testBit := range testBits {
		count := getOnesCount(data, testBit)
		fmt.Printf("Count of %d is %d\n", testBit, count)
		if count > len(data)/2 {
			gamma = gamma | int(testBit)
		}
	}
	return uint16(gamma)
}

func SolveEpsilon(data []uint16, maskSize int) uint16 {
	gamma := SolveGamma(data)

	bitMask := int(math.Pow(2, float64(maskSize))) - 1

	return uint16(bitMask) ^ gamma
}

func SolveOxygenRating(data []uint16, maskSize int) uint16 {
	current := data
	for i := maskSize - 1; i > -1; i-- {
		fmt.Println(current)
		current = FilterNumbersWithBitSet(current, testBits[i], func(f1, f2 float64) bool { return f1 >= f2 })
		if len(current) == 1 {
			break
		}
	}

	return uint16(current[0])
}

func SolveCO2Rating(data []uint16, maskSize int) uint16 {
	current := data
	for i := maskSize - 1; i > -1; i-- {
		fmt.Println(current)
		current = FilterNumbersWithBitSet(current, testBits[i], func(f1, f2 float64) bool { return f1 < f2 })
		if len(current) == 1 {
			break
		}
	}

	return uint16(current[0])
}

func FilterNumbersWithBitSet(data []uint16, testBit uint16, compareFunc compareFloatsFunc) []uint16 {
	filtered := []uint16{}

	count := getOnesCount(data, testBit)

	if compareFunc(float64(count), float64(len(data))/float64(2.0)) {
		for i := 0; i < len(data); i++ {
			if data[i]&testBit == testBit {
				filtered = append(filtered, data[i])
			}
		}
	} else {
		for i := 0; i < len(data); i++ {
			if data[i]&testBit != testBit {
				filtered = append(filtered, data[i])
			}
		}
	}

	return filtered
}

func getOnesCount(data []uint16, testBit uint16) int {
	onesCount := 0
	for i := 0; i < len(data); i++ {
		if data[i]&testBit == testBit {
			onesCount++
		}
	}
	return onesCount
}
