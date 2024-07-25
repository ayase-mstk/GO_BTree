package btree

// queueが使えたらもっと簡単に実装できそう。

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	level := BTreeLevelCount(root)

	for lc := 0; lc < level; lc++ {
		BTreeApplyByLevelHelper(root, f, lc)
	}
}

func BTreeApplyByLevelHelper(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	}

	if root.Left != nil {
		BTreeApplyByLevelHelper(root.Left, f, level-1)
	}
	if root.Right != nil {
		BTreeApplyByLevelHelper(root.Right, f, level-1)
	}
}
