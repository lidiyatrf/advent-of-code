package main

import (
	"advent-of-code/2021/file"
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

	result, err := getPowerConsumption(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func getPowerConsumption(data []string) (int, error) {
	onesCounter, err := countOnes(data)
	if err != nil {
		return 0, err
	}

	gammaRate, epsilonRate := getRates(onesCounter, len(data))

	gammaRateDec := mustConvertBinToDec(gammaRate)
	epsilonRateDec := mustConvertBinToDec(epsilonRate)

	return gammaRateDec * epsilonRateDec, nil
}

func countOnes(data []string) (map[int]int, error) {
	length := len(data[0])
	onesCounter := make(map[int]int) // position - amount

	for _, nextVal := range data {
		if len(nextVal) != length {
			return nil, fmt.Errorf("different length of data in report")
		}
		for i := 0; i < length; i++ {
			if nextVal[i] == '1' {
				onesCounter[i]++
			} else if nextVal[i] != '0' {
				return nil, fmt.Errorf("not binary format in report %q", nextVal)
			}
		}
	}

	return onesCounter, nil
}

func getRates(onesCounter map[int]int, rows int) (gammaRate string, epsilonRate string) {
	var (
		gammaRateSB   strings.Builder
		epsilonRateSB strings.Builder
	)
	for i := 0; i < len(onesCounter); i++ {
		if onesCounter[i] > rows/2 {
			gammaRateSB.WriteRune('1')
			epsilonRateSB.WriteRune('0')
		} else {
			gammaRateSB.WriteRune('0')
			epsilonRateSB.WriteRune('1')
		}
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
