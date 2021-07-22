package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	overlapped := maxNotOverlapIntervalsGreedy(intervals)
	if overlapped == 0 {
		return len(intervals) - 1
	}
	return len(intervals) - overlapped
}

// 贪心思想
// 区间问题：按起点或者终点排序，并且需要注意当相同的时候另一个端点该升序还是降序
// 1：按起点排序（s1 < s2 < e1）， 如果按照区间起点升序的话，就可以省去第一步的比较；算相交的时候方便，不用互相比较；另外这里一致的时候第二顺序位是升序，原因在于 s1 <s2 < e1，e1 越小， s2 越不容易与之相交
// 2：考虑边界case比如全部相交的时候函数返回0，该移除多少。比如上述所提第一排序字段相等的时候第二顺序排序位该怎么做
// 1'：按终点排序（s2 < e1 < e2）， 如果按照区间终点升序的话，就可以省去第二步e2>e1的比较；算相交的时候方便，不用互相比较；且第二排序无需考虑
// 贪心思想是end尽可能短。
func maxNotOverlapIntervalsGreedy(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})

	fmt.Println(intervals)

	count := 0
	compareEnd := intervals[0][1]
	for _, v := range intervals {
		if v[0] >= compareEnd {
			compareEnd = v[1]
			count++
		}
	}
	if count > 0 {
		count++ // 初始两个区间互不相交自己没算上1
	}
	return count
}
