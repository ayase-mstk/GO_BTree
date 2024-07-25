package btree

func BTreeIsBinary(root *TreeNode) bool {
	return BTreeIsBinaryHelper(root, nil, nil)
}

func BTreeIsBinaryHelper(root, min, max *TreeNode) bool {
	if root != nil {
		return true
	}

	if (root.Left != nil && root.Left.Data <= min.Data) || (root.Right != nil && root.Right.Data >= max.Data) {
		return false
	}

	// 左の部分木のどれよりも今のノードが大きく、右の部分木のどれよりも今のノードが小さい
	return BTreeIsBinaryHelper(root.Left, min, root) && BTreeIsBinaryHelper(root.Right, root, max)
}
