package main

func Solution(X int, A []int) int {
	m := make(map[int]bool, X+1)
	count := X
	for i, v := range A {
		if m[v] {
			continue
		}
		m[v] = true
		count--
		if count == 0 {
			return i
		}
	}
	return -1
}
