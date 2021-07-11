package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// scan能够按行扫描输入
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	caseNumS := input.Text()
	caseNum := str2Int64(caseNumS)

	for i := 0; i < int(caseNum); i++ {

		input.Scan()
		line1 := getNumsInLine(input.Text())
		input.Scan()
		line2 := getNumsInLine(input.Text())
		input.Scan()
		line3 := getNumsInLine(input.Text())

		leftDiagnose := line1[0] + line3[2]
		middleCol := line1[1] + line3[1]
		middleRow := line2[0] + line2[1]
		rightDiagnose := line3[0] + line1[2]

		var counter = map[int64]int{}
		if leftDiagnose%2 == 0 {
			counter[leftDiagnose]++
		}
		if middleCol%2 == 0 {
			counter[middleCol]++
		}
		if middleRow%2 == 0 {
			counter[middleRow]++
		}
		if rightDiagnose%2 == 0 {
			counter[rightDiagnose]++
		}
		var (
			maxV = 0
		)
		for _, v := range counter {
			if maxV < v {
				maxV = v
			}
		}

		if line1[0]+line1[2] == 2*line1[1] {
			maxV++
		}
		if line3[0]+line3[2] == 2*line3[1] {
			maxV++
		}
		if line1[0]+line3[0] == 2*line2[0] {
			maxV++
		}
		if line1[2]+line3[2] == 2*line2[1] {
			maxV++
		}

		fmt.Printf("Case #%d: %d\n", i+1, maxV)
	}
}
