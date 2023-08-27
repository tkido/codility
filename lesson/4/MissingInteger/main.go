package main

func Solution(A []int) int {
	a := make([]int, 1000000)
	for _, v := range A {
		if v <= 0 {
			continue
		}
		a[v-1]++
	}
	for i, v := range a {
		if v == 0 {
			return i + 1
		}
	}
	return -1 // MUST NOT HAPPEN
}
