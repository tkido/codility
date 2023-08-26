package main

func Solution(A []int) int {
	limit := len(A) + 1
	m := make(map[int]bool, limit)
	for _, v := range A {
		m[v] = true
	}
	for i := 1; i <= limit; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 0 // MUST NOT HAPPEN
}
