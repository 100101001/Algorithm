package main

func checkInclusion(s1 string, s2 string) bool {
	//
	return checkInclusionMW(s1, s2)
}

func checkInclusionMW(s1 string, s2 string) bool {

	left, right := 0, 0
	var (
		okNum    int
		need     = make(map[byte]int, len(s2))
		inWindow = make(map[byte]int, len(s2))
	)
	for _, s := range s2 {
		need[byte(s)]++
	}

	for right < len(s1) {

		add := s1[right]
		right++
		if _, ex := need[add]; ex {
			inWindow[add]++
			if inWindow[add] == need[add] {
				okNum++
			}
		}

		for right-left >= len(s2) {
			if okNum == len(need) {
				return true
			}
			moved := s1[left]
			left++
			if _, ex := need[moved]; ex {
				if inWindow[moved] == need[moved] {
					okNum--
				}
				inWindow[moved]--
			}
		}
	}

	return false
}
