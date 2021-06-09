package main

var mem [][]int

func numTrees(n int) int {
	mem = make([][]int, n+1)
	for i, _ := range mem {
		mem[i] = make([]int, n+1)
	}
	return numTreesHelper(1, n)
}

func numTreesHelper(lo, hi int) int {
	if lo > hi {
		return 1
	}
	if mem[lo][hi] != 0 {
		return mem[lo][hi]
	}
	num := 0
	for i := lo; i <= hi; i++ {
		lefts := numTreesHelper(lo, i-1)
		rights := numTreesHelper(i+1, hi)
		num += lefts * rights
	}
	mem[lo][hi] = num
	return num
}
