package main

func Solution(A []int) int {
	size := len(A) + 1
	sum := make([]int, size)
	for i, v := range A {
		sum[i+1] = sum[i] + v
	}
	min := float64(10000)
	rst := 0
	for i := 0; i < size-3; i++ {
		avg2 := float64(sum[i+2]-sum[i]) / float64(2)
		if avg2 < min {
			min = avg2
			rst = i
		}
		avg3 := float64(sum[i+3]-sum[i]) / float64(3)
		if avg3 < min {
			min = avg3
			rst = i
		}
	}
	avg := float64(A[len(A)-2]+A[len(A)-1]) / float64(2)
	if avg < min {
		min = avg
		rst = len(A) - 2
	}
	return rst
}
