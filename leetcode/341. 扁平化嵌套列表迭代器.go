package main

type NestedInteger struct {
	isInteger bool
	i         int
	l         []*NestedInteger
}

func (i NestedInteger) IsInteger() bool { return i.isInteger }

func (i NestedInteger) GetInteger() int {
	if i.isInteger {
		return i.i
	}
	panic("not a single integer")
}

func (i *NestedInteger) SetInteger(value int) {
	i.isInteger = true
	i.i = value
	i.l = nil
}

func (i *NestedInteger) Add(elem NestedInteger) {
	i.isInteger = false
	i.l = append(i.l, &elem)
}

func (i NestedInteger) GetList() []*NestedInteger {
	if !i.isInteger {
		return i.l
	}
	panic("not list")
}

type NestedIterator struct {
	l []*NestedInteger
}

func NestedIteratorConstructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{l: nestedList}
}

func (n *NestedIterator) Next() int {

	return 0
}

func (n *NestedIterator) HasNext() bool {
	return false
}
