package util

import (
	"fmt"
	"os"
)

// Check prints an error and exits the program if an error has occurred
func Check(e *error) {
	if e != nil {
		fmt.Printf("\nepic fail, %v", e)
		os.Exit(1)
	}
}
