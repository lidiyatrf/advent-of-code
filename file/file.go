package file

import (
	"fmt"
	"log"
	"os"

	"advent-of-code/file/reader"
)

func ToString(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()
	return reader.ToString(file)
}

func ToStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()
	return reader.ToStrings(file)
}

func ToInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("unable to close input file: %v", err)
		}
	}()
	return reader.ToInts(file)
}
