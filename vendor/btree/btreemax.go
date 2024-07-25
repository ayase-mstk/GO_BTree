package btree

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root != nil && root.Right == nil {
		return root
	}

	return BTreeMax(root.Right)
}
