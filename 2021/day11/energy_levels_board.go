package main

type Board struct {
	levels [][]int
}

func newBoard(data []string) *Board {
	var levels [][]int
	for _, nextLine := range data {
		var row []int
		for _, nextVal := range nextLine {
			num := int(nextVal - '0')
			if num < 0 && num > 9 {
				panic("cannot parse not number")
			}
			row = append(row, num)
		}
		levels = append(levels, row)
	}

	return &Board{
		levels: levels,
	}
}

func (b *Board) simulate() {
	var has9 bool

	for i := 0; i < len(b.levels); i++ {
		for j := 0; j < len(b.levels); j++ {
			if !has9 && b.levels[i][j] == 9 {
				has9 = true
			}
			b.levels[i][j]++
		}
	}

	for has9 {
		has9 = false
		for i := 0; i < len(b.levels); i++ {
			for j := 0; j < len(b.levels); j++ {
				if b.levels[i][j] <= 9 {
					continue
				}
				b.flash(i, j)
				b.levels[i][j] = 0
				has9 = true
			}
		}
	}
}

func (b *Board) flash(i, j int) {
	b.levels[i][j]++
	if i != 0 {
		if b.levels[i-1][j] != 0 {
			b.levels[i-1][j]++ // up
		}
		if j != 0 && b.levels[i-1][j-1] != 0 {
			b.levels[i-1][j-1]++ // up left
		}
		if j != len(b.levels)-1 && b.levels[i-1][j+1] != 0 {
			b.levels[i-1][j+1]++ // up right
		}
	}
	if i != len(b.levels)-1 {
		if b.levels[i+1][j] != 0 {
			b.levels[i+1][j]++ // down
		}
		if j != 0 && b.levels[i+1][j-1] != 0 {
			b.levels[i+1][j-1]++ // down left
		}
		if j != len(b.levels)-1 && b.levels[i+1][j+1] != 0 {
			b.levels[i+1][j+1]++ // down right
		}
	}
	if j != 0 && b.levels[i][j-1] != 0 {
		b.levels[i][j-1]++ // left
	}
	if j != len(b.levels)-1 && b.levels[i][j+1] != 0 {
		b.levels[i][j+1]++ // right
	}
}

func (b *Board) getJustFlashed() int {
	counter := 0
	for i := 0; i < len(b.levels); i++ {
		for j := 0; j < len(b.levels); j++ {
			if b.levels[i][j] == 0 {
				counter++
			}
		}
	}
	return counter
}

func (b *Board) doAllOctopusesFlash() bool {
	for i := 0; i < len(b.levels); i++ {
		for j := 0; j < len(b.levels); j++ {
			if b.levels[i][j] != 0 {
				return false
			}
		}
	}
	return true

}
