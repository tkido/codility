package main

func Solution(A []int) int {
	// Implement your solution here
	m := make(map[int]int, 1000000)
	for _, n := range A {
		m[n]++
	}
	for k, v := range m {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}
