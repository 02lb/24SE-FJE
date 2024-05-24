package main

type JSONTree struct {
	Root AbstractNode
}

func NewJSONTree(factory JSONFactory, data map[string]interface{}) *JSONTree {
	root := factory.CreateNode("root", "")
	buildTree(root, data, factory)
	return &JSONTree{Root: root}
}

func buildTree(node AbstractNode, data map[string]interface{}, factory JSONFactory) {
	for key, value := range data {
		//TODO: 由于范例中对于value为string的情况需要进行输出，所以在buildTree时进行处理(已完成)
		a, ok := value.(string)
		var val string
		if ok {
			val = a
		} else {
			val = ""
		}
		child := factory.CreateNode(key, val)

		//TODO: 由于没有实现很好的封装，需要对style作出判断才能使用AddChild方法(已完成)
		_, ok = factory.(*TreeFactory)
		if ok {
			//指示当前工厂是TreeFactory，即style为tree
			node.(*TreeNode).AddChild(child.(*TreeNode))
		}
		_, ok = factory.(*RectangleFactory)
		if ok {
			//指示当前工厂是RectangleFactory，即style为rectangle
			node.(*RectangleNode).AddChild(child.(*RectangleNode))

		}

		if v, ok := value.(map[string]interface{}); ok {
			buildTree(child, v, factory)
		}
	}
}

func (t *JSONTree) Display() {
	t.Root.Display("")
}
