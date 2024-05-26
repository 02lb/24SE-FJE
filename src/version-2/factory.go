package main

type Factory interface {
	CreateContainer(name string, value string) *Container
	CreateLeaf(name string, value string) *Leaf
}

// TreeFactory 树形工厂
type TreeFactory struct {
	ContainerIcon string
	LeafIcon      string
}

func (f *TreeFactory) CreateContainer(name string, value string) *Container {
	return &Container{Name: name, Value: value, Icon: f.ContainerIcon, Style: "tree", Last: false}
}

func (f *TreeFactory) CreateLeaf(name string, value string) *Leaf {
	return &Leaf{Name: name, Value: value, Icon: f.LeafIcon, Style: "tree", Last: false}
}

// RectangleFactory 矩形工厂
type RectangleFactory struct {
	ContainerIcon string
	LeafIcon      string
}

func (f *RectangleFactory) CreateContainer(name string, value string) *Container {
	return &Container{Name: name, Value: value, Icon: f.ContainerIcon, Style: "rectangle", Last: false}
}

func (f *RectangleFactory) CreateLeaf(name string, value string) *Leaf {
	return &Leaf{Name: name, Value: value, Icon: f.LeafIcon, Style: "rectangle", Last: false}
}
