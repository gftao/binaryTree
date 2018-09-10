package main

import (
	"fmt"
	"binaryTree/bTree"
)

func main() {
	tree := &bTree.Node{Key: 6, Left: &bTree.Node{Key: 3}}
	//Insert
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(1)
	fmt.Println(tree)
	//Delete
	tree.Delete(10, nil)
	fmt.Println(tree)
	//Search
	fmt.Println(tree.Search(6))
	tree.Delete(6, nil)
	fmt.Println(tree)
	fmt.Println(tree.Search(6))

	//Max
	fmt.Println(tree.Max())
	//Min
	fmt.Println(tree.Min())

}
