package main

var firstRoot bool = true //特判第一个root节点
const maxLength = 50      //矩阵对齐最大长度

// Component 接口：中间节点和叶节点都实现这个接口
type Component interface {
	Draw(string)
	SetLast(bool) //设置是否是最后一个child
	SetHead(bool)
}

// 区分中间节点和叶节点的主要原因：需要使用不同的Icon来显示
type Leaf struct {
	Name  string
	Value string
	Icon  string
	Last  bool
	Head  bool
	Style string
}

type Container struct {
	Name     string
	Value    string
	Icon     string
	Last     bool //用于判断是否是最后一个child，以便在显示时使用
	Head     bool
	Style    string
	Children []Component
}

func (l *Leaf) Draw(indent string) {
	//TODO: 如何在这里判断是TreeStyle还是RectangleStyle
	switch l.Style {
	case "tree":
		l.TreeDraw(indent)
	case "rectangle":
		l.RectangleDraw(indent)
	default:
		panic("Unknown style")
	}
}

func (c *Container) Draw(indent string) {
	switch c.Style {
	case "tree":
		c.TreeDraw(indent)
	case "rectangle":
		c.RectangleDraw(indent)
	default:
		panic("Unknown style")
	}
}

// 以下是辅助工具方法
func (l *Leaf) SetLast(b bool) {
	l.Last = b
}

func (l *Leaf) SetHead(b bool) {
	l.Head = b

}

func (c *Container) AddChild(child Component) {
	c.Children = append(c.Children, child)
}

func (c *Container) SetLast(b bool) {
	c.Last = b
}

func (c *Container) SetHead(b bool) {
	c.Head = b
}
