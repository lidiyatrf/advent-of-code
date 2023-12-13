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

func Min(number1, number2 int) int {
	if number1 < number2 {
		return number1
	}
	return number2
}

func Abs(value int64) int64 {
	if value >= 0 {
		return value
	}
	return value * -1
}
