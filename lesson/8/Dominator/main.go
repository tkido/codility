package main

func Solution(A []int) int {
	m := make(map[int]int)
	count := 0
	for i, v := range A {
		m[v]++
		if m[v] > count {
			count = m[v]
		}
		if count > len(A)/2 {
			return i
		}
	}
	return -1
}
