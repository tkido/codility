package main

import (
	"sort"
)

const limit = 10000000

func Solution(A []int) int {
	count := 0
	open := 0
	lefts := make([]int, len(A))
	rights := make([]int, len(A))
	for i, v := range A {
		lefts[i] = i - v
		rights[i] = i + v
	}
	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})
	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})
	for len(rights) > 0 {
		if len(lefts) > 0 && lefts[0] <= rights[0] {
			open++
			lefts = lefts[1:]
		} else {
			open--
			rights = rights[1:]
			count += open
		}
		if count > limit {
			return -1
		}
	}
	return count
}
