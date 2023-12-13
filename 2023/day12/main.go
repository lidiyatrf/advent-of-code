package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2023/day12/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
}

var numbersRegexp = regexp.MustCompile("\\d+")

func puzzle1(arr []string) int {
	var res int

	// #.#.### 1,1,3
	for _, line := range arr {
		splitted := strings.Split(line, " ")

		numbers := numbersRegexp.FindAllString(splitted[1], -1)
		lineRegexp := regexp.MustCompile(buildRegexpString(numbers))

		res += len(line)
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
