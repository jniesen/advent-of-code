package main

import (
	"bytes"
	"testing"
)

func TestBuildSeatingGrid(t *testing.T) {
	cases := []struct {
		Desc       string
		Rows, Cols int
	}{
		{Desc: "3 x 2 filled with 0", Rows: 3, Cols: 3},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			grid := BuildSeatingGrid(tc.Rows, tc.Cols)
			if len(grid) != tc.Rows {
				t.Errorf("Expected %v Rows, Got %v", tc.Rows, len(grid))
			}

			for i, row := range grid {
				if len(row) != tc.Cols {
					t.Errorf("Row %v, Expected %v Columns, Got %v", i, tc.Cols, len(row))
				}
			}

			grid[2][1] = 1
			if grid[0][1] != 0 {
				t.Errorf("Underlying byte slice is shared!")
			}
		})
	}
}

func TestMarkOccupiedSeats(t *testing.T) {
	cases := []struct {
		Desc            string
		Seats, Expected [][]byte
		Occupied        [][]int
	}{
		{
			Desc: "All Occupied",
			Seats: [][]byte{
				[]byte{0, 0},
				[]byte{0, 0},
			},
			Occupied: [][]int{
				[]int{0, 0}, []int{0, 1},
				[]int{1, 0}, []int{1, 1},
			},
			Expected: [][]byte{
				[]byte{1, 1},
				[]byte{1, 1},
			},
		},
		{
			Desc: "One Occupied",
			Seats: [][]byte{
				[]byte{0, 0},
				[]byte{0, 0},
			},
			Occupied: [][]int{
				[]int{0, 0},
			},
			Expected: [][]byte{
				[]byte{1, 0},
				[]byte{0, 0},
			},
		},
		{
			Desc: "Two Occupied",
			Seats: [][]byte{
				[]byte{0, 0},
				[]byte{0, 0},
			},
			Occupied: [][]int{
				[]int{0, 0},
				[]int{1, 1},
			},
			Expected: [][]byte{
				[]byte{1, 0},
				[]byte{0, 1},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			seats := tc.Seats
			for _, seat := range tc.Occupied {
				seats = MarkOccupiedSeats(tc.Seats, seat[0], seat[1])
			}

			for i := range seats {
				if bytes.Compare(seats[i], tc.Expected[i]) != 0 {
					t.Errorf("\nExpected:\n\t%v\nGot\n\t%v", tc.Expected[i], seats[i])
				}
			}
		})
	}
}

func TestFindEmptySeats(t *testing.T) {
	cases := []struct {
		Desc            string
		Row             []byte
		RowID, Expected int
	}{
		{
			Desc:     "All Full",
			Row:      []byte{1, 1},
			RowID:    1,
			Expected: -1,
		},
		{
			Desc:     "One Available",
			Row:      []byte{0, 1},
			RowID:    1,
			Expected: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			_, col := FindEmptySeats(tc.Row, tc.RowID)

			if tc.Expected != col {
				t.Errorf("Expected %v, Got %v", tc.Expected, col)
			}
		})
	}
}
