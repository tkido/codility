package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		X, Y, D, Want int
	}{
		{10, 85, 30, 3},
		{1, 1, 1, 0},
		{1, 2, 1, 1},
		{2, 4, 2, 1},
		{2, 5, 2, 2},
		{2, 6, 2, 2},
	}
	for _, c := range cases {
		got := Solution(c.X, c.Y, c.D)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
