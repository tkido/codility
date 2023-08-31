package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		A    []int
		Want int
	}{
		{[]int{10, 2, 5, 1, 5, 20}, 1},
		{[]int{10, 50, 5, 1}, 0},
		{[]int{}, 0},
		{[]int{-1}, 0},
		{[]int{-1, -2}, 0},
		{[]int{-1, -2, -3}, 0},
		{[]int{10, 50, 5, 100}, 0},
		{[]int{1147483647, 1147483647, 2147483647}, 1},
	}
	for _, c := range cases {
		got := Solution(c.A)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
