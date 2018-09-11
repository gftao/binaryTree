package main

import (
	"fmt"
	"binaryTree/bTree"
)

func main() {
	tree := &bTree.Node{Key: 6}

	/*
	err := tree.delete(6, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	*/

	fmt.Println(tree)
	tree, err := tree.DeleteNode(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tree)
	//insert
	tree, err = tree.InsertNode(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	tree, err = tree.InsertNode(4)
	tree, err = tree.InsertNode(12)
	//tree, err = tree.InsertNode(5)
	fmt.Println(tree)
	//delete
	tree.DeleteNode(10)
	fmt.Println(tree)

	////Max
	//fmt.Println(tree.Max())
	////Min
	//fmt.Println(tree.Min())

}
