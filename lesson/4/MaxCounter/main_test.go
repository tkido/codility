package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		N    int
		A    []int
		Want []int
	}{
		{5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2}},
	}
	for _, c := range cases {
		got := Solution(c.N, c.A)
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
