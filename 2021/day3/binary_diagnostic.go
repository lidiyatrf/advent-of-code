package main

import (
	"fmt"
	"math"
	"strings"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	powerConsumption, err := getPowerConsumption(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("power consumption:", powerConsumption)

	lifeSupportRating, err := getLifeSupportRating(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("life support rating:", lifeSupportRating)
}

func getLifeSupportRating(data []string) (int, error) {
	oxygenGeneratorRating, err := getOxygenGeneratorRating(data)
	if err != nil {
		return 0, err
	}
	co2ScrubberRating, err := getCo2ScrubberRating(data)
	if err != nil {
		return 0, err
	}

	return mustConvertBinToDec(oxygenGeneratorRating) * mustConvertBinToDec(co2ScrubberRating), nil
}

func getOxygenGeneratorRating(data []string) (string, error) {
	return getRatingRec(data, 0, findMostCommonBit)
}

func getCo2ScrubberRating(data []string) (string, error) {
	return getRatingRec(data, 0, findLeastCommonBit)
}

// Co2ScrubberRating && OxygenGeneratorRating
func getRatingRec(data []string, startPosition int, findCommonBitFunc func(data []string, position int) (rune, error)) (string, error) {
	if len(data) == 1 {
		return data[0], nil
	}
	mostCommon, err := findCommonBitFunc(data, startPosition)
	if err != nil {
		return "", err
	}
	filtered := []string{}
	for _, next := range data {
		runes := []rune(next)
		if runes[startPosition] == mostCommon {
			filtered = append(filtered, next)
		}
	}
	return getRatingRec(filtered, startPosition+1, findCommonBitFunc)
}

// findMostCommonBit, position starts from 0
func findMostCommonBit(data []string, position int) (rune, error) {
	onesCount, err := countOnesOnPosition(data, position)
	if err != nil {
		return '-', err
	}
	if onesCount*2 >= len(data) {
		return '1', nil
	} else {
		return '0', nil
	}
}

func findLeastCommonBit(data []string, position int) (rune, error) {
	onesCount, err := countOnesOnPosition(data, position)
	if err != nil {
		return '-', err
	}
	if onesCount*2 < len(data) {
		return '1', nil
	} else {
		return '0', nil
	}
}

func countOnesOnPosition(data []string, position int) (int, error) {
	var onesCount int
	for _, nextVal := range data {
		if nextVal[position] == '1' {
			onesCount++
			continue
		}
		if nextVal[position] != '0' {
			return 0, fmt.Errorf("not binary format in report %q", nextVal)
		}
	}
	return onesCount, nil
}

// getPowerConsumption returns the result for the first part
func getPowerConsumption(data []string) (int, error) {
	onesCounter, err := countOnesInPositions(data)
	if err != nil {
		return 0, err
	}

	gammaRate, epsilonRate := getGammaAndEpsilonRates(onesCounter, len(data))

	gammaRateDec := mustConvertBinToDec(gammaRate)
	epsilonRateDec := mustConvertBinToDec(epsilonRate)

	return gammaRateDec * epsilonRateDec, nil
}

func countOnesInPositions(data []string) (map[int]int, error) {
	length := len(data[0])
	onesCounter := make(map[int]int) // position - amount

	for _, nextVal := range data {
		if len(nextVal) != length {
			return nil, fmt.Errorf("different length of data in report")
		}
		for i := 0; i < length; i++ {
			if nextVal[i] == '1' {
				onesCounter[i]++
				continue
			}
			if nextVal[i] != '0' {
				return nil, fmt.Errorf("not binary format in report %q", nextVal)
			}
		}
	}

	return onesCounter, nil
}

func getGammaAndEpsilonRates(onesCounter map[int]int, rows int) (gammaRate string, epsilonRate string) {
	var (
		gammaRateSB   strings.Builder
		epsilonRateSB strings.Builder
	)
	for i := 0; i < len(onesCounter); i++ {
		if onesCounter[i] > rows/2 {
			gammaRateSB.WriteRune('1')
			epsilonRateSB.WriteRune('0')
			continue
		}
		gammaRateSB.WriteRune('0')
		epsilonRateSB.WriteRune('1')
	}
	return gammaRateSB.String(), epsilonRateSB.String()
}

func mustConvertBinToDec(num string) int {
	result := 0
	for i, val := range num {
		if val == '0' {
			continue
		}
		if val != '1' {
			panic("cannot convert non binary")
		}
		pow := len(num) - i - 1
		result += int(math.Pow(2, float64(pow)))
	}
	return result
}
