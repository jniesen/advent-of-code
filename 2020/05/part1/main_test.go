package main

import (
	"testing"
)

func TestDecodeBoardingPass(t *testing.T) {
	cases := []struct {
		BoardingPass   string
		Row, Col, Seat int
	}{
		{BoardingPass: "BFFFBBFRRR", Row: 70, Col: 7, Seat: 567},
		{BoardingPass: "FFFBBBFRRR", Row: 14, Col: 7, Seat: 119},
		{BoardingPass: "BBFFBBFRLL", Row: 102, Col: 4, Seat: 820},
	}

	for _, tc := range cases {
		t.Run(tc.BoardingPass, func(t *testing.T) {
			row, col, seat := DecodeBoardingPass(tc.BoardingPass)
			if row != tc.Row {
				t.Errorf("Expected Row %v, Got Row %v", tc.Row, row)
			}

			if col != tc.Col {
				t.Errorf("Expected Col %v, Got Col %v", tc.Col, col)
			}

			if seat != tc.Seat {
				t.Errorf("Expected Seat %v, Got Seat %v", tc.Seat, seat)
			}
		})
	}
}
