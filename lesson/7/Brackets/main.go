package main

func Solution(S string) int {
	stack := make([]rune, 0)
	closeOf := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, r := range S {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return 0
			}
			if closeOf[stack[len(stack)-1]] != r {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return 0
	}
	return 1
}
