package main

import (
	"advent-of-code/file"
	"fmt"
	"math"
	"strings"
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

func getPowerConsumption(data []string) (int, error) {
	gammaRate, epsilonRate, err := getGammaAndEpsilonRates(data)
	if err != nil {
		return 0, err
	}
	return mustConvertBinToDec(gammaRate) * mustConvertBinToDec(epsilonRate), nil
}

func getGammaAndEpsilonRates(data []string) (gammaRate string, epsilonRate string, err error) {
	onesCounter, err := countOnesInAllPositions(data)
	if err != nil {
		return "", "", err
	}

	var (
		gammaRateSB   strings.Builder
		epsilonRateSB strings.Builder
	)
	for i := 0; i < len(onesCounter); i++ {
		if onesCounter[i] > len(data)/2 {
			gammaRateSB.WriteRune('1')
			epsilonRateSB.WriteRune('0')
			continue
		}
		gammaRateSB.WriteRune('0')
		epsilonRateSB.WriteRune('1')
	}
	return gammaRateSB.String(), epsilonRateSB.String(), nil
}

func countOnesInAllPositions(data []string) (map[int]int, error) {
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

// getRatingRec for both Co2ScrubberRating && OxygenGeneratorRating recursively
func getRatingRec(data []string, startPosition int, findCommonBitFunc func(data []string, position int) (rune, error)) (string, error) {
	if len(data) == 1 {
		return data[0], nil
	}

	mostCommon, err := findCommonBitFunc(data, startPosition)
	if err != nil {
		return "", err
	}

	var filtered []string
	for _, next := range data {
		runes := []rune(next)
		if runes[startPosition] == mostCommon {
			filtered = append(filtered, next)
		}
	}
	return getRatingRec(filtered, startPosition+1, findCommonBitFunc)
}

func findMostCommonBit(data []string, position int) (rune, error) {
	onesCount, err := countOnesInPosition(data, position)
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
	onesCount, err := countOnesInPosition(data, position)
	if err != nil {
		return '-', err
	}
	if onesCount*2 < len(data) {
		return '1', nil
	} else {
		return '0', nil
	}
}
func countOnesInPosition(data []string, position int) (int, error) {
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
