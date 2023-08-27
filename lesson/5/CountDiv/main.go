package main

func Solution(A int, B int, K int) int {
	r := B/K - A/K
	if A%K == 0 {
		r++
	}
	return r
}
