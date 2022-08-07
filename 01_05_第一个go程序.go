// 第1天06命令行运行程序

// 1.go语言以包作为管理单位
// 2.每个文件必须先声明包
// 3.程序必须有一个main包（重要）

// 这是定义的包名
package main

import "fmt"

// 入口主函数
func main() { //左括号必须与函数名同行
	// 打印 到屏幕 是通过1个包 println会自动换行
	//
	fmt.Println("hello go") //go语句结尾是为没有分号的
}

// 入口：go有且只有1个入口函数main
