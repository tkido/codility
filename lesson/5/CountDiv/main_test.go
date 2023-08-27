package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		A, B, K, Want int
	}{
		{2, 4, 2, 2},
		{6, 11, 2, 3},
		{6, 12, 2, 4},
		{0, 10, 2, 6},
		{0, 5, 1, 6},
		{1, 5, 2, 2},
		{1, 6, 2, 3},
		{1, 2, 3, 0},
		{2, 3, 3, 1},
		{3, 4, 3, 1},
		{4, 5, 3, 0},
		{5, 6, 3, 1},
		{0, 6, 3, 3},
		{1, 7, 3, 2},
		{2, 8, 3, 2},
		{3, 9, 3, 3},
		{1, 2, 2000000000, 0},
		{0, 1, 2000000000, 1},
		{0, 2000000000, 2000000000, 2},
	}
	for _, c := range cases {
		fmt.Println(c)
		got := Solution(c.A, c.B, c.K)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
