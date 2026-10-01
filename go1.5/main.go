package main

import "fmt"

// Go 1.5（2015-08-19）语言特性演示入口。
//
// 发布说明中“Changes to the language”只有 map 复合字面量键类型省略这一项；
// 具体说明和运行方法见 README.md。
func main() {
	fmt.Println("Go 1.5：map 复合字面量键可省略类型")
	demoMapKeys()
}
