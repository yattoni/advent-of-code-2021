package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

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
	testBits := []uint16{
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

	// minBits := bits.Len16(gamma)
	bitMask := 0
	for i := 0; i < maskSize; i++ {
		bitMask += int(math.Pow(2, float64(i)))
	}

	return uint16(bitMask) ^ gamma
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
