package main

func findRotateSteps(ring string, key string) int {
	// dp[i][j] 处于ring[i](==key[j]]) 拼出 key[..j]，所需最少拨动和摁键次数
	// dp[i][j] = min({dp[x][j-1]+ min(n-abs(i-x), abs(i-x)) + 1 | x ∈ (pos[j-1]), n=len(ring) })
	// n=len(ring), m=len(key), min({dp[x][m-1] | x ∈ (pos[j-1])})

	var pos = make([][]int, len(key))
	for i, k := range key {
		for j, r := range ring {
			if r == k {
				pos[i] = append(pos[i], j)
			}
		}
	}

	n := len(ring)
	m := len(key)

	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = 100 * 100 * 2
		}
	}

	for _, p := range pos[0] {
		rotate := minIntV2(n-p, p)
		press := 1
		dp[p][0] = rotate + press
	}

	for j := 1; j < len(key); j++ {
		for _, p := range pos[j] {
			for _, x := range pos[j-1] {
				rotate := minIntV2(n-absInt(x-p), absInt(x-p))
				press := 1
				dp[p][j] = minIntV2(dp[p][j], dp[x][j-1]+rotate+press)
			}
		}
	}

	var res = 100 * 100 * 2
	for _, p := range pos[m-1] {
		res = minIntV2(res, dp[p][m-1])
	}
	return res
}
