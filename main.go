package main

import (
	"bst"
	"fmt"
)

func main() {
	// insert data
	{
		fmt.Println("===== BTreeInsertData =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		fmt.Println(root.Left.Data)
		fmt.Println(root.Data)
		fmt.Println(root.Right.Left.Data)
		fmt.Println(root.Right.Data)
	}
	fmt.Println()
	// apply in order
	{
		fmt.Println("===== BTreeApplyInorder =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		bst.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// apply pre order
	{
		fmt.Println("===== BTreeApplyPreorder =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		bst.BTreeApplyPreorder(root, fmt.Println)
	}
	fmt.Println()
	// search item
	{
		fmt.Println("===== BTreeSearchItem =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		selected := bst.BTreeSearchItem(root, "7")
		fmt.Print("Item selected -> ")
		if selected != nil {
			fmt.Println(selected.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Parent of selected item -> ")
		if selected.Parent != nil {
			fmt.Println(selected.Parent.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Left child of selected item -> ")
		if selected.Left != nil {
			fmt.Println(selected.Left.Data)
		} else {
			fmt.Println("nil")
		}
		fmt.Print("Right child of selected item -> ")
		if selected.Right != nil {
			fmt.Println(selected.Right.Data)
		} else {
			fmt.Println("nil")
		}
	}
	fmt.Println()
	// level count
	{
		fmt.Println("===== BTreeLevelCount =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		fmt.Println(bst.BTreeLevelCount(root))
	}
	fmt.Println()
	// is binary
	{
		fmt.Println("===== BTreeIsBinary =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		fmt.Println(bst.BTreeIsBinary(root))
	}
	fmt.Println()
	// apply by level
	{
		fmt.Println("===== BTreeApplyByLevel =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "3")
		bst.BTreeInsertData(root, "2")
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		bst.BTreeApplyByLevel(root, fmt.Println)
	}
	fmt.Println()
	// btree max
	{
		fmt.Println("===== BTreeMax =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		max := bst.BTreeMax(root)
		fmt.Println(max.Data)
		max = bst.BTreeMax(nil)
	}
	fmt.Println()
	// btree min
	{
		fmt.Println("===== BTreeMin =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		min := bst.BTreeMin(root)
		fmt.Println(min.Data)
		min = bst.BTreeMin(nil)
	}
	fmt.Println()
	// trans plant
	{
		fmt.Println("===== BTreeTransPlant =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		node := bst.BTreeSearchItem(root, "1")
		replacement := &bst.TreeNode{Data: "3"}
		root = bst.BTreeTransplant(root, node, replacement)
		bst.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// deletenode
	{
		fmt.Println("===== BTreeDeleteNode =====")
		root := &bst.TreeNode{Data: "4"}
		bst.BTreeInsertData(root, "1")
		bst.BTreeInsertData(root, "7")
		bst.BTreeInsertData(root, "5")
		bst.BTreeInsertData(root, "6")
		node := bst.BTreeSearchItem(root, "4")
		fmt.Println("Before delete:")
		bst.BTreeApplyInorder(root, fmt.Println)
		root = bst.BTreeDeleteNode(root, node)
		fmt.Println("After delete:")
		bst.BTreeApplyInorder(root, fmt.Println)
	}
}
