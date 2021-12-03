package util

import (
	"bufio"
	"os"
)

var FileName string = ""

// ReadInputs reads the input file and appends each line to a slice of strings
func ReadInputs() (inputs []string) {
	file, err := os.Open(FileName)
	Handle(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return
}
