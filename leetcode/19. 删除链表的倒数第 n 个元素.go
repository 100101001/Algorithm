package main

// 快慢指针的移动差
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEndFastSlow(head, n)
}

func removeNthFromEndFastSlow(head *ListNode, n int) *ListNode {
	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快指针走到头代表第n个节点是头节点
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
}