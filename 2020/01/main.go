package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const DESIRED_SUM = 2020

type empty struct{}

func main() {
	inputs := map[int64]empty{}
	inputBytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(inputBytes)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputs[num] = empty{}
	}

	if scanner.Err() != nil {
		panic(err)
	}

	workingCandidateFound := false
	for num1, _ := range inputs {
		for num2, _ := range inputs {
			candidate := DESIRED_SUM - num1 - num2
			if _, candidateWorks := inputs[candidate]; candidateWorks {
				fmt.Printf("%v + %v + %v = %v\n", num1, candidate, num2, DESIRED_SUM)
				fmt.Printf("%v * %v * %v = %v\n", num1, candidate, num2, num1*candidate*num2)
				workingCandidateFound = true
				break
			}
		}

		if workingCandidateFound {
			break
		}
	}
}
