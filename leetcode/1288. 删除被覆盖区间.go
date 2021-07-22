package main

import "sort"

// s1 <= s2 < e2 <= e1
// 按起点升序,  去掉一半
// e2 <= e1，由于是任意，迭代的时候前半段一直是满足的，后半段一直更新e1为当前最长即可。
// 注意按照迭代顺序 s1=s2的时候，令 e1 >= e2，即（起点升序，终点降序）
func removeCoveredIntervals(intervals [][]int) int {
	//
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			intervals[i][0] == intervals[j][0] && intervals[i][1] >= intervals[j][1]
	})

	e1 := intervals[0][1]
	count := 0
	for _, interval := range intervals {
		e2 := interval[1]
		if e2 <= e1 {
			count++
		} else {
			e1 = e2
		}
	}

	return count
}
