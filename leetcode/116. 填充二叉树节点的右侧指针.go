package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 填充两个相邻节点的指针
// 后序遍历
// 树中：让左孩子指向右孩子
// 跨树：让左子树最右侧节点指向右子树最左侧节点
// 先序遍历（左孩子指向右孩子）
// 最小结构不是三角形，而是子和子子层的完全二叉树形，所有两个子树相连
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
