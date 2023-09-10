package main

func Solution(H []int) int {
	s := []int{} // stack
	count := 0
	for _, h := range H {
		if len(s) > 0 && s[len(s)-1] == h {
			continue
		}
		for len(s) > 0 && s[len(s)-1] > h {
			s = s[:len(s)-1]
		}
		if len(s) == 0 || s[len(s)-1] < h {
			s = append(s, h)
			count++
			continue
		}
	}
	return count
}
