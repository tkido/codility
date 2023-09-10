package main

func Solution(S string) int {
	stack := 0
	for _, r := range S {
		if r == '(' {
			stack++
		} else {
			if stack == 0 {
				return 0
			}
			stack--
		}
	}
	if stack != 0 {
		return 0
	}
	return 1
}
