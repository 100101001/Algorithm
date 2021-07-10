package main

import "math"

var superEggDropMem [][]int

//
func superEggDrop(k int, n int) int {

	// 初始化备忘录
	superEggDropMem = make([][]int, k+1)
	for i := range superEggDropMem {
		superEggDropMem[i] = make([]int, n+1)
		for j := range superEggDropMem[i] {
			superEggDropMem[i][j] = -1
		}
	}

	return superEggDropHelper(k, n)
}

// 高楼扔鸡蛋
func superEggDropHelper(k, n int) int {

	// 鸡蛋只有1个，只能试n次
	if k == 1 {
		return n
	}
	// 楼层只有0层，不用试了
	if n == 0 {
		return 0
	}

	if superEggDropMem[k][n] >= 0 {
		return superEggDropMem[k][n]
	}

	// 暴力穷举
	var res = math.MaxInt64
	for i := 1; i <= n; i++ {

		// 摔碎的情况下还需要的最少尝试次数
		broken := superEggDropHelper(k-1, i-1)
		// 没摔碎的情况下还需要的最少尝试次数
		perfect := superEggDropHelper(k, n-i)

		bad := maxInt(broken, perfect)

		// 本次又算是尝试了一次
		res = minInt(res, bad+1)
	}

	superEggDropMem[k][n] = res

	return res
}
