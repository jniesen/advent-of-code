package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(partOne(inputBytes))
	fmt.Println(partTwo(inputBytes))
}

func partOne(input []byte) int {
	validCount := validateEach(input, func(s string) bool {
		validation := strings.Split(s, " ")
		min, _ := strconv.ParseInt(strings.Split(validation[0], "-")[0], 10, 64)
		max, _ := strconv.ParseInt(strings.Split(validation[0], "-")[1], 10, 64)
		requiredChar := strings.TrimSuffix(validation[1], ":")
		password := validation[2]
		charCount := strings.Count(password, requiredChar)

		return (charCount >= int(min) && charCount <= int(max))
	})

	return validCount
}

func partTwo(input []byte) int {
	validCount := validateEach(input, func(s string) bool {
		validation := strings.Split(s, " ")
		pos1, _ := strconv.ParseInt(strings.Split(validation[0], "-")[0], 10, 64)
		pos2, _ := strconv.ParseInt(strings.Split(validation[0], "-")[1], 10, 64)
		requiredChar := strings.TrimSuffix(validation[1], ":")
		password := validation[2]
		splitPassword := strings.Split(password, "")

		firstPos := splitPassword[pos1-1] == requiredChar && splitPassword[pos2-1] != requiredChar
		secondPos := splitPassword[pos1-1] != requiredChar && splitPassword[pos2-1] == requiredChar

		return (firstPos || secondPos)
	})

	return validCount
}

func validateEach(input []byte, validateFunc func(string) bool) int {
	validPasswords := 0
	reader := bytes.NewReader(input)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if validateFunc(scanner.Text()) {
			validPasswords++
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return validPasswords
}
