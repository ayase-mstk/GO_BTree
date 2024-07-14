package main

import (
	"fmt"
	"piscine"
)

func main() {
	// insert data
	{
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
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyInorder(root, fmt.Println)
	}
	fmt.Println()
	// apply pre order
	{
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyPreorder(root, fmt.Println)
	}
	fmt.Println()
	// search item
	{
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
	{
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		fmt.Println(piscine.BTreeLevelCount(root))
	}
	fmt.Println()
}
