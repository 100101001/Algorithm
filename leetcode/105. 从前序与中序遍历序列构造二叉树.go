package main

func buildTreePI(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

// 构建一个树
func buildTreeHelper(preorder []int, preLo, preHi int, inorder []int, inLo, inHi int) *TreeNode {

	if preLo > preHi || inLo > inHi {
		return nil
	}

	rootVal := preorder[preLo]
	leftTreeSize := 0
	rootIdx := 0
	for i := inLo; i <= inHi; i++ {
		if inorder[i] == rootVal {
			leftTreeSize = i - inLo
			rootIdx = i
			break
		}
	}

	root := &TreeNode{Val: rootVal}
	root.Left = buildTreeHelper(preorder, preLo+1, preLo+leftTreeSize, inorder, inLo, rootIdx-1)
	root.Right = buildTreeHelper(preorder, preLo+leftTreeSize+1, preHi, inorder, rootIdx+1, inHi)
	return root
}
