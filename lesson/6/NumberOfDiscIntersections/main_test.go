package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		A    []int
		Want int
	}{
		{[]int{1, 5, 2, 1, 4, 0}, 11},
	}
	for _, c := range cases {
		got := Solution(c.A)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
