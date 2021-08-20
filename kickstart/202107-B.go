package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//smallSetB202107BMain()
	bigSet202107BMain()
}

func smallSetB202107BMain() {
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

func bigSet202107BMain() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	caseNum := str2Int64(input.Text())

	for i := 0; i < int(caseNum); i++ {
		input.Scan()
		NK := getNumsInLine(input.Text())
		n, k := int(NK[0]), int(NK[1])
		intervals := make([][]int, n)
		for j := 0; j < n; j++ {
			input.Scan()
			interval := getNumsInLine(input.Text())
			intervals[j] = []int{int(interval[0]), int(interval[1])}
		}
		fmt.Printf("Case #%d: %d\n", i+1, B202107BigSet(intervals, k))
	}
}

// the spirit is the number of points lies in the intervals
func B202107BigSet(intervals [][]int, k int) int {

	// We construct a sorted map Mx which maps the coordinate X to the number of
	// additional intervals it lies within as compared to X-1.

	// That is, for each starting interval endpoint Li, we increment M(Li+1).
	// And for each ending interval endpoint Ri, we decrement M(Ri)

	// The above solution is based on the observation that the number of additional intervals obtained by a cut
	// at two consecutive points X and X+1 are the same, except, possibly, when some intervals start or end at these points
	var additionalOverlapsMap = make(map[int]int, 128)
	var orderedList []int
	for _, interval := range intervals {
		additionalOverlapsMap[interval[0]+1]++
		additionalOverlapsMap[interval[1]]--
	}
	for k := range additionalOverlapsMap {
		orderedList = append(orderedList, k)
	}
	sort.Ints(orderedList)

	// then iterate over the map in the sorted order of keys.
	// keep track of the number of overlapping intervals j
	// the previous key K1, the current key K2.
	// We increment Aj by K2-K1.
	// Now Aj can be used to compute the final solution as describe above
	K2, K1, j := 0, 0, 0
	var Aj = make(map[int]int, len(intervals)) // key: num of overlapping intervals, val: num of points in it
	for _, point := range orderedList {
		K2 = point
		Aj[j] += K2 - K1
		K1 = point
		j += additionalOverlapsMap[point]
	}

	// And With Aj, we can compute the final results
	var orderedJ []int
	for overlap := range Aj {
		orderedJ = append(orderedJ, overlap)
	}
	sort.Ints(orderedJ)
	finalAdditional := 0
	for i := len(orderedJ) - 1; i >= 0 && k > 0; i-- {
		curJ := orderedJ[i]
		curPoints := Aj[curJ]
		if k > curPoints {
			k -= curPoints
			finalAdditional += curJ * curPoints
		} else {
			finalAdditional += curJ * k
			k = 0
		}
	}

	// Constructing the map requires adding 2*N points to it, with each addition requiring O(logN).
	// There for time complexity is O(N*logN)
	return finalAdditional + len(intervals)
}

func getNumsInLine(str string) []int64 {
	vs := strings.Split(str, " ")
	var ret []int64
	for _, v := range vs {
		ret = append(ret, str2Int64(v))
	}
	return ret
}

func str2Int64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}
