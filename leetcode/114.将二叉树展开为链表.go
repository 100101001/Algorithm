package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

// 将以n为根的树拉成右直链
func helper(n *TreeNode) {
	if n == nil {
		return
	}
	// 将左子树为根的树拉成右直链，将右子树为根的树拉成右直链
	helper(n.Left)
	helper(n.Right)

	//
	left := n.Left
	right := n.Right

	n.Left = nil
	n.Right = left

	var r = n
	for ; r.Right != nil; r = r.Right {
	}
	r.Right = right
}
