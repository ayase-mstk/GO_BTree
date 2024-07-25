package piscine

import "fmt"

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node == nil {
		return root
	}

	if node.Left == nil {
		fmt.Println("LEft node is nil")
		BTreeTransplant(root, node, node.Right)
		return root
	}
	if node.Right == nil {
		fmt.Println("Right node is nil")
		BTreeTransplant(root, node, node.Left)
		return root
	}

	// どっちの子ノードもある場合
	// nodeのRigtの最小のノードを抜き出して、nodeと置き換える
	var minNode *TreeNode = BTreeMin(node.Right)
	minNode.Left = node.Left
	node.Left.Parent = minNode
	BTreeTransplant(root, node, minNode)

	return root
}
