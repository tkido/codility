package main

func Solution(A []int) int {
	leader := 0
	size := 0
	value := 0
	for _, v := range A {
		if size == 0 {
			size++
			value = v
		} else {
			if value != v {
				size--
			} else {
				size++
			}
		}
	}
	if size > 0 {
		leader = value
	}
	counts := make([]int, len(A))
	count := 0
	for i, v := range A {
		if v == leader {
			count++
		}
		counts[i] = count
	}
	if count <= len(A)/2 {
		return 0
	}
	rst := 0
	for i, _ := range A {
		if counts[i] > (i+1)/2 && count-counts[i] > (len(A)-i-1)/2 {
			rst++
		}
	}
	return rst
}
