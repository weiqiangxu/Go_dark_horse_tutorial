// 常量的使用
package main

import "fmt"

func main() {
	// 变量：程序运行期间可以改变的量
	// 常量：程序运行期间不可以改变的量。声明常量需要const
	const a int = 10
	// a = 20 //常量允许赋值修改的
	fmt.Println(a)
	const b = 10
	// const b := 10 const不能这样推导初始化
	fmt.Printf("%T\n", b)
	fmt.Println("b=", b)
}
