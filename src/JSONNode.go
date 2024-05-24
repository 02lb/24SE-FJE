package main

import "fmt"

type JSONNode struct {
	Name     string
	Value    string
	Children []*JSONNode
}

func NewJSONNode(name, value string) *JSONNode {
	return &JSONNode{Name: name, Value: value}
}

func (n *JSONNode) AddChild(child *JSONNode) {
	n.Children = append(n.Children, child)
}

func (n *JSONNode) Display(indent string) {
	// display the node
	fmt.Println("The Display method is not implemented yet!")
}
