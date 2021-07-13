package main

func minWindows(s, t string) string {
	//
	return minWindowsMW(s, t)
}

// 滑窗O(n)解法，暴力需O(n^2)
func minWindowsMW(s, t string) string {

	left, right := 0, 0

	var need = make(map[byte]int, len(t))     // 需要的各字符数量, 无需更新
	var inWindow = make(map[byte]int, len(t)) // 窗口内的个字符数量，需更新
	var okNum int                             // 每次操作inWindow都要检查
	for _, v := range t {
		need[byte(v)]++
	}

	var (
		minCoverStr = s
		shrunk      bool
	)
	for right < len(s) {
		add := s[right]
		right++
		if _, ex := need[add]; ex {
			inWindow[add]++
			// 放入后刚好一样，这次添加导致一个字符满足数量条件
			if inWindow[add] == need[add] {
				okNum++
			}
		}

		// 收缩窗口是while满足收缩条件
		for okNum == len(need) {
			shrunk = true
			if len(minCoverStr) > right-left {
				minCoverStr = s[left:right]
			}

			moved := s[left]
			left++
			if _, ex := need[moved]; ex {
				// 移走前刚好一样，这次移除导致一个字符不满足数量条件
				if inWindow[moved] == need[moved] {
					okNum--
				}
				//
				inWindow[moved]--
			}
		}
	}

	if !shrunk {
		return ""
	}
	return minCoverStr
}
