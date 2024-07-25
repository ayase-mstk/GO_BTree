package main

import (
	"btree"
	"fmt"
)

func main() {
	// insert data
	{
		fmt.Println("===== BTreeInsertData =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		fmt.Println(root.Left.Data)
		fmt.Println(root.Data)
		fmt.Println(root.Right.Left.Data)
		fmt.Println(root.Right.Data)
	}
	fmt.Println()
	// apply in order
	{
		fmt.Println("===== BTreeApplyInorder =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		btree.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// apply pre order
	{
		fmt.Println("===== BTreeApplyPreorder =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		btree.BTreeApplyPreorder(root, fmt.Println)
	}
	fmt.Println()
	// search item
	{
		fmt.Println("===== BTreeSearchItem =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		selected := btree.BTreeSearchItem(root, "7")
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
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		fmt.Println(btree.BTreeLevelCount(root))
	}
	fmt.Println()
	// is binary
	{
		fmt.Println("===== BTreeIsBinary =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		fmt.Println(btree.BTreeIsBinary(root))
	}
	fmt.Println()
	// apply by level
	{
		fmt.Println("===== BTreeApplyByLevel =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "3")
		btree.BTreeInsertData(root, "2")
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		btree.BTreeApplyByLevel(root, fmt.Println)
	}
	fmt.Println()
	// btree max
	{
		fmt.Println("===== BTreeMax =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		max := btree.BTreeMax(root)
		fmt.Println(max.Data)
		max = btree.BTreeMax(nil)
	}
	fmt.Println()
	// btree min
	{
		fmt.Println("===== BTreeMin =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		min := btree.BTreeMin(root)
		fmt.Println(min.Data)
		min = btree.BTreeMin(nil)
	}
	fmt.Println()
	// trans plant
	{
		fmt.Println("===== BTreeTransPlant =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		node := btree.BTreeSearchItem(root, "1")
		replacement := &btree.TreeNode{Data: "3"}
		root = btree.BTreeTransplant(root, node, replacement)
		btree.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// deletenode
	{
		fmt.Println("===== BTreeDeleteNode =====")
		root := &btree.TreeNode{Data: "4"}
		btree.BTreeInsertData(root, "1")
		btree.BTreeInsertData(root, "7")
		btree.BTreeInsertData(root, "5")
		btree.BTreeInsertData(root, "6")
		node := btree.BTreeSearchItem(root, "4")
		fmt.Println("Before delete:")
		btree.BTreeApplyInorder(root, fmt.Println)
		root = btree.BTreeDeleteNode(root, node)
		fmt.Println("After delete:")
		btree.BTreeApplyInorder(root, fmt.Println)
	}
}
