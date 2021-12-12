package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// GetFileInputs reads a file and returns the content in a slice
func GetFileInputs(file string) []int {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	input := strings.Split(string(text), "\n")

	var nums []int
	for _, value := range input {
		nums = append(nums, atoi(value))
	}

	return nums
}

func atoi(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Could not convert string to int: %v", err)
	}

	return value
}
