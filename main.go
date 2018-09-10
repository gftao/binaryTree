package main

import (
	"fmt"
	"binaryTree/bTree"
)

func main() {
	tree := &bTree.Node{Key: 6, Left: &bTree.Node{Key: 3}}
	tree.Insert(10)
	tree.Insert(4)
	tree.Insert(1)
	fmt.Println(tree)
	//tree.Delete(1, nil)
	//tree.Delete(2, nil)
	tree.Delete(6, nil)
	fmt.Println(tree)

}
