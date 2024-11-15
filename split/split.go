package split

import (
	"strconv"
	"strings"
)

func ToInts(str string, delimiter string) ([]int, error) {
	var result []int
	strs := strings.Split(str, delimiter)
	for _, next := range strs {
		num, err := strconv.Atoi(strings.Trim(next, " "))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}
