package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// GetFileInts reads a file and returns the content in an int slice
func GetFileInts(file string) []int {
	input := GetFileStrs(file)

	var nums []int
	for _, value := range input {
		nums = append(nums, Atoi(value))
	}

	return nums
}

// GetFileStrs reads a file and returns a string slice of the contents
func GetFileStrs(file string) []string {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	input := strings.Split(string(text), "\n")
	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	return input
}

func Atoi(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Could not convert string to int: %v", err)
	}

	return value
}

func Btoi(binary string) int {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert binary to integer: %v", err)
	}

	return int(output)
}
