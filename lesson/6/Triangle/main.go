package main

import "sort"

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	for i := 0; i <= len(A)-3; i++ {
		if A[i] <= 0 {
			continue
		}
		if int64(A[i])+int64(A[i+1]) > int64(A[i+2]) {
			return 1
		}
	}
	return 0
}
