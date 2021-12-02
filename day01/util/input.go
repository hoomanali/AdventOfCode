package util

import (
	"bufio"
	"os"
)

// ReadInputs reads the input file and appends each line to a slice of strings
func ReadInputs(fileName string) (inputs []string) {
	file, err := os.Open(fileName)
	Check(&err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return
}
