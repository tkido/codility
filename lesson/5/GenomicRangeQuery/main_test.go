package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		S    string
		P    []int
		Q    []int
		Want []int
	}{
		{
			"CAGCCTA",
			[]int{2, 5, 0},
			[]int{4, 5, 6},
			[]int{2, 4, 1},
		},
		{
			"AC",
			[]int{0, 0, 1},
			[]int{0, 1, 1},
			[]int{1, 1, 2},
		},
	}
	for _, c := range cases {
		got := Solution(c.S, c.P, c.Q)
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
