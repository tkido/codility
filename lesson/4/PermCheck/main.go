package main

func Solution(A []int) int {
	length := len(A)
	m := make(map[int]bool, length+1)
	count := length
	for _, v := range A {
		if v > length {
			return 0
		}
		if m[v] {
			return 0
		}
		m[v] = true
		count--
		if count == 0 {
			return 1
		}
	}
	return 0
}
