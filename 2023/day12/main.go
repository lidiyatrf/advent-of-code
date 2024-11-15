package main

import (
	"advent-of-code/file"
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := file.ToStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

var numbersRegexp = regexp.MustCompile("\\d+")
var questionRegexp = regexp.MustCompile("\\?")
var sharpRegexp = regexp.MustCompile("#")

func puzzle1(arr []string) int {
	var res int

	for _, line := range arr {
		splitted := strings.Split(line, " ")

		numbers := numbersRegexp.FindAllString(splitted[1], -1)

		questionsCount := len(questionRegexp.FindAllString(splitted[0], -1))
		sharpsCount := len(sharpRegexp.FindAllString(splitted[0], -1))

		pattern := regexp.MustCompile(buildRegexpString(numbers))

		var sum int
		for _, num := range numbers {
			v, _ := strconv.Atoi(num)
			sum += v
		}

		guessBitsCount := sum - sharpsCount
		variants := generateVariants(questionsCount, guessBitsCount)

		var count int
		for _, variant := range variants {
			stringVariant := setupVariant(splitted[0], variant, questionsCount)
			if pattern.MatchString(stringVariant) {
				count++
			}
		}

		res += count
	}

	return res
}

var splitDotsRegexp = regexp.MustCompile("[.]{2,}")

func puzzle2(arr []string) int {
	var res int

	for _, line := range arr {
		fmt.Println(line)
		splitted := strings.Split(line, " ")

		patternStr := splitted[0]
		numbersStr := splitted[1]
		for i := 0; i < 4; i++ {
			patternStr += "?" + splitted[0]
			numbersStr += "," + splitted[1]
		}

		numbers := numbersRegexp.FindAllString(numbersStr, -1)

		count, ok := continueParse(patternStr, 0, numbers, make(map[int]map[string]int))
		if !ok {
			panic(fmt.Sprintf("nothing is possible: %s %s", patternStr, numbersStr))
		}

		res += count
	}

	return res
}

func buildRegexpString(numbers []string) string {
	var s strings.Builder // .*[#]{3}.+[#]{2}.+[#]{1}.*
	s.WriteString(".*")
	for i, number := range numbers {
		s.WriteString("[#]{")
		s.WriteString(number)
		s.WriteString("}")
		if i != len(numbers)-1 {
			s.WriteString(".+")
		}
	}
	s.WriteString(".*")
	return s.String()
}

func generateVariants(questionsCount, guessBitsCount int) []uint {
	var min, max uint
	for i := 0; i < guessBitsCount; i++ {
		min += uint(math.Pow(2, float64(i)))
		max += uint(math.Pow(2, float64(questionsCount-i-1)))
	}

	var res []uint
	for i := min; i <= max; i++ {
		if bits.OnesCount(i) != guessBitsCount {
			continue
		}

		res = append(res, i)
	}
	return res
}

func setupVariant(stringPattern string, variant uint, count int) string {
	for i := 0; i < count; i++ {
		var val = "#"
		if variant&(1<<i) == 0 {
			val = "."
		}
		stringPattern = strings.Replace(stringPattern, "?", val, 1)
	}
	return stringPattern
}

func countByPart(stringPattern string, p *regexp.Regexp, guessBitsCount int) int {
	questionsCount := len(questionRegexp.FindAllString(stringPattern, -1))
	t := time.Now()
	variants := generateVariants(questionsCount, guessBitsCount)
	fmt.Printf("Variants %d took %s\n", len(variants), time.Since(t))
	t = time.Now()
	var count int
	for _, variant := range variants {
		stringVariant := setupVariant(stringPattern, variant, questionsCount)
		if p.MatchString(stringVariant) {
			count++
		}
	}
	fmt.Println("Count took", time.Since(t))

	return count
}

func generateStringOf(val string, count int) string {
	if count == 0 {
		return ""
	}
	sb := strings.Builder{}
	for i := 0; i < count; i++ {
		sb.WriteString(val)
	}
	return sb.String()
}

func generatePartialStringOf(count, from, to int) string {
	s := generateStringOf(".", from)
	s += generateStringOf("#", to-from+1)
	s += generateStringOf(".", count-from+1)
	return s
}

func continueParse(str string, numIdx int, numbers []string, cache map[int]map[string]int) (int, bool) {
	str = strings.TrimLeft(str, ".")
	if len(str) == 0 {
		return 1, numIdx == len(numbers)
	}

	if numIdx == len(numbers) {
		sharpIndex := strings.Index(str, "#")
		if sharpIndex >= 0 {
			return 0, false
		}
		return 1, true
	}

	solutions, found := cache[numIdx][str]
	if found {
		return solutions, true
	}

	val, _ := strconv.Atoi(numbers[numIdx])

	if val > len(str) {
		return 0, false
	}

	for i := 0; i < val; i++ {
		if str[i] == '.' {
			if str[0] == '?' {
				return continueParse(str[1:], numIdx, numbers, cache)
			} else {
				return 0, false
			}
		}
	}
	if val != len(str) && str[val] == '#' {
		if str[0] == '?' {
			return continueParse(str[1:], numIdx, numbers, cache)
		} else {
			return 0, false
		}
	}
	if val == len(str) {
		return 1, numIdx+1 == len(numbers)
	}
	count := 0
	num, ok := continueParse(str[val+1:], numIdx+1, numbers, cache)
	if ok {
		count += num
	}
	if str[0] == '?' {
		num, ok := continueParse(str[1:], numIdx, numbers, cache)
		if ok {
			count += num
		}
	}
	_, ok = cache[numIdx]
	if !ok {
		cache[numIdx] = make(map[string]int)
	}
	cache[numIdx][str] = count
	return count, true
}
