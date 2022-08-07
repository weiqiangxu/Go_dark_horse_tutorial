package main

import "fmt"

func main() {
	// 声明
	var f1 float32
	f1 = 3.14
	fmt.Println("f1=", f1)
	// 自动推导类型
	f2 := 3.14
	fmt.Printf("f2 type is %T", f2)
	// float64 存储小数 比float32更准确
}
