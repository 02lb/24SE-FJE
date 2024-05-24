package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := flag.String("f", "", "Path to the JSON file")
	style := flag.String("s", "", "Style to be applied (tree, rectangle)")
	// iconFamily := flag.String("i", "default", "Icon family to be used")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a JSON file path with -f")
		return
	}

	fileContent, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	//output the json file content for testing
	//for key, value := range data {
	//	// fmt.Println(key, value)
	//	a := value.(map[string]interface{})[key]
	//	content, ok := a.(string)
	//	if ok {
	//		fmt.Println(content)
	//	}
	//}

	var factory JSONFactory
	fmt.Println("Creating a JSON tree with style:", *style)
	switch *style {
	case "tree":
		factory = &TreeFactory{}
	case "rectangle":
		factory = &RectangleFactory{}
	default:
		fmt.Println("Unknown style. Use 'tree' or 'rectangle'.")
		return
	}

	tree := NewJSONTree(factory, data)
	tree.Display()
}
