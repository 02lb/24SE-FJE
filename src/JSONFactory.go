package main

type JSONFactory interface {
	CreateNode(name, value string) AbstractNode
}
