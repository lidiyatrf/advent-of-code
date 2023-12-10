package main

import (
	"advent-of-code/file"
	"advent-of-code/parse"
	"advent-of-code/utils"
	"fmt"
	"sort"
)

func main() {
	data, err := file.ParseToStrings("2023/day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(lines []string) int {
	sourceArray, _ := parse.StringToInts(lines[0][7:], " ")

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
		mapLine, _ := parse.StringToInts(line, " ")

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

type destinationRangeMapping struct {
	start int
	end   int
	shift int
}

func puzzle2(lines []string) int {
	seedsFromFile, _ := parse.StringToInts(lines[0][7:], " ")
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

func fromPairsToSourceRanges(pairs []int) []sourceRange {
	var newArr []sourceRange
	for i := 0; i < len(pairs); i += 2 {
		newArr = append(newArr, sourceRange{pairs[i], pairs[i] + pairs[i+1] - 1})
	}
	return newArr
}

func fromMappingLinesToDestinationRangeMapping(mappingLines []string) []destinationRangeMapping {
	var newArr []destinationRangeMapping
	for _, line := range mappingLines {
		mapLine, _ := parse.StringToInts(line, " ")
		newArr = append(newArr, destinationRangeMapping{mapLine[1], mapLine[1] + mapLine[2], mapLine[0] - mapLine[1]})
	}
	return newArr
}

func mapRangesArray(arraySource []sourceRange, mappingLines []string) []sourceRange {
	destinationRangeMappings := fromMappingLinesToDestinationRangeMapping(mappingLines)

	var rangesNoOverlap []sourceRange
	for _, currentSource := range arraySource {
		sourceStart := currentSource.start
		sourceEnd := currentSource.end

		var overlapFound bool

		for _, mapping := range destinationRangeMappings {
			destinationEnd := mapping.end
			destinationStart := mapping.start

			if sourceStart > destinationEnd || sourceEnd < destinationStart {
				continue
			}

			overlapFound = true

			intSlice := []int{sourceStart, sourceEnd, destinationEnd, destinationStart}
			sort.Ints(intSlice)

			rangesNoOverlap = append(rangesNoOverlap, sourceRange{intSlice[1], intSlice[2]})

			if sourceStart > destinationStart && sourceEnd < destinationEnd {
				continue
			}
			if destinationStart < sourceStart && destinationEnd < sourceEnd {
				rangesNoOverlap = append(rangesNoOverlap, sourceRange{intSlice[2] + 1, intSlice[3]})
				continue
			}
			if destinationEnd > sourceEnd && destinationStart > sourceStart {
				rangesNoOverlap = append(rangesNoOverlap, sourceRange{intSlice[0], intSlice[1] - 1})
				continue
			}
			rangesNoOverlap = append(rangesNoOverlap, sourceRange{intSlice[0], intSlice[1] - 1})
			rangesNoOverlap = append(rangesNoOverlap, sourceRange{intSlice[2] + 1, intSlice[3]})
			continue
		}

		if !overlapFound {
			rangesNoOverlap = append(rangesNoOverlap, currentSource)
		}
	}

	var newSource []sourceRange

	for _, currentSource := range rangesNoOverlap {
		sourceStart := currentSource.start
		sourceEnd := currentSource.end

		var overlapFound bool

		for _, mapping := range destinationRangeMappings {
			destinationEnd := mapping.end
			destinationStart := mapping.start

			if sourceStart > destinationEnd || sourceEnd < destinationStart {
				continue
			}

			overlapFound = true
			newSource = append(newSource, sourceRange{sourceStart + mapping.shift, sourceEnd + mapping.shift})
			break
		}

		if !overlapFound {
			newSource = append(newSource, currentSource)
		}
	}
	return newSource
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
