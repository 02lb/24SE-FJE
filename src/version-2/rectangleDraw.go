package main

import "fmt"

func (l *Leaf) RectangleDraw(indent string) {
	if l.Last {
		//已经是递归的最后一个节点
		newIndent := "└"
		for i := 0; i < len(indent)-1; i++ {
			newIndent += "—"
		}
		var val string = l.Value
		if val != "" {
			val = ": " + val
		}

		var str = fmt.Sprintf("%s└─ %s %s", newIndent, l.Icon+l.Name, val)
		var strTheSameLen = fmt.Sprintf("%s└─ %s %s", indent, l.Icon+l.Name, val)
		//主要是符号长度和对对齐比较麻烦
		for i := 0; i < maxLength-len(strTheSameLen); i++ {
			str += "—"
		}
		str += "┘"
		fmt.Println(str)

		return
	}

	if l.Head {
		var val string = l.Value
		if val != "" {
			val = ": " + val
		}

		var str = fmt.Sprintf("%s┌─ %s %s", indent, l.Icon+l.Name, val)
		lenStr := len(str)
		for i := 0; i < maxLength-lenStr; i++ {
			str += "—"
		}
		str += "┐"
		fmt.Printf("l.Head = true")
		fmt.Println(str)
	} else {
		var val string = l.Value
		if val != "" {
			val = ": " + val
		}

		var str = fmt.Sprintf("%s├─ %s %s", indent, l.Icon+l.Name, val)
		lenStr := len(str)
		for i := 0; i < maxLength-lenStr; i++ {
			str += "—"
		}
		str += "┤"
		fmt.Println(str)
	}

}

func (c *Container) RectangleDraw(indent string) {
	//TODO: 实现矩形风格的显示
	if c == nil {
		return
	}
	if c.Name == "root" && firstRoot {
		firstRoot = false
		for idx, child := range c.Children {
			if idx == 0 {
				child.SetHead(true)
			}
			if idx == len(c.Children)-1 {
				child.SetLast(true)
			}
			child.Draw("")
		}
		return
	}

	if c.Last {
		//判断是否是递归最后一个节点
		if len(c.Children) == 0 {
			// 输出矩阵的最后一行
			newIndent := "└"
			for i := 0; i < len(indent)-1; i++ {
				newIndent += "—"
			}
			var val string = c.Value
			if val != "" {
				val = ": " + val
			}

			var str = fmt.Sprintf("%s└─ % %s", newIndent, c.Icon+c.Name, val)
			var strTheSameLen = fmt.Sprintf("%s└─ %s %s", indent, c.Icon+c.Name, val)
			//主要是符号长度和对对齐比较麻烦
			for i := 0; i < maxLength-len(strTheSameLen); i++ {
				str += "—"
			}
			str += "┘"
			fmt.Println(str)

			return
		}
	}

	if c.Head {
		var val string = c.Value
		if val != "" {
			val = ": " + val
		}
		var str = fmt.Sprintf("%s┌─ %s %s", indent, c.Icon+c.Name, val)
		lenStr := len(str)
		for i := 0; i < maxLength-lenStr; i++ {
			str += "—"
		}
		str += "┐"
		fmt.Println(str)

	} else {
		var val string = c.Value
		if val != "" {
			val = ": " + val
		}
		var str = fmt.Sprintf("%s├─ %s %s", indent, c.Icon+c.Name, val)
		lenStr := len(str)
		for i := 0; i < maxLength-lenStr; i++ {
			str += "—"
		}
		str += "┤"
		fmt.Println(str)
	}

	newIndent := indent + "|  "
	for idx, child := range c.Children {
		if idx == len(c.Children)-1 && c.Last {
			child.SetLast(true)
		} else {
			child.SetLast(false)
		}
		child.Draw(newIndent)
	}

}
