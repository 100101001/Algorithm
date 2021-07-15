package main

func detectCycle(head *ListNode) *ListNode {
	return detectCycleFastSlow(head)
}

func detectCycleFastSlow(head *ListNode) *ListNode {

	var hasLoop bool
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			hasLoop = true
			break
		}
	}
	if !hasLoop {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
