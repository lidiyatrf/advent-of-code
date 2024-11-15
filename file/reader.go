package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ToString(path string) (string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()

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

func ToStrings(path string) ([]string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()

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

func ToInts(path string) ([]int, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()

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
