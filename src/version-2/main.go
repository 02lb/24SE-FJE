package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 定义命令行选项
	file := flag.String("f", "", "JSON file")
	style := flag.String("s", "", "Style (tree or rectangle)")
	iconReferred := flag.String("i", "default", "Icon family")
	help := flag.Bool("help", false, "Show this help message")

	// 自定义Usage函数
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// 解析命令行选项
	flag.Parse()

	// 如果--help被触发，则显示帮助信息并退出
	if *help {
		flag.Usage()
		return
	}

	// icon-dealing
	configFilePath := "../config/icon_config.json"
	iconConfig, err := LoadConfig(configFilePath)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	icon := iconConfig.IconFamilies[*iconReferred]

	// style-dealing
	var factory Factory
	switch *style {
	case "tree":
		factory = &TreeFactory{icon.Intermediate, icon.Leaf}
	case "rectangle":
		factory = &RectangleFactory{icon.Intermediate, icon.Leaf}
	//TODO: 在这里可以添加新的样式的抽象工厂
	default:
		fmt.Println("Unknown style")
		return
	}

	explorer := &FunnyJSONExplorer{}
	explorer.Load(factory)
	explorer.Show(*file)

	// 打印一行信息：即是否使用IconFamily 用的什么IconFamily:
	fmt.Println("\nUsing IconFamily:", *iconReferred)
	fmt.Println("中间节点:", icon.Intermediate, "	叶子节点:", icon.Leaf)
}
