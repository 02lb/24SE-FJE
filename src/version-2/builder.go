package main

type Builder struct {
	factory Factory
}

// Build方法 根据 JSON 数据构建树结构，返回根节点
func (b *Builder) Build(root Component, data map[string]interface{}) {
	for key, value := range data {
		a, ok := value.(string)
		var val string
		if ok {
			val = a
		} else {
			val = ""
		}
		var child Component
		if _, ok := value.(map[string]interface{}); ok {
			child = b.factory.CreateContainer(key, val)
			root.(*Container).AddChild(child)
		} else {
			// Leaf and NoneTy
			child := b.factory.CreateLeaf(key, val)
			root.(*Container).AddChild(child)
		}

		if v, ok := value.(map[string]interface{}); ok {
			b.Build(child, v)
		}
	}
}
