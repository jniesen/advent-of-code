package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var TREE = "#"

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pattern = append(pattern, strings.Split(scanner.Text(), ""))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}
	treesFound := 0
	for _, slope := range slopes {
		treesInSlope := plotTrees(slope[0], slope[1], pattern)
		if treesFound == 0 {
			treesFound = treesInSlope
		} else {
			treesFound = treesFound * treesInSlope
		}
	}

	fmt.Println(treesFound)
}

func plotTrees(xTravel, yTravel int, pattern [][]string) int {
	xPos := 1
	yPos := 1
	treesFound := 0

	mapGrid := make([][]string, len(pattern))
	mapGrid = recordNewGround(mapGrid, pattern)

	for range pattern {
		yIndex := yPos - 1
		xIndex := xPos - 1

		if yIndex >= len(mapGrid) {
			break
		}

		fmt.Printf("xPos = %v; yPos = %v\n", xPos, yPos)
		if xIndex >= len(mapGrid[yIndex]) {
			mapGrid = recordNewGround(mapGrid, pattern)
		}

		fmt.Printf("Location contains a '%v'\n", mapGrid[yIndex][xIndex])
		if mapGrid[yIndex][xIndex] == TREE {
			treesFound++
		}

		xPos = xPos + xTravel
		yPos = yPos + yTravel
	}

	return treesFound
}

func recordNewGround(mapGrid, newGround [][]string) [][]string {
	for i := 0; i < len(newGround); i++ {
		mapGrid[i] = append(mapGrid[i], newGround[i]...)
	}

	return mapGrid
}
