package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		A, B []int
		Want int
	}{
		{
			[]int{4, 3, 2, 1, 5},
			[]int{0, 1, 0, 0, 0},
			2,
		},
	}
	for _, c := range cases {
		got := Solution(c.A, c.B)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
