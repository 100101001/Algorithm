package main

func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstringMW(s)
}

// LYX_TODO 内存占用率过高可优化
func lengthOfLongestSubstringMW(s string) int {

	window := make(map[byte]int, len(s))
	res := 0

	left, right := 0, 0
	for right < len(s) {
		add := s[right]
		right++
		// 右移更新
		window[add]++

		for window[add] > 1 {
			removed := s[left]
			left++
			window[removed]--
		}

		if res < right-left {
			res = right - left
		}
	}

	return res
}
