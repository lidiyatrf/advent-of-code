package main

import (
	"advent-of-code/file"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2023/day7/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

type handAndBid struct {
	hand        string
	combination int
	bid         int
}

const (
	combinationHighCard     int = iota
	combinationOnePair          // 2+1+1+1
	combinationTwoPair          // 2+2+1
	combinationThreeOfAKind     // 3+1+1
	combinationFullHouse        // 3+2
	combinationFourOfAKind      // 4+1
	combinationFiveOfAKind      // 5
)

var letters1 = map[byte]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

func puzzle1(lines []string) int {
	var hands []handAndBid

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, handAndBid{hand, getCombination1(hand), bid})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].combination == hands[j].combination {
			for k := 0; k < len(hands[i].hand); k++ {
				if hands[i].hand[k] != hands[j].hand[k] {
					return letters1[hands[i].hand[k]] > letters1[hands[j].hand[k]]
				}
			}
		}
		return hands[i].combination < hands[j].combination
	})

	var result int
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return result
}

func getCombination1(hand string) int {
	letters := make(map[rune]int)
	for _, letter := range hand {
		letters[letter]++
	}

	switch len(letters) {
	case 1:
		return combinationFiveOfAKind
	case 4:
		return combinationOnePair
	case 5:
		return combinationHighCard
	case 2:
		for _, v := range letters {
			if v == 4 || v == 1 {
				return combinationFourOfAKind
			} else {
				return combinationFullHouse
			}
		}
		panic("unsupported combination")
	case 3:
		for _, v := range letters {
			if v == 3 {
				return combinationThreeOfAKind
			}
			if v == 2 {
				return combinationTwoPair
			}
		}
		panic("unsupported combination")
	default:
		return -1
	}
}

var letters2 = map[byte]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'T': 3,
	'9': 4,
	'8': 5,
	'7': 6,
	'6': 7,
	'5': 8,
	'4': 9,
	'3': 10,
	'2': 11,
	'J': 12,
}

func puzzle2(lines []string) int {
	var hands []handAndBid

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, handAndBid{hand, getCombination2(hand), bid})
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].combination == hands[j].combination {
			for k := 0; k < len(hands[i].hand); k++ {
				if hands[i].hand[k] != hands[j].hand[k] {
					return letters2[hands[i].hand[k]] > letters2[hands[j].hand[k]]
				}
			}
		}
		return hands[i].combination < hands[j].combination
	})

	var result int
	for i, hand := range hands {
		result += (i + 1) * hand.bid
	}
	return result
}

func getCombination2(hand string) int {
	letters := make(map[rune]int)
	var joker int
	for _, letter := range hand {
		if letter == 'J' {
			joker++
			continue
		}
		letters[letter]++
	}

	if joker != 0 {
		var maxCount int
		var maxLetter rune
		for k, v := range letters {
			if v > maxCount {
				maxCount = v
				maxLetter = k
			}
		}
		letters[maxLetter] += joker
	}

	switch len(letters) {
	case 1:
		return combinationFiveOfAKind
	case 4:
		return combinationOnePair
	case 5:
		return combinationHighCard
	case 2:
		for _, v := range letters {
			if v == 4 || v == 1 {
				return combinationFourOfAKind
			} else {
				return combinationFullHouse
			}
		}
		panic("unsupported combination")
	case 3:
		for _, v := range letters {
			if v == 3 {
				return combinationThreeOfAKind
			}
			if v == 2 {
				return combinationTwoPair
			}
		}
		panic("unsupported combination")
	default:
		return -1
	}
}
