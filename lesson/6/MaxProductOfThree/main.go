package main

import "sort"

func Solution(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	size := len(A)
	r1 := A[size-1] * A[size-2] * A[size-3]
	r2 := A[0] * A[1] * A[size-1]
	r := r1
	if r2 > r1 {
		r = r2
	}
	return r
}
