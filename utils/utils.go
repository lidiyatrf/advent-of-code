package utils

func GetMin(array []int) int {
	if len(array) == 0 {
		panic("unable to get minimum of empty array")
	}
	minimum := array[0]
	for _, num := range array[1:] {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}
