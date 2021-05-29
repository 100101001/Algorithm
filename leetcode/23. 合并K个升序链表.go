package main

import "fmt"

var MergeTwoListCase = [][]int{{1, 3, 7}, {4, 5, 6, 9}, {2, 3, 7, 8, 10}, {-5, -1, 3, 9, 10}, {-10, 0, 100}}
var TestMerge2ListCase []*ListNode

func testMergeTwoLists() {
	PrintList(mergeKLists(TestMerge2ListCase))
	//PrintList(mergeTwoLists(TestMerge2ListCase[0], TestMerge2ListCase[1]))
}

func PrintList(t *ListNode) {
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}

func init() {
	for _, l := range MergeTwoListCase {
		var t = &ListNode{}
		var tt = t
		for _, item := range l {
			h := &ListNode{Val: item}
			tt.Next = h
			tt = tt.Next
		}
		TestMerge2ListCase = append(TestMerge2ListCase, t.Next)
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并k个链表方法1
func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		var nextLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				nextLists = append(nextLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				nextLists = append(nextLists, lists[i])
			}
		}
		lists = nextLists
	}
	if len(lists) == 0 {
		return nil
	}

	return lists[0]
}

func mergeTwoLists(first, second *ListNode) *ListNode {
	//
	head := &ListNode{}
	firstP, secondP, newT := first, second, head
	for firstP != nil && secondP != nil {
		if firstP.Val < secondP.Val {
			newT.Next = firstP
			firstP = firstP.Next
		} else {
			newT.Next = secondP
			secondP = secondP.Next
		}
		newT = newT.Next
	}

	for firstP != nil {
		newT.Next = firstP
		firstP = firstP.Next
		newT = newT.Next
	}

	for secondP != nil {
		newT.Next = secondP
		secondP = secondP.Next
		newT = newT.Next
	}
	return head.Next
}


func mergeKLists2(lists []*ListNode) *ListNode {
	return nil
}