package file

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}

func ParseToStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return result, nil
}

func ParseToInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	result := []int{}
	scanner := bufio.NewScanner(file)
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
