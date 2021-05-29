package main

import "fmt"

var TestCase32 = []string{"(()", "())", "(())", "()()))", "(()()"}

func TestLongestValidParentheses() {
	for _, v := range TestCase32 {
		//longestValidParenthesesStack(v)
		longestValidParenthesesCounter(v)
	}
}

func longestValidParenthesesDP(s string) {
	// dp[i]代表以第i个字符结尾的有效的最长长度

	// 如果 s[i] 是 左括号，那么 dp[i] = 0
	// 如果 s[i] 是右括号，那么按s[i - 1] 是左括号或者右括号进行分类讨论
	// "...()" , dp[i] = dp[i-2] + 2
	// "...))", dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 (且 s[i-dp[i-1]-1] 是 '(')

	dp := make([]int, len(s), len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		}
		if s[i] == ')' && s[i-1] == ')' {
			if idx := i - dp[i-1] - 1; idx >= 0 && s[idx] == '(' {
				dp[i] = dp[i-1] + 2
				if idx-1 >= 0 {
					dp[i] += dp[idx-1]
				}
			}
		}
	}
	maxN := 0
	for i := 0; i < len(dp); i++ {
		if maxN < dp[i] {
			maxN = dp[i]
		}
	}

	fmt.Println(s, maxN)
}

func longestValidParenthesesStack(s string) {
	// 用栈是因为后进先出
	// 栈内存放目前遇到的所有左括号，右括号视下述情况放或不放
	// 遇到右括号弹出栈顶，和栈内剩余的开始标记位（匹配的最左位置）相减作为最长，如果栈内没有剩余代表当前右括号是目前的匹配的开始则放入栈内

	// 栈底始终是有效匹配的起点, 上面则是一堆左括号的位置
	maxN := 0
	stack := make([]int, 0, len(s))
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack)-1 >= 0 {
				maxN = maxInt(maxN, i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}
	fmt.Println(s, maxN)
}

func longestValidParenthesesCounter(s string) {
	// 左括号和右括号出现次数进行计数，当且仅当二者一样，计算为maxN
	// 然后直接将左右括号出现次数清零重新开始计数。
	// 对于 (() 的情况，左右来回一遍解决，完全反过来

	maxN := 0
	right, left := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right && 2*left > maxN {
			maxN = 2 * left
		}
		if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}
		if left == right && 2*left > maxN {
			maxN = 2 * left
		}
		if left > right {
			left, right = 0, 0
		}
	}

	fmt.Println(s, maxN)
}
