package main

import (
	"bufio"
	"bytes"
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
	seat := coordsToSeatID(row, col)

	return row, col, seat
}

func coordsToSeatID(row, col int) int {
	return row*8 + col
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

func BuildSeatingGrid(rowCnt, colCnt int) [][]byte {
	cols := make([][]byte, rowCnt)
	for c := range cols {
		rows := make([]byte, colCnt)
		for r := range rows {
			rows[r] = 0
		}
		cols[c] = rows
	}

	return cols
}

func MarkOccupiedSeats(seats [][]byte, row, col int) [][]byte {
	seats[row][col] = 1
	return seats
}

func FindEmptySeats(row []byte, rowID int) (int, int) {
	col := bytes.IndexByte(row, 0)
	return rowID, col
}

func main() {
	seats := BuildSeatingGrid(128, 8)

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row, col, _ := DecodeBoardingPass(scanner.Text())
		seats = MarkOccupiedSeats(seats, row, col)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for i, row := range seats {
		if bytes.Count(row, []byte{0}) == 1 {
			r, c := FindEmptySeats(row, i)
			if c != -1 {
				fmt.Printf(
					"Empty seat found Row: %v, Col: %v, Seat: %v!\n",
					r, c, coordsToSeatID(r, c))
			}
		} else {
			continue
		}
	}
}
