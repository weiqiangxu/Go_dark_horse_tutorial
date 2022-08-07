// bool类型
package main

import "fmt"

func main() {
	// 声明变量没有初始化 零值是false
	var a bool
	a = true
	fmt.Println("a=", a)
	// 自动推导类型
	var b = false
	fmt.Println(b)
	c := false
	fmt.Println(c)
}
