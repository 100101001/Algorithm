package main

import "fmt"

//https://zhuanlan.zhihu.com/p/70532099
func getProcessedStr(origin string) string {
	if len(origin) == 0 {
		return "^$"
	}
	var s = "^"
	for _, v := range origin {
		s += fmt.Sprintf("#%s", string(v))
	}
	return s + "#$"
}


// 马拉车算法
func longestPalindrome(origin string) string {

	origin = getProcessedStr(origin)
	n := len(origin)
	P := make([]int, n)
	C, R := 0, 0
	for i := 1; i < n-1; i++ {
		iMirror := 2*C - i

		if R > i {
			P[i] = minInt(R-i, P[iMirror]) // 防止超出R
		} else {
			P[i] = 0 // i==R
		}

		// 碰到之前讲的三种情况时候，需要利用中心扩展法
		for origin[i+1+P[i]] == origin[i-1-P[i]] {
			P[i]++
		}

		// 更新C和R
		if i+P[i] > R {
			R = i + P[i]
			C = i
		}
	}

	return ""
}
