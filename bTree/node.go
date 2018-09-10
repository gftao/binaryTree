package bTree

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.Key > key {
		return n.Left.Search(key)
	} else if n.Key < key {
		return n.Right.Search(key)
	}
	return true
}
func (n *Node) Insert(key int) {
	if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()

}
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if parent == nil {
		return n, nil
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}
func (n *Node) findMin(parent *Node) (*Node, *Node) {
	if parent == nil {
		return n, nil
	}
	if n.Left == nil {
		return n, parent
	}
	return n.Left.findMin(n)
}
func (n *Node) replaceNode(parent, replacement *Node) {
	//头结点
	//if parent == nil && replacement == nil {
	//	return
	//}
	if n == parent.Left {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}
}

func (n *Node) Delete(key int, p *Node) error {
	if n == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}

	if n.Key > key {
		n.Left.Delete(key, n)
	} else if n.Key < key {
		n.Right.Delete(key, n)
	} else {
		//该节点是头结点
		if p == nil {
			var replacement, replParent *Node
			if n.Left == nil && n.Right == nil {
				return errors.New("can not delete root node without child in the tree")
			} else if n.Left == nil {
				replacement, replParent = n.Right.findMin(n)
			} else if n.Right == nil {
				replacement, replParent = n.Left.findMax(n)
			} else {
				replacement, replParent = n.Left.findMax(n)
			}
			n.Key = replacement.Key
			return replacement.Delete(replacement.Key, replParent)
		} else {
			if n.Left == nil && n.Right == nil {
				n.replaceNode(p, nil)
				return nil
			} else if n.Left == nil {
				n.replaceNode(p, n.Right)
				return nil
			} else if n.Right == nil {
				n.replaceNode(p, n.Left)
				return nil
			}
			replacement, replParent := n.Left.findMax(n)
			n.Key = replacement.Key
			return replacement.Delete(replacement.Key, replParent)
		}
	}
	return nil
}

func (n *Node) String() string {
	if n == nil {
		return "empty node"
	}
	vs := make([]int, 0, 10)
	n.Traverse(n, func(node *Node) {
		//fmt.Println("->", node.Key)
		vs = append(vs, node.Key)
	})
	s := ""
	for i, _ := range vs {
		s += fmt.Sprintf("%d ", vs[i])
	}
	return strings.TrimSpace(s)
}

func (n *Node) Traverse(c *Node, f func(*Node)) {
	if c == nil {
		return
	}
	//f(c) //先序遍历
	n.Traverse(c.Left, f)
	f(c) //中序遍历
	n.Traverse(c.Right, f)
	//f(c)  //后序遍历
}
