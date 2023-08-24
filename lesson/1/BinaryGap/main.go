package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello!")
}

func Solution(N int) int {
	// Implement your solution here
	fmt.Printf("given: %b\n", N)
	r := 0
	l := 0
	flag := false
	for N > 0 {
		fmt.Printf("%b: %v\n", N, flag)
		if N&1 == 1 {
			if flag {
				if l > r {
					r = l
				}
			} else {
				flag = true
			}
			l = 0
		} else {
			if flag {
				l++
			}
		}
		N = N >> 1
	}

	return r
}
