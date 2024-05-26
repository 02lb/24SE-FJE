package main

type RectangleFactory struct{}

func (f *RectangleFactory) CreateNode(name, value string) AbstractNode {
	return NewRectangleNode(name, value)
}
