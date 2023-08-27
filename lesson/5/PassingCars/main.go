package main

func Solution(A []int) int {
	limit := 1000000000
	passing := 0
	easts := 0
	for _, v := range A {
		if v == 0 {
			easts++
		} else {
			passing += easts
			if passing > limit {
				return -1
			}
		}
	}
	return passing
}
