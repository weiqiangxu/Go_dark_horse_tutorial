// 变量的输入

package main

import "fmt"

func main() {
	var a int //声明变量
	fmt.Printf("请输入变量a：")
	fmt.Scan(&a)
	// fmt.Scanf("%d", &a)
	fmt.Println("a=", a)

}
