package common

import (
	"io/ioutil"
	"log"
	"strings"
)

// GetFileInputs reads a file and returns the content in a slice
func GetFileInputs(file string) []string {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	input := strings.Split(string(text), "\n")

	return input
}
