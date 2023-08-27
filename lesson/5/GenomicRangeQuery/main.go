package main

const (
	A = 0
	C = 1
	G = 2
	T = 3
)

func Solution(S string, P []int, Q []int) []int {
	N := len(S)
	sum := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		sum[i] = make([]int, 4)
	}
	ss := make([]int, N)
	for i, v := range S {
		switch v {
		case 'A':
			ss[i] = A
		case 'C':
			ss[i] = C
		case 'G':
			ss[i] = G
		case 'T':
			ss[i] = T
		default:
			panic("invalid input!!!!")
		}
	}
	for i, v := range ss {
		for j := 0; j < 4; j++ {
			sum[i+1][j] = sum[i][j]
		}
		sum[i+1][v]++
	}
	M := len(P)
	rst := make([]int, M)
	for i, p := range P {
		q := Q[i]
		for j := 0; j < 4; j++ {
			if sum[p][j] != sum[q+1][j] {
				rst[i] = j + 1
				break
			}
		}
	}
	return rst
}
