package main

import (
	"fmt"
	"binaryTree/bTree"
)

func main() {
	tree := &bTree.Node{Key: 6}

	/*
	err := tree.Delete(6, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	*/
	fmt.Println(tree)
	//Insert
	tree.Insert(10)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	fmt.Println(tree)
	//Delete
	tree.Delete(10, nil)
	fmt.Println(tree)
	tree.Delete(6, nil)
	fmt.Println(tree)


	//Max
	//fmt.Println(tree.Max())
	//Min
	//fmt.Println(tree.Min())

}
