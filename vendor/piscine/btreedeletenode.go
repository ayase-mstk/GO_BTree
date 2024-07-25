package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node == nil {
		return root
	}

	if node.Left == nil {
		root = BTreeTransplant(root, node, node.Right)
		return root
	}
	if node.Right == nil {
		root = BTreeTransplant(root, node, node.Left)
		return root
	}

	// どっちの子ノードもある場合
	var minNode *TreeNode = BTreeMin(node.Right)
	// nodeのRigtの最小のノードがnodeのRightでなければ、そのRightと入れ替える
	if minNode.Parent != node {
		root = BTreeTransplant(root, minNode, minNode.Right)
		minNode.Right = node.Right
		node.Left.Parent = minNode
	}
	// nodeのRigtの最小のノードを抜き出して、nodeと置き換える
	root = BTreeTransplant(root, node, minNode)
	minNode.Left = node.Left
	node.Left.Parent = minNode

	return root
}
