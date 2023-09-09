package main

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		S    string
		Want int
	}{
		{"{[()()]}", 1},
		{"([)()]", 0},
	}
	for _, c := range cases {
		got := Solution(c.S)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
