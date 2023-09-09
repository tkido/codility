package main

func Solution(A []int, B []int) int {
	alive := 0
	s := []int{}
	for i, v := range A {
		if B[i] == 1 {
			s = append(s, v)
		} else {
			for len(s) > 0 {
				if s[len(s)-1] > v {
					break
				}
				s = s[:len(s)-1]
			}
			if len(s) == 0 {
				alive++
			}
		}
	}
	return alive + len(s)
}
