package main

func calculateMinimumHP(dungeon [][]int) int {
	// dp[i][j] 从(i, j) => (m, n) 耗血最小
	var dp = make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[0]))
	}

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
				dp[i][j] = getBlood(-dungeon[i][j])
			} else if i == len(dungeon)-1 {
				dp[i][j] = getBlood(dp[i][j+1] - dungeon[i][j])
			} else if j == len(dungeon[0])-1 {
				dp[i][j] = getBlood(dp[i+1][j] - dungeon[i][j])
			} else {
				dp[i][j] = getBlood(minIntV2(dp[i][j+1], dp[i+1][j]) - dungeon[i][j])
			}
		}
	}

	return dp[0][0] + 1
}

func getBlood(grid int) int {
	if grid < 0 {
		return 0
	} else {
		return grid
	}
}
