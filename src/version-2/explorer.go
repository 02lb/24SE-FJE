package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// FunnyJSONExplorer 负责读取 JSON 文件并显示其内容
type FunnyJSONExplorer struct {
	factory Factory
}

func (fje *FunnyJSONExplorer) Show(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	builder := &Builder{factory: fje.factory}
	root := fje.factory.CreateContainer("root", "")
	builder.Build(root, jsonData)

	root.Draw("")
}

func (fje *FunnyJSONExplorer) Load(factory Factory) {
	fje.factory = factory
}
