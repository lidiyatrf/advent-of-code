package parse

import (
	"strconv"
	"strings"
)

func StringToInts(str string, delimiter string) ([]int, error) {
	var result []int
	strs := strings.Split(str, delimiter)
	for _, next := range strs {
		num, err := strconv.Atoi(next)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
