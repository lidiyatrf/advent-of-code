package main

import "strings"

type Polymer struct {
	polymer string
	rules   map[string]string
}

func newPolymer(polymer string, rulesArr []string) *Polymer {
	rules := make(map[string]string)
	for _, ruleStr := range rulesArr {
		splitted := strings.Split(ruleStr, " -> ")
		if len(splitted) != 2 {
			panic("error when splitting")
		}
		rules[splitted[0]] = splitted[1]
	}
	return &Polymer{
		polymer: polymer,
		rules:   rules,
	}
}

func (p *Polymer) polymerize() {
	var result strings.Builder
	result.WriteString(p.polymer[0:1])

	for i := 1; i < len(p.polymer); i++ {
		pair := p.polymer[i-1 : i+1]
		additionalLetter := p.rules[pair]
		result.WriteString(additionalLetter)
		result.WriteString(p.polymer[i : i+1])
	}

	p.polymer = result.String()
}

func (p *Polymer) getMostCommonBitOccurrence() int {
	maxAmount := 0
	counter := make(map[rune]int)
	for _, next := range p.polymer {
		counter[next]++
		if counter[next] > maxAmount {
			maxAmount = counter[next]
		}
	}
	return maxAmount
}

func (p *Polymer) getLeastCommonBitOccurrence() int {
	minAmount := 1
	minRune := []rune(p.polymer)[0]

	counter := make(map[rune]int)
	for _, next := range p.polymer[1:] {
		counter[next]++
		if next == minRune {
			minAmount = counter[next]
		}
		if counter[next] < counter[minRune] {
			minAmount = counter[next]
			minRune = next
		}
	}
	return minAmount
}
