package main

import (
	"sort"
	"strings"
)

var correctDisplay = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type Display struct {
	currLetterToCorrect  map[rune]rune
	digitToCurrentSignal map[int]string
}

func newDisplay(signalWires []string) *Display {
	currLetterToCorrect := make(map[rune]rune)
	digitToCurrentSignal := make(map[int]string)

	d := &Display{
		currLetterToCorrect:  currLetterToCorrect,
		digitToCurrentSignal: digitToCurrentSignal,
	}

	signalWires = d.populate1478(signalWires)
	d.populateSegmentA()
	signalWires = d.populate6andSegmentsCandF(signalWires)
	signalWires = d.populate5andSegmentE(signalWires)
	signalWires = d.populate9(signalWires)
	signalWires = d.populate0andSegmentD(signalWires) // ??
	d.populateSegmentB()
	signalWires = d.populate3(signalWires)
	signalWires = d.populate2(signalWires)
	d.populateSegmentG()

	return d
}

func (d *Display) populate1478(signalWires []string) []string {
	var leftSignalWires []string
	for _, next := range signalWires {
		switch len(next) {
		case 2:
			d.digitToCurrentSignal[1] = next
		case 4:
			d.digitToCurrentSignal[4] = next
		case 3:
			d.digitToCurrentSignal[7] = next
		case 7:
			d.digitToCurrentSignal[8] = next
		default:
			leftSignalWires = append(leftSignalWires, next)
		}
	}
	return leftSignalWires
}

func (d *Display) populateSegmentA() {
	one, seven := d.digitToCurrentSignal[1], d.digitToCurrentSignal[7]
	for _, next := range seven {
		if !strings.ContainsRune(one, next) {
			d.currLetterToCorrect[next] = 'a'
			return
		}
	}
	panic("wrong seven-one relation")
}

func (d *Display) populate6andSegmentsCandF(signalWires []string) []string {
	var leftSignalWires []string

	for _, next := range signalWires {
		if len(next) != 6 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		oneRunes := []rune(d.digitToCurrentSignal[1])
		containsZero := strings.ContainsRune(next, oneRunes[0])
		containsFirst := strings.ContainsRune(next, oneRunes[1])

		if containsZero && containsFirst {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		d.digitToCurrentSignal[6] = next
		if containsZero {
			d.currLetterToCorrect[oneRunes[0]] = 'f'
			d.currLetterToCorrect[oneRunes[1]] = 'c'
		} else {
			d.currLetterToCorrect[oneRunes[0]] = 'c'
			d.currLetterToCorrect[oneRunes[1]] = 'f'
		}
	}
	return leftSignalWires
}

func (d *Display) populate5andSegmentE(signalWires []string) []string {
	var leftSignalWires []string

	for _, next := range signalWires {
		if len(next) != 5 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		var notIncludedRune rune
		for _, r := range d.digitToCurrentSignal[6] {
			if !strings.ContainsRune(next, r) {
				if notIncludedRune == 0 {
					notIncludedRune = r
				} else {
					notIncludedRune = 0
					break
				}
			}
		}
		if notIncludedRune == 0 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		d.digitToCurrentSignal[5] = next
		d.currLetterToCorrect[notIncludedRune] = 'e'
	}
	return leftSignalWires
}

func (d *Display) populate0andSegmentD(signalWires []string) []string {
	var leftSignalWires []string

	for _, next := range signalWires {
		if len(next) != 6 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		var notIncludedRune rune
		for _, r := range d.digitToCurrentSignal[8] {
			if !strings.ContainsRune(next, r) {
				if notIncludedRune == 0 {
					notIncludedRune = r
				} else {
					notIncludedRune = 0
					break
				}
			}
		}
		if notIncludedRune == 0 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		d.digitToCurrentSignal[0] = next
		d.currLetterToCorrect[notIncludedRune] = 'd'
	}
	return leftSignalWires
}

func (d *Display) populate9(signalWires []string) []string {
	var currE rune
	for k, v := range d.currLetterToCorrect {
		if v == 'e' {
			currE = k
			break
		}
	}

	var leftSignalWires []string

	for _, next := range signalWires {
		if len(next) != 6 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		var nineWasDetected bool
		for _, r := range d.digitToCurrentSignal[8] {
			if !strings.ContainsRune(next, r) && r == currE {
				nineWasDetected = true
				break
			}
		}

		if !nineWasDetected {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		d.digitToCurrentSignal[9] = next
	}
	return leftSignalWires
}

func (d *Display) populateSegmentB() {
	var currC rune
	var currF rune
	var currD rune
	for k, v := range d.currLetterToCorrect {
		switch v {
		case 'c':
			currC = k
		case 'f':
			currF = k
		case 'd':
			currD = k
		}
	}

	for _, next := range d.digitToCurrentSignal[4] {
		if next != currC && next != currF && next != currD {
			d.currLetterToCorrect[next] = 'b'
			return
		}
	}
	panic("cannot populate segment b")
}

func (d *Display) populate3(signalWires []string) []string {
	var leftSignalWires []string

	for _, next := range signalWires {
		var notIncludedRune rune
		for _, r := range d.digitToCurrentSignal[9] {
			if !strings.ContainsRune(next, r) {
				if notIncludedRune == 0 {
					notIncludedRune = r
				} else {
					notIncludedRune = 0
					break
				}
			}
		}
		if notIncludedRune == 0 {
			leftSignalWires = append(leftSignalWires, next)
			continue
		}

		d.digitToCurrentSignal[3] = next
	}
	return leftSignalWires
}

func (d *Display) populate2(signalWires []string) []string {
	d.digitToCurrentSignal[2] = signalWires[0]
	return nil
}

func (d *Display) populateSegmentG() {
	var currA rune
	var currC rune
	var currD rune
	var currE rune
	for k, v := range d.currLetterToCorrect {
		switch v {
		case 'a':
			currA = k
		case 'c':
			currC = k
		case 'd':
			currD = k
		case 'e':
			currE = k
		}
	}

	for _, next := range d.digitToCurrentSignal[2] {
		if next != currA && next != currC && next != currD && next != currE {
			d.currLetterToCorrect[next] = 'g'
			return
		}
	}
	panic("cannot populate segment b")

}

func (d *Display) decodeOutputValue(signals []string) int {
	result := 0
	for _, signal := range signals {
		correctName := d.decodeSignal(signal)
		digit := correctDisplay[correctName]
		result = result*10 + digit
	}
	return result
}

func (d *Display) decodeSignal(signal string) string {
	var result []rune
	for _, currChar := range signal {
		correctChar := d.currLetterToCorrect[currChar]
		result = append(result, correctChar)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return string(result)
}
