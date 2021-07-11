package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	caseNum := str2Int64(input.Text())

	for i := 0; i < int(caseNum); i++ {
		input.Scan()
		nc := getNumsInLine(input.Text())
		n := nc[0]
		c := nc[1]

		// 获取n个区间
		var tmpMap = map[int64]int{} // 记录了每个整数的最大值
		for j := 0; j < int(n); j++ {
			input.Scan()
			lr := getNumsInLine(input.Text())
			for ii := lr[0] + 1; ii < lr[1]; ii++ {
				tmpMap[ii]++
			}
		}

		res := getFirstC(tmpMap, int(c))
		fmt.Printf("Case #%d: %d\n", i+1, res+int(n))
	}

}

func getFirstC(tmpMap map[int64]int, c int) int {
	var res = make([]int, 0, len(tmpMap))
	for _, v := range tmpMap {
		res = append(res, v)
	}
	sort.Ints(res)
	var ret int
	for i, j := len(res)-1, 0; i >= 0 && j < c; i, j = i-1, j+1 {
		ret += res[i]
	}
	return ret
}
