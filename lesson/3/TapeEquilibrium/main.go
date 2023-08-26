package main

func Solution(A []int) int {
	right := 0
	for _, v := range A {
		right += v
	}
	left := 0
	min := 100000 * 1000
	for i, v := range A {
		if i == len(A)-1 {
			break
		}
		left += v
		right -= v
		diff := left - right
		if diff < 0 {
			diff = -diff
		}
		if diff < min {
			min = diff
		}
	}
	return min
}
