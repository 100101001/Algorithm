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
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	caseNum := str2Int64(input.Text())
	for i := 0; i < int(caseNum); i++ {
		input.Scan()
		NS := getNumsInLine(input.Text())
		n := int(NS[0])

		var problems []int
		for i := 0; i < n; i++ {
			input.Scan()
			ab := getNumsInLine(input.Text())
			if ab[0] == ab[1] {
				problems = append(problems, int(ab[0]))
				continue
			}
			for i := int(ab[0]); i <= int(ab[1]); i++ {
				problems = append(problems, i)
			}
		}
		sort.Ints(problems)

		fmt.Println(problems)

		var res string
		input.Scan()
		skills := getNumsInLine(input.Text())
		for _, skill := range skills {
			i := normalSearch(problems, int(skill))
			fmt.Println("i===", i)
			res += fmt.Sprintf(" %d", problems[i])
			problems = removeArrI(problems, i)
		}
		fmt.Printf("Case #%d:%s\n", i+1, res)
	}
}

func normalSearch(arr []int, skill int) int {
	if arr[0] > skill {
		return 0
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == skill {
			return i
		}
		if arr[i] < skill && arr[i+1] > skill {
			if skill-arr[i] <= arr[i+1]-skill {
				return i
			} else {
				return i + 1
			}
		}
	}
	return len(arr) - 1
}

func removeArrI(arr []int, target int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if i != target {
			res = append(res, arr[i])
		}
	}
	return res
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
