package main

import (
	"fmt"
	"log"

	"advent-of-code/file"
)

func main() {
	data, err := file.ToString("./input.txt")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}

	fmt.Println("puzzle 1:", puzzle(data, 4))
	fmt.Println("puzzle 2:", puzzle(data, 14))
}

func puzzle(line string, distictSymbols int) int {
	symbolsMap := make(map[rune]int, distictSymbols)
	currentStartPosition := 0
	for i, symbol := range line {
		if prevSymbolPosition, contains := symbolsMap[symbol]; contains {
			currentStartPosition = prevSymbolPosition + 1
			symbolsMap[symbol] = i

			for key, value := range symbolsMap {
				if value < prevSymbolPosition {
					delete(symbolsMap, key)
				}
			}
			continue
		}

		symbolsMap[symbol] = i

		if len(symbolsMap) == distictSymbols {
			return currentStartPosition + distictSymbols
		} // mjqjp qmgbljsphdztnv jfqwrcgsmlb
	}

	return -1
}
