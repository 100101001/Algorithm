package main

func main() {
	//TestLongestValidParentheses()

	//testMergeTwoLists()
	//for i := 2; i < 100; i++ {
	//	fmt.Println(i, addIn(i))
	//}
	//fmt.Println(23, addIn(23))

	//KMP("ABABCABCAB", "ABCAB")

	//longestPalindrome("cbcbccde")

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	root.Right = &TreeNode{Val: 4}
	kthSmallest(root, 1)
}
