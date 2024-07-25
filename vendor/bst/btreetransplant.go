package bst

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if root == node {
		return rplc
	}

	if node == rplc {
		return root
	}

	if node.Parent != nil {
		if rplc != nil {
			rplc.Parent = node.Parent
		}
		if node.Parent.Left == node {
			node.Parent.Left = rplc
		} else {
			node.Parent.Right = rplc
		}
	}

	return root
}
