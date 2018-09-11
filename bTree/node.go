package bTree

import (
	"errors"
	"fmt"
	"strings"
)

var EORROOT error = errors.New("can not delete root node without child in the tree")

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
func (n *Node) InsertNode(key int) (*Node, error) {
	if n == nil {
		return &Node{Key: key}, nil
	} else {
		n.insert(key)
	}
	return n, nil
}
func (n *Node) insert(key int) {
	if n == nil {
		return
	}
	if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.insert(key)
		}
	} else if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.insert(key)
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

func (n *Node) DeleteNode(key int) (*Node, error) {
	err := n.delete(key, nil)
	if err == EORROOT {
		n = nil
		return nil, nil
	}
	return n, err
}

func (n *Node) delete(key int, p *Node) error {
	if n == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}

	if n.Key > key {
		n.Left.delete(key, n)
	} else if n.Key < key {
		n.Right.delete(key, n)
	} else {
		//该节点是头结点
		if p == nil {
			var replacement, replParent *Node
			if n.Left == nil && n.Right == nil {
				return EORROOT
			} else if n.Left == nil {
				replacement, replParent = n.Right.findMin(n)
			} else if n.Right == nil {
				replacement, replParent = n.Left.findMax(n)
			}
			n.Key = replacement.Key
			return replacement.delete(replacement.Key, replParent)
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
			return replacement.delete(replacement.Key, replParent)
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
