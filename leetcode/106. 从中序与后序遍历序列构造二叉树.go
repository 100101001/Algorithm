package main

func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	return buildTreeHelperIP(postorder, 0, len(postorder)-1, inorder, 0, len(inorder)-1)
}

// 构建一个树
func buildTreeHelperIP(postorder []int, postLo, postHi int, inorder []int, inLo, inHi int) *TreeNode {
	if postLo > postHi || inLo > inHi {
		return nil
	}

	rootVal := postorder[postHi]
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
	root.Left = buildTreeHelperIP(postorder, postLo, postLo+leftTreeSize-1, inorder, inLo, rootIdx-1)
	root.Right = buildTreeHelperIP(postorder, postLo+leftTreeSize, postHi-1, inorder, rootIdx+1, inHi)
	return root
}
