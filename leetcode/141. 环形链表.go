package main

func hasCycle(head *ListNode) bool {
	return hasCycleFastSlow(head)
}

//
func hasCycleFastSlow(head *ListNode) bool {

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
