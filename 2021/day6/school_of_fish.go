package main

import (
	"strconv"
	"strings"
)

const newBornTimer = 8
const givenBirthTimer = 6

type School struct {
	fishesPerDay map[int]int
}

func newSchool(str string) (*School, error) {
	fishesPerDay := make(map[int]int)
	for i := 0; i <= newBornTimer; i++ {
		fishesPerDay[i] = 0
	}

	strs := strings.Split(str, ",")
	for _, next := range strs {
		day, err := strconv.Atoi(next)
		if err != nil {
			return nil, err
		}
		fishesPerDay[day]++
	}

	return &School{fishesPerDay: fishesPerDay}, nil
}

func (s *School) grow() {
	prev := 0
	var curr int
	for i := newBornTimer; i >= 0; i-- {
		curr = s.fishesPerDay[i]
		s.fishesPerDay[i] = prev
		prev = curr
	}

	s.fishesPerDay[newBornTimer] = prev
	s.fishesPerDay[givenBirthTimer] += prev
}

func (s *School) getTotal() int {
	total := 0
	for _, v := range s.fishesPerDay {
		total += v
	}
	return total
}
