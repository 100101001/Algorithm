package main

import "fmt"

//https://www.acwing.com/activity/content/code/content/43108/
//http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
// s是字符串，p是待匹配字符串
func KMP(s, p string) {

	// 构造部分匹配表
	helper := make([]int, len(p))
	for i, j := 1, 0; i < len(p); i++ {
		// j回跳到子串中的匹配部分
		for j > 0 && p[i] != p[j] {
			j = helper[j]
		}
		// 基于子串中最长匹配部分，看能否拓展1位
		if p[i] == p[j] {
			j++
		}
		helper[i] = j
	}
	fmt.Println(helper)

	// 借助部分匹配表
	for i, j := 0, 0; i < len(s); i++ {
		for s[i] != p[j] && j > 0 {
			j = helper[j]
		}
		if s[i] == p[j] {
			j++
		}
		if j == len(p) {
			fmt.Println("匹配成功")
			j = helper[len(p)-1]
		}
	}
}
