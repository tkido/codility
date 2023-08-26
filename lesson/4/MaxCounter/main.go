package main

func Solution(N int, A []int) []int {
	a := make([]int, N+1)
	max := 0
	floor := 0
	for _, v := range A {
		if v == N+1 {
			floor = max
		} else {
			if a[v] < floor {
				a[v] = floor
			}
			a[v]++
			if max < a[v] {
				max = a[v]
			}
		}
	}
	for i, v := range a {
		if v < floor {
			a[i] = floor
		}
	}
	return a[1:]
}
