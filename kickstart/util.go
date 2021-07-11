package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		fmt.Printf("cur line is %s\n", s)
		counts[s]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d  %s\n", n, line)
		}
	}
}
