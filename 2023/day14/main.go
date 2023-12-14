package main

import (
	"advent-of-code/file"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	data, err := file.ParseToStrings("2023/day14/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(lines []string) int {
	var sum int

	height := len(lines)

	for j := 0; j < len(lines[0]); j++ {
		nextRoundedRockPosition := height
		for i := 0; i < len(lines); i++ {
			if lines[i][j] == 'O' {
				sum += nextRoundedRockPosition
				nextRoundedRockPosition--
				continue
			}
			if lines[i][j] == '#' {
				nextRoundedRockPosition = height - i - 1
				continue
			}
		}
	}

	return sum
}

func puzzle2(lines []string) int {
	fmt.Println(runtime.NumCPU())
	dish := stringsToRune(lines)
	for i := 0; i < 1000000000; i++ {
		//t := time.Now()
		dish = oneCycle(dish)
		//fmt.Println(i, time.Since(t))
	}
	for i := 0; i < len(dish); i++ {
		fmt.Println(string(dish[i]))
	}
	return calcBoard(dish)
}

func calcBoard(dish [][]rune) int {
	var sum int

	for j := 0; j < len(dish[0]); j++ {
		height := len(dish)
		for i := 0; i < len(dish); i++ {
			if dish[i][j] == 'O' {
				sum += height - i
			}
		}
	}

	return sum
}

func stringsToRune(lines []string) [][]rune {
	dish := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		dish[i] = make([]rune, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			dish[i][j] = rune(lines[i][j])
		}
	}
	return dish
}

func runeToStrings(dish [][]rune) []string {
	var lines []string
	for i := 0; i < len(dish); i++ {
		lines = append(lines, string(dish[i]))
	}
	return lines
}

func oneCycle(dish [][]rune) [][]rune {
	return rollEast(rollSouth(rollWest(rollNorth(dish))))
}

func rollNorth(dish [][]rune) [][]rune {
	n := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for k := 0; k < n; k++ {
		go func(k int) {
			for j := k; j < len(dish[0]); j += n {
				nextRoundedRockPosition := 0
				for i := 0; i < len(dish); i++ {
					if dish[i][j] == 'O' {
						if i != nextRoundedRockPosition {
							dish[i][j] = '.'
							dish[nextRoundedRockPosition][j] = 'O'
						}
						nextRoundedRockPosition++
						continue
					}
					if dish[i][j] == '#' {
						nextRoundedRockPosition = i + 1
						continue
					}
				}
			}
			wg.Done()
		}(k)
	}

	wg.Wait()
	return dish
}

func rollSouth(dish [][]rune) [][]rune {
	height := len(dish) - 1

	n := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for k := 0; k < n; k++ {
		go func(k int) {
			for j := k; j < len(dish[0]); j += n {
				nextRoundedRockPosition := height
				for i := len(dish) - 1; i >= 0; i-- {
					if dish[i][j] == 'O' {
						dish[i][j] = '.'
						dish[nextRoundedRockPosition][j] = 'O'
						nextRoundedRockPosition--
						continue
					}
					if dish[i][j] == '#' {
						nextRoundedRockPosition = i - 1
						continue
					}
				}
			}
			wg.Done()
		}(k)
	}

	wg.Wait()
	return dish
}

func rollWest(dish [][]rune) [][]rune {
	n := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for k := 0; k < n; k++ {
		go func(k int) {
			for i := k; i < len(dish); i += n {
				nextRoundedRockPosition := 0
				for j := 0; j < len(dish[0]); j++ {
					if dish[i][j] == 'O' {
						if j != nextRoundedRockPosition {
							dish[i][j] = '.'
							dish[i][nextRoundedRockPosition] = 'O'
						}
						nextRoundedRockPosition++
						continue
					}
					if dish[i][j] == '#' {
						nextRoundedRockPosition = j + 1
						continue
					}
				}
			}
			wg.Done()
		}(k)
	}

	wg.Wait()
	return dish
}

func rollEast(dish [][]rune) [][]rune {
	width := len(dish) - 1

	n := runtime.NumCPU()
	wg := sync.WaitGroup{}
	wg.Add(n)
	for k := 0; k < n; k++ {
		go func(k int) {
			for i := k; i < len(dish); i += n {
				nextRoundedRockPosition := width
				for j := len(dish[0]) - 1; j >= 0; j-- {
					if dish[i][j] == 'O' {
						if j != nextRoundedRockPosition {
							dish[i][j] = '.'
							dish[i][nextRoundedRockPosition] = 'O'
						}
						nextRoundedRockPosition--
						continue
					}
					if dish[i][j] == '#' {
						nextRoundedRockPosition = j - 1
						continue
					}
				}
			}
			wg.Done()
		}(k)
	}

	wg.Wait()
	return dish
}
