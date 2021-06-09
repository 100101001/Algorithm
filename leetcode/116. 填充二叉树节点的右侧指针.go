package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

// 连接相邻两个树
func connectTwoNode(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}

	// 连接相邻节点本身
	left.Next = right

	// 分别连接两个树的子层相邻
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)
	connectTwoNode(left.Right, right.Left)
}
