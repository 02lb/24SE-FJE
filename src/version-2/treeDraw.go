package main

import "fmt"

func (l *Leaf) TreeDraw(indent string) {
	value := ""
	if l.Value != "" {
		value = ": " + l.Value
	}
	if l.Last {
		fmt.Println(indent + "└─ " + l.Icon + l.Name + value)
	} else {
		fmt.Println(indent + "├─ " + l.Icon + l.Name + value)
	}
}

func (c *Container) TreeDraw(indent string) {
	//TODO: 实现树形风格的显示
	if c == nil {
		return
	}
	if c.Name == "root" && firstRoot {
		firstRoot = false
		//fmt.Println("这里只能进来一次")
		for idx, child := range c.Children {
			if idx == len(c.Children)-1 {
				child.SetLast(true)
			}
			child.Draw("")
		}
		return
	}
	if c.Last {
		// 这里处理了当Value为空时需要去掉冒号的情况
		value := ""
		if c.Value != "" {
			value = ": " + c.Value
		}
		fmt.Printf("%s└─ %s%s\n", indent, c.Icon+c.Name, value)
	} else {
		value := ""
		if c.Value != "" {
			value = ": " + c.Value
		}
		fmt.Printf("%s├─ %s%s\n", indent, c.Icon+c.Name, value)
	}

	newIndent := ""
	if c.Last {
		newIndent = indent + "   "
	} else {
		newIndent = indent + "│  "
	}
	for idx, child := range c.Children {
		if idx == len(c.Children)-1 {
			child.SetLast(true)
		}
		child.Draw(newIndent)
	}
}
