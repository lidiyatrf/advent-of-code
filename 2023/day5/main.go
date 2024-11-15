package main

import (
	"advent-of-code/file"
	"advent-of-code/split"
	"advent-of-code/utils"
	"fmt"
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

func puzzle2(lines []string) int {
	seedsFromFile, _ := split.ToInts(lines[0][7:], " ")
	sourceArray := fromPairsToSourceRanges(seedsFromFile)

	for i := 3; i < len(lines); {
		mappingsLastLine := i
		for ; mappingsLastLine < len(lines); mappingsLastLine++ {
			if lines[mappingsLastLine] == "" {
				break
			}
		}

		sourceArray = mapRangesArray(sourceArray, lines[i:mappingsLastLine])
		i = mappingsLastLine + 2
	}

	return getRangesMinimum(sourceArray)
}

func getRangesMinimum(sourceRange []sourceRange) int {
	minimum := sourceRange[0].start
	for _, source := range sourceRange {
		if source.start < minimum {
			minimum = source.start
		}
	}
	return minimum
}

func fromPairsToSourceRanges(pairs []int) []sourceRange {
	var newArr []sourceRange
	for i := 0; i < len(pairs); i += 2 {
		newArr = append(newArr, sourceRange{pairs[i], pairs[i] + pairs[i+1] - 1})
	}
	return newArr
}

func mapRangesArray(arraySource []sourceRange, mappingLines []string) (res []sourceRange) {
	rules := fromMappingLinesToDestinationRangeMapping(mappingLines)

	for i := 0; i < len(arraySource); i++ {
		source := arraySource[i]
		var updated bool
		for _, rule := range rules {
			upd, new, ok := applyRule(source, rule)
			arraySource = append(arraySource, new...)
			updated = updated || ok
			if ok {
				res = append(res, upd)
				break
			}
		}
		if !updated {
			res = append(res, source)
		}
	}

	return res
}

func applyRule(source sourceRange, rule rule) (sourceRange, []sourceRange, bool) {
	var newSources []sourceRange
	updatedSource := source
	var updated bool

	if source.start >= rule.start && source.end <= rule.end {
		source.start += rule.shift
		source.end += rule.shift
		updatedSource = source
		updated = true
	}
	if source.start < rule.start && source.end > rule.end {
		newSources = []sourceRange{{source.start, rule.start - 1}, {rule.end + 1, source.end}}
		updatedSource = sourceRange{rule.start + rule.shift, rule.end + rule.shift}
		updated = true
	}
	if source.start < rule.start && source.end <= rule.end && source.end >= rule.start {
		newSources = []sourceRange{{source.start, rule.start - 1}}
		updatedSource = sourceRange{rule.start + rule.shift, source.end + rule.shift}
		updated = true
	}
	if source.start >= rule.start && source.end > rule.end && source.start <= rule.end {
		newSources = []sourceRange{{rule.end + 1, source.end}}
		updatedSource = sourceRange{source.start + rule.shift, rule.end + rule.shift}
		updated = true
	}
	return updatedSource, newSources, updated
}

func fromMappingLinesToDestinationRangeMapping(mappingLines []string) []rule {
	var newArr []rule
	for _, line := range mappingLines {
		mapLine, _ := split.ToInts(line, " ")
		newArr = append(newArr, rule{mapLine[1], mapLine[1] + mapLine[2] - 1, mapLine[0] - mapLine[1]})
	}
	return newArr
}

func puzzle1(lines []string) int {
	sourceArray, _ := split.ToInts(lines[0][7:], " ")

	for i := 3; i < len(lines); {
		mappingsLastLine := i
		for ; mappingsLastLine < len(lines); mappingsLastLine++ {
			if lines[mappingsLastLine] == "" {
				break
			}
		}

		sourceArray = mapArray(sourceArray, lines[i:mappingsLastLine])
		i = mappingsLastLine + 2
	}

	return utils.GetMin(sourceArray)
}

func mapArray(arraySource []int, mappingLines []string) []int {
	var newSource []int

	for _, line := range mappingLines {
		mapLine, _ := split.ToInts(line, " ")

		destinationMin := mapLine[1]
		destinationMax := mapLine[1] + mapLine[2]

		for _, value := range arraySource {
			if value > destinationMin && value < destinationMax {
				shift := value - destinationMin
				newSource = append(newSource, mapLine[0]+shift)
			}
		}
	}

	return newSource
}

type sourceRange struct {
	start int
	end   int
}

type rule struct {
	start int
	end   int
	shift int
}
