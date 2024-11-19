package reader

import (
	"bufio"
	"io"
	"strconv"
)

func ToString(reader io.Reader) (string, error) {
	var result string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		result += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}

func ToStrings(reader io.Reader) ([]string, error) {
	result := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return result, nil
}

func ToInts(reader io.Reader) ([]int, error) {
	result := []int{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []int{}, err
		}
		result = append(result, num)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	return result, nil
}
