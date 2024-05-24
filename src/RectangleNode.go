package main

import "fmt"

type RectangleNode struct {
	*JSONNode
	Children  []*RectangleNode
	maxLength int
	head      bool
	end       bool
}

func NewRectangleNode(name, value string) *RectangleNode {
	//TODO：如何设置maxLength这个参数，或者启发性调整
	return &RectangleNode{JSONNode: NewJSONNode(name, value), maxLength: 50, head: false, end: true}
}

func (n *RectangleNode) Display(indent string) {
	if n == nil {
		return
	}
	//首先进行root初始化
	if n.Name == "root" {
		for i, child := range n.Children {
			if i == 0 {
				child.head = true
			}
			if i == len(n.Children)-1 {
				child.end = true
			} else {
				child.end = false
			}
			child.Display("")
		}
		return
	}

	if n.end {
		//判断是否是递归最后一个节点
		if len(n.Children) == 0 {
			// 输出矩阵的最后一行
			newindent := "└"
			for i := 0; i < len(indent)-1; i++ {
				newindent += "—"
			}
			var str string = fmt.Sprintf("%s└─ %s: %s", newindent, n.Name, n.Value)
			var str_thesamelen = fmt.Sprintf("%s└─ %s: %s", indent, n.Name, n.Value)
			//主要是符号长度和对对齐比较麻烦
			for i := 0; i < n.maxLength-len(str_thesamelen); i++ {
				str += "—"
			}
			str += "┘"
			fmt.Println(str)

			return
		}
	}

	if n.head {
		var str string = fmt.Sprintf("%s┌─ %s: %s", indent, n.Name, n.Value)
		len_str := len(str)
		for i := 0; i < n.maxLength-len_str; i++ {
			str += "—"
		}
		str += "┐"
		fmt.Println(str)

	} else {
		var str string = fmt.Sprintf("%s├─ %s: %s", indent, n.Name, n.Value)
		len_str := len(str)
		for i := 0; i < n.maxLength-len_str; i++ {
			str += "—"
		}
		str += "┤"
		fmt.Println(str)
	}

	newIndent := indent + "|  "
	for idx, child := range n.Children {
		if idx == len(n.Children)-1 && n.end {
			child.end = true
		} else {
			child.end = false
		}
		child.Display(newIndent)
	}

}

func (n *RectangleNode) AddChild(child *RectangleNode) {
	n.Children = append(n.Children, child)
}
