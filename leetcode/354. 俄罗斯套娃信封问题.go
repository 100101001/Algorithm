package main

import "sort"

// w升序，h降序; 在h中寻找LIS即可。（降序的原因是w1>w2 && h1>h2才能嵌套）
func maxEnvelopes(envelopes [][]int) int {

	// 将信封按 w一排序位升序，h二排序位降序
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})

	// 利用LIS的二分算法，在h序列中查找最大升序子序列即可
	var top = make([]int, len(envelopes))
	var piles int
	for i := 0; i < len(envelopes); i++ {
		poker := envelopes[i][1]
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] == poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else if top[mid] > poker {
				right = mid
			}
		}

		if left == piles {
			piles++
		}
		top[left] = poker
	}

	return piles
}
