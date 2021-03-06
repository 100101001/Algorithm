package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaxHelper(nums, 0, len(nums)-1)
}

// 先序遍历
// 函数定义：返回数组段构造的树根
// 最大值作为根，对左数组段和右数组段递归作为左右子树
// 从数组构造一个最大二叉树
func constructMaxHelper(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	//
	max := nums[lo]
	maxIdx := lo
	for i := lo; i <= hi; i++ {
		if nums[i] > max {
			maxIdx = i
			max = nums[i]
		}
	}

	// LYX_TODO 边界条件需要自己注意有问题
	root := &TreeNode{Val: max}
	root.Left = constructMaxHelper(nums, lo, maxIdx-1)
	root.Right = constructMaxHelper(nums, maxIdx+1, hi)
	return root
}
