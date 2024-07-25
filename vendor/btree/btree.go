package btree

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
