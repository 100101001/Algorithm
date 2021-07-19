package main

//子序列问题：i,j双指针，以字符
func minDistance(word1 string, word2 string) int {
	return minDistanceDP(word1, word2)
}

var minDistanceMem [][]int

func minDistanceDFS(word1 string, word2 string) int {
	minDistanceMem = fillTwoArrayInt(len(word1), len(word2), -1)
	minDistanceDFSHelper(word1, word2, len(word1), len(word2))
	return minDistanceMem[len(word1)][len(word2)]
}

// 没法理解最优子结构的原因是没有穷举（穷举后就是最优子结构了）
func minDistanceDFSHelper(word1 string, word2 string, i, j int) int {

	if i < 0 {
		return maxIntV2(0, j)
	}
	if j < 0 {
		return maxIntV2(0, i)
	}

	if minDistanceMem[i][j] != -1 {
		return minDistanceMem[i][j]
	}

	if word1[i] == word2[j] {
		minDistanceMem[i][j] = minDistanceDFSHelper(word1, word2, i-1, j-1)
	} else {
		minDistanceMem[i][j] = minIntV2(
			minDistanceDFSHelper(word1, word2, i-1, j),   // 删除
			minDistanceDFSHelper(word1, word2, i-1, j-1), // 替换
			minDistanceDFSHelper(word1, word2, i, j-1),   // 插入
		) + 1
	}
	return minDistanceMem[i][j]
}

// s1[..i], s2[...j]
func minDistanceDP(word1 string, word2 string) int {

	maxOp := len(word2) + len(word1) + 1
	var dp = fillTwoArrayInt(len(word1)+1, len(word2)+1, maxOp)

	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	// ...i, ...j
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minIntV2(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
