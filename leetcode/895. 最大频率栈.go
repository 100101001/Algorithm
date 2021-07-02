package main

type FreqStack struct {
	maxFreq  int
	freq2val map[int]Stack
	val2freq map[int]int
}

func FreqStackConstructor() FreqStack {
	return FreqStack{}
}

func (f *FreqStack) Push(val int) {

}

func (f *FreqStack) Pop() int {
	return 0
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
