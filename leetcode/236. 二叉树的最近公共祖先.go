package main

// 后序遍历，
// 函数定义：如果p,q在root的树中，是的话root就是lca，返回lca；
// 函数定义：如果p,q有一个不在root的树中，则返回p,q自己，停止继续往子树遍历的过程；
// 后序遍历之前：判断节点是不是p,q是的话就返回，停止继续往子树遍历；(此处如果p,在以q为根的子树内，就停止也是合理的)
// 后序遍历之后：返回lca；如果p,q在两颗子树上返回root作为lca; 否则返回子树中已找到的lca；
// 由于判断p,q是否在root的树中是后序遍历，所以依据后续遍历的特质，子节点先被遍历到，因此返回的root就是LCA
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 两个都不为空，返回root
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
