## FJE-GO：一个JSON文件可视化的命令行界面小工具

**运行方式：**

```bash
$ fje-go -f <json file> -s <style> -i <icon family>
```

**`-i`指定进行可视化转换的JSON源文件路径**

**`-s`可以快速切换风格（style），包括：树形（tree）、矩形（rectangle）:**

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

poker-face-icon-family: 中间节点icon：♢ 叶节点icon：♤    
```

