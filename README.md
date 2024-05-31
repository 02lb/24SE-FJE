## FJE-GO：一个JSON文件可视化的命令行界面小工具

#### DirTree:

<img src="tex-report/figures/dirtree.png" style="zoom: 33%;" />

#### **运行方式：**

```bash
$ fje-go -f <json file> -s <style> -i <icon family> -help <help>
```
**`-f`可以指定进行可视化转换的JSON源文件路径**

**`-s`可以快速切换风格(style)，包括：树形（tree）、矩形（rectangle）:**

**`-i`可以指定需要显示的图标族(IconFamily)**

**`-help`查看Options帮助**


```
├─ oranges                             ┌─ oranges ───────────────────────────────┐
│  └─ mandarin                         │  ├─ mandarin ───────────────────────────┤
│     ├─ clementine                    │  │  ├─ clementine ──────────────────────┤
│     └─ tangerine: cheap & juicy!     │  │  ├─ tangerine: cheap & juicy! ───────┤
└─ apples                              ├─ apples ────────────────────────────────┤
   └─ gala                             └──┴─gala ────────────────────────────────┘

        树形（tree）                                   矩形（rectangle）
```

**`-i`可以指定图标族（icon family），为中间节点或叶节点指定一套icon**：

```
├─♢oranges                                 
│  └─♢mandarin                             
│     ├─♤clementine                        
│     └─♤tangerine: cheap & juicy!    
└─♢apples                                  
   └─♤gala                                 

Using IconFamily: poker-face
中间节点: ♢     叶子节点: ♤  
```

**可以在[src/config/icon_config.json](src/config/icon_config.json)下配置图标族**

**可以在[src/json](src/json)下配置需要可视化的JSON文件**

#### 设计模式：

> 抽象工厂模式、工厂方法模式、组合模式、建造者模式、策略模式、模版模式……..

#### UML-类图：

![](tex-report/figures/uml-class.png)

#### 效果展示：

>简单JSON文件为 json/test.json

<img src="tex-report/figures/short-show.png" style="zoom:50%;" />

> 复杂JSON文件为 json/test_long.json

<img src="tex-report/figures/long-show.png" style="zoom: 67%;" />

