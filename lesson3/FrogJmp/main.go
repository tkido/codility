package main

func Solution(X int, Y int, D int) int {
	if X == Y {
		return 0
	}
	r := (Y - X) / D
	if (Y-X)%D != 0 {
		r++
	}
	return r
}
