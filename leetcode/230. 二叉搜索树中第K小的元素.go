package main

var (
	R int
	K int
)

func kthSmallest(root *TreeNode, k int) int {
	R = 0
	K = 0
	kthSmallHelper(root, k)
	return R
}

// BST的特性是中序遍历就是升序
func kthSmallHelper(root *TreeNode, k int) {
	if root == nil {
		return
	}
	kthSmallHelper(root.Left, k)
	K++
	if K == k {
		R = root.Val
		return
	}
	kthSmallHelper(root.Right, k)
}

// 算树大小不是算排行
func treeSize(root *TreeNode, k int) int {
	R = 0
	treeSizeHelper(root, k)
	return R
}

// BST的特性是中序遍历就是升序
func treeSizeHelper(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	rank := treeSizeHelper(root.Left, k)
	if rank+1 == k {
		R = root.Val
	}
	return rank + 1 + treeSizeHelper(root.Right, k)
}
