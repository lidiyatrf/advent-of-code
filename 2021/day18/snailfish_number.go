package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var pairRegex = regexp.MustCompile("\\[\\d+,\\d+\\]")

type SnailfishNumber struct {
	value string
}

func (sn *SnailfishNumber) add(newNum string) {
	sn.value = "[" + sn.value + "," + newNum + "]"
}

func (sn *SnailfishNumber) String() string {
	return sn.value
}

func (sn *SnailfishNumber) reduce() {
	reduced := true
	for reduced {
		reduced = sn.explode()
		if reduced {
			continue
		}
		reduced = sn.split()
	}
}

func (sn *SnailfishNumber) explode() bool {
	countOpenBrackets := 0
	beg := -1
	for i, next := range sn.value {
		switch next {
		case '[':
			countOpenBrackets++
		case ']':
			countOpenBrackets--
		}
		if countOpenBrackets == 5 {
			beg = i
			break
		}
	}

	if beg == -1 {
		return false
	}

	comma := strings.Index(sn.value[beg:], ",") + beg
	end := strings.Index(sn.value[beg:], "]") + beg

	left, err := strconv.Atoi(sn.value[beg+1 : comma])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(sn.value[comma+1 : end])
	if err != nil {
		panic(err)
	}

	// update right
	for i := end; i < len(sn.value); i++ {
		if unicode.IsDigit(rune(sn.value[i])) {
			numLength := 0
			for _, digit := range sn.value[i:] {
				if !unicode.IsDigit(digit) {
					break
				}
				numLength++
			}

			num, err := strconv.Atoi(sn.value[i : i+numLength])
			if err != nil {
				panic(err)
			}
			sum := num + right
			sn.value = sn.value[:i] + strconv.Itoa(sum) + sn.value[i+numLength:]
			break
		}
	}

	sn.value = sn.value[:beg] + "0" + sn.value[end+1:]

	// update left
	for i := beg - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(sn.value[i])) {
			numLength := 1
			for j := i - 1; j >= 0; j-- {
				if !unicode.IsDigit(rune(sn.value[j])) {
					break
				}
				numLength++
				i--
			}

			num, err := strconv.Atoi(sn.value[i : i+numLength])
			if err != nil {
				panic(err)
			}
			sum := num + left
			sn.value = sn.value[:i] + strconv.Itoa(sum) + sn.value[i+numLength:]
			break
		}
	}

	return true
}

func (sn *SnailfishNumber) split() bool {
	var prev = rune(sn.value[0])
	for i, curr := range sn.value[1:] {
		if unicode.IsDigit(curr) && unicode.IsDigit(prev) {
			number, err := strconv.Atoi(sn.value[i : i+2]) // 15
			if err != nil {
				panic(err)
			}
			first := number / 2
			second := number - first

			sn.value = sn.value[:i] + "[" + strconv.Itoa(first) + "," + strconv.Itoa(second) + "]" + sn.value[i+2:]
			return true
		}
		prev = curr
	}

	return false
}

func (sn *SnailfishNumber) getMagnitude() int {
	allPairs := pairRegex.FindAllStringIndex(sn.value, -1)

	for i := len(allPairs) - 1; i >= 0; i-- {
		pair := allPairs[i]

		beg := pair[0]
		end := pair[1]

		comma := strings.Index(sn.value[beg:end], ",") + beg

		left, err := strconv.Atoi(sn.value[beg+1 : comma])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(sn.value[comma+1 : end-1])
		if err != nil {
			panic(err)
		}

		sum := 3*left + 2*right

		ending := ""
		if end != len(sn.value) {
			ending = sn.value[end:]
		}
		sn.value = sn.value[:beg] + strconv.Itoa(sum) + ending
	}

	if strings.Contains(sn.value, "[") {
		return sn.getMagnitude()
	}

	result, err := strconv.Atoi(sn.value)
	if err != nil {
		panic(err)
	}
	return result
}
