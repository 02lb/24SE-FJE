package main

import "fmt"

// the tree-style node

type TreeNode struct {
	*JSONNode
	Children []*TreeNode
	last     bool //指示是否是最后一个孩子，以便在显示时使用
}

func NewTreeNode(name, value string) *TreeNode {
	return &TreeNode{JSONNode: NewJSONNode(name, value)}
}

func (n *TreeNode) Display(indent string) {
	if n == nil {
		return
	}
	if n.Name == "root" {
		for i, child := range n.Children {
			if i == len(n.Children)-1 {
				child.last = true
			}
			child.Display("")
		}
		return
	}

	if n.last {
		fmt.Printf("%s└─ %s: %s\n", indent, n.Name, n.Value)
	} else {
		fmt.Printf("%s├─ %s: %s\n", indent, n.Name, n.Value)
	}
	newIndent := ""
	if n.last {
		newIndent = indent + "   "
	} else {
		newIndent = indent + "│  "

	}
	for i, child := range n.Children {
		if i == len(n.Children)-1 {
			child.last = true
		}
		child.Display(newIndent)
	}
}

func (n *TreeNode) AddChild(child *TreeNode) {
	n.Children = append(n.Children, child)
}
