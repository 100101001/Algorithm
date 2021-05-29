package main

func addIn(i int) int {

	m := make(map[int]bool, i)
	left := 1
	ret := 1
	for !m[left] {
		m[left] = true
		left = (left*10 + 1) % i
		if left == 0 {
			return ret + 1
		}
		ret = ret + 1
	}

	return -1
}
