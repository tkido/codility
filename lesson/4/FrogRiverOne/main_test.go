package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		X    int
		A    []int
		Want int
	}{
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
	}
	for _, c := range cases {
		got := Solution(c.X, c.A)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
