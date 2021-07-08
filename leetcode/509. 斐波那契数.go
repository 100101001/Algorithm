package main


// 状态压缩，自底向上
// base case : n==0, n==1
func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	var prev1, prev2 = 1, 1
	var res int
	for i := 3; i <= n; i++ {
		res = prev1 + prev2
		tmp := prev1
		prev1 = res
		prev2 = tmp
	}
	return res
}
