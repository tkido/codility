package main

const null = -1000000001

func Solution(A []int) int {
	m := make(map[int]int)
	leftLeaders := make([]int, len(A))
	leader := null
	leaderCount := 0
	for i, v := range A {
		m[v]++
		if m[v] > (i+1)/2 {
			leader = v
			leaderCount = m[v]
			leftLeaders[i] = v
		} else if leaderCount <= (i+1)/2 {
			leader = null
			leftLeaders[i] = null
		} else {
			leftLeaders[i] = leader
		}
	}
	m = make(map[int]int)
	rightLeaders := make([]int, len(A))
	leader = null
	leaderCount = 0
	for i := len(A) - 1; i >= 0; i-- {
		v := A[i]
		m[v]++
		if m[v] > (len(A)-i)/2 {
			leader = v
			leaderCount = m[v]
			rightLeaders[i] = v
		} else if leaderCount <= (len(A)-i)/2 {
			leader = null
			rightLeaders[i] = null
		} else {
			rightLeaders[i] = leader
		}
	}
	count := 0
	for i := 0; i < len(A)-1; i++ {
		if leftLeaders[i] != null && leftLeaders[i] == rightLeaders[i+1] {
			count++
		}
	}
	return count
}
