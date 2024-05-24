package main

type TreeFactory struct{}

func (f *TreeFactory) CreateNode(name, value string) AbstractNode {
	return NewTreeNode(name, value)
}
