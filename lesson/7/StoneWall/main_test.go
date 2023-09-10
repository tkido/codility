package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		H    []int
		Want int
	}{
		{[]int{8, 8, 5, 7, 9, 8, 7, 4, 8}, 7},
	}
	for _, c := range cases {
		got := Solution(c.H)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
