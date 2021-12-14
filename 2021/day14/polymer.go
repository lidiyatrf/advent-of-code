package main

import (
	"math"
	"strings"
)

type Polymer struct {
	polymerPairs map[string]int64
	rules        map[string]string
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

	polymerPairs := make(map[string]int64)

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		polymerPairs[pair]++
	}
	polymerPairs[polymer[len(polymer)-1:]]++

	return &Polymer{
		polymerPairs: polymerPairs,
		rules:        rules,
	}
}

func (p *Polymer) polymerize() {
	newPairs := make(map[string]int64)
	for k, v := range p.polymerPairs {
		if len(k) == 1 {
			newPairs[k]++
			continue
		}
		additionalLetter := p.rules[k]
		first := k[0:1] + additionalLetter
		second := additionalLetter + k[1:2]
		newPairs[first] += v
		newPairs[second] += v
	}
	p.polymerPairs = newPairs
}

func (p *Polymer) getCommonBitsOccurrence() (max, min int64) {
	letterMap := make(map[string]int64)
	for k, v := range p.polymerPairs {
		letterMap[k[:1]] += v
	}

	max, min = 0, math.MaxInt
	for _, v := range letterMap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func (p *Polymer) getPolymerLength() int64 {
	counter := int64(0)
	for _, v := range p.polymerPairs {
		counter += v
	}
	return counter
}
