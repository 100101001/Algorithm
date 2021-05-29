package main

func UglyNumber(n int) int {
	if n < 1 {
		return 0
	}

	uglyParam := &UglyParam{}
	for i := 0; i < n; i++ {
		uglyParam.Next()
	}

	return 0
}

type UglyParam struct {
	X2   int
	Y3   int
	Z5   int
	Cur  int
	Prev int
}

func (param *UglyParam) Next() {

}
