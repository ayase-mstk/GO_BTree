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
	// apply in order
	{
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyInorder(root, fmt.Println)
	}
	// apply pre order
	{
		root := &piscine.TreeNode{Data: "4"}
		piscine.BTreeInsertData(root, "1")
		piscine.BTreeInsertData(root, "7")
		piscine.BTreeInsertData(root, "5")
		piscine.BTreeApplyPreorder(root, fmt.Println)
	}
}
