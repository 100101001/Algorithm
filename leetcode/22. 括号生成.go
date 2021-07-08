package main

var GENRES []string

var INT2STR = []string{"(", ")"}

func generateParenthesis(n int) []string {
	GENRES = []string{}
	generateParenthesisHelper(n, []int{0, 0}, "")
	return GENRES
}

// counter 的 0 是左括号，1 是右括号
func generateParenthesisHelper(n int, counter []int, path string) {

	// 终止条件
	if counter[0] >= n && counter[1] >= n {
		GENRES = append(GENRES, path)
		return
	}

	for i := 0; i < 2; i++ {
		if !isValidChoice(counter, i, n) {
			continue
		}
		// 做选择
		counter[i]++
		// 下一步
		generateParenthesisHelper(n, counter, path+INT2STR[i])
		// 撤销选择
		counter[i]--
	}
}

func isValidChoice(counter []int, choice, n int) bool {
	// 做了选择以后右括号数大于左括号数
	counter[choice]++
	right := counter[1]
	left := counter[0]
	counter[choice]--
	// 右括号数量要小于等于左括号数
	return right <= left && right <= n && left <= n
}

//
func generateParenthesis1(n int) []string {
	return nil
}

func generateParenthesisHelper1() {

}
