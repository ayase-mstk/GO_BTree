package main

import (
	"fmt"
	"piscine"
)

func main() {
	// insert data
	{
		fmt.Println("===== BTreeInsertData =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		fmt.Println(root.Left.Data)
		fmt.Println(root.Data)
		fmt.Println(root.Right.Left.Data)
		fmt.Println(root.Right.Data)
	}
	fmt.Println()
	// apply in order
	{
		fmt.Println("===== BTreeApplyInorder =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// apply pre order
	{
		fmt.Println("===== BTreeApplyPreorder =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyPreorder(root, fmt.Println)
	}
	fmt.Println()
	// search item
	{
		fmt.Println("===== BTreeSearchItem =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		selected := piscine.BTreeSearchItem(root, "7")
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
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		fmt.Println(piscine.BTreeLevelCount(root))
	}
	fmt.Println()
	// is binary
	{
		fmt.Println("===== BTreeIsBinary =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		fmt.Println(piscine.BTreeIsBinary(root))
	}
	fmt.Println()
	// apply by level
	{
		fmt.Println("===== BTreeApplyByLevel =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "3")
		piscine.BTreeInsertData(root, "2")
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyByLevel(root, fmt.Println)
	}
	fmt.Println()
	// btree max
	{
		fmt.Println("===== BTreeMax =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		max := piscine.BTreeMax(root)
		fmt.Println(max.Data)
		max = piscine.BTreeMax(nil)
	}
	fmt.Println()
	// btree min
	{
		fmt.Println("===== BTreeMin =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		min := piscine.BTreeMin(root)
		fmt.Println(min.Data)
		min = piscine.BTreeMin(nil)
	}
	fmt.Println()
	// trans plant
	{
		fmt.Println("===== BTreeTransPlant =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		node := piscine.BTreeSearchItem(root, "1")
		replacement := &piscine.TreeNode{Data: "3"}
		root = piscine.BTreeTransplant(root, node, replacement)
		piscine.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// deletenode
	{
		fmt.Println("===== BTreeDeleteNode =====")
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeInsertData(root, "6")
		node := piscine.BTreeSearchItem(root, "4")
		fmt.Println("Before delete:")
		piscine.BTreeApplyInorder(root, fmt.Println)
		root = piscine.BTreeDeleteNode(root, node)
		fmt.Println("After delete:")
		piscine.BTreeApplyInorder(root, fmt.Println)
	}
}
