package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestSolution(t *testing.T) {
	cases := []struct {
		given, want int
	}{
		{9, 2},
		{529, 4},
		{20, 1},
		{15, 0},
		{32, 0},
	}
	for _, c := range cases {
		got := Solution(c.given)
		want := c.want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
