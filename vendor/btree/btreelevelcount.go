package btree

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
