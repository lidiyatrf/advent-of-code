package main

import (
	"fmt"
	"strconv"
	"strings"
)

const boardSize = 5

type Board struct {
	board  [boardSize][boardSize]int
	marked [boardSize][boardSize]bool

	stepsToWin int
}

func (b *Board) SetRow(rowNum int, str string) error {
	nums := strings.FieldsFunc(str, splitRow)
	if len(nums) != boardSize {
		return fmt.Errorf("incorrect amount of numbers in row")
	}
	for i := 0; i < boardSize; i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return err
		}
		b.board[rowNum][i] = num
	}
	return nil
}

func (b *Board) Mark(num int) (marked bool, i, j int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.board[i][j] == num {
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
				sumOfUnmarked += b.board[i][j]
			}
		}
	}
	return sumOfUnmarked
}

func splitRow(c rune) bool {
	return c == ' '
}

func parseBoards(data []string) ([]Board, error) {
	result := []Board{}

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			continue
		}
		nextBoard := Board{}
		for j := 0; j < boardSize; j, i = j+1, i+1 {
			err := nextBoard.SetRow(j, data[i])
			if err != nil {
				return nil, err
			}
		}
		result = append(result, nextBoard)
	}

	return result, nil
}
