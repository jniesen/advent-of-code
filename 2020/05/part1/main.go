package main

import (
	"bufio"
	"fmt"
	"os"
)

func DecodeBoardingPass(boardingPass string) (int, int, int) {
	seatRows := make([]int, 128)
	for i := 0; i < len(seatRows); i++ {
		seatRows[i] = i
	}

	seatCols := make([]int, 8)
	for i := 0; i < len(seatCols); i++ {
		seatCols[i] = i
	}

	instructions := []byte(boardingPass)
	rowInstructions := instructions[:7]
	colInstructions := instructions[7:]

	row := narrowDown(seatRows, rowInstructions)[0]
	col := narrowDown(seatCols, colInstructions)[0]
	seat := row*8 + col

	return row, col, seat
}

func narrowDown(a []int, instructions []byte) []int {
	if len(a) == 1 {
		return a
	}

	half := len(a) / 2
	instruction := instructions[:1]

	switch string(instruction) {
	case "F", "L":
		return narrowDown(a[:half], instructions[1:])
	case "B", "R":
		return narrowDown(a[half:], instructions[1:])
	default:
		return a
	}
}

func main() {
	highestSeat := 0

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, _, seat := DecodeBoardingPass(scanner.Text())
		if seat > highestSeat {
			highestSeat = seat
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(highestSeat)
}
