package main

import (
	"fmt"
	"strconv"
	"strings"
)

const boardSize = 5

type Board struct {
	values [boardSize][boardSize]int
	marked [boardSize][boardSize]bool
}

func newBoard(rows []string) (*Board, error) {
	b := Board{}
	for j := 0; j < boardSize; j++ {
		err := b.appendRow(j, rows[j])
		if err != nil {
			return nil, err
		}
	}
	return &b, nil
}

func (b *Board) appendRow(row int, str string) error {
	nums := strings.FieldsFunc(str, func(c rune) bool {
		return c == ' '
	})
	if len(nums) != boardSize {
		return fmt.Errorf("incorrect amount of numbers in row")
	}
	for i := 0; i < boardSize; i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return err
		}
		b.values[row][i] = num
	}
	return nil
}

func (b *Board) countStepsToWinLessThanLimit(numbers []int, maxStepsLimit int) (winBeforeLimit bool, steps int) {
	for k := 0; k < maxStepsLimit; k++ {
		marked, i, j := b.mark(numbers[k])
		if !marked {
			continue
		}
		if b.isRowCompleted(i) || b.isColumnCompleted(j) {
			return true, k
		}
	}
	return false, 0
}

func (b *Board) countStepsToWinMoreThanLimit(numbers []int, minStepsLimit int) (winAfterLimit bool, steps int) {
	for k := 0; k < len(numbers); k++ {
		marked, i, j := b.mark(numbers[k])
		if !marked {
			continue
		}
		if b.isRowCompleted(i) || b.isColumnCompleted(j) {
			return k > minStepsLimit, k
		}
	}
	return true, len(numbers)
}

func (b *Board) mark(num int) (marked bool, i, j int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.values[i][j] == num {
				b.marked[i][j] = true
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func (b *Board) isColumnCompleted(col int) bool {
	for i := 0; i < boardSize; i++ {
		if !b.marked[i][col] {
			return false
		}
	}
	return true
}

func (b *Board) isRowCompleted(row int) bool {
	for i := 0; i < boardSize; i++ {
		if !b.marked[row][i] {
			return false
		}
	}
	return true
}

func (b *Board) getSumOfUnmarked() int {
	sumOfUnmarked := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b.marked[i][j] {
				sumOfUnmarked += b.values[i][j]
			}
		}
	}
	return sumOfUnmarked
}
