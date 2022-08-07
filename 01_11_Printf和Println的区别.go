// printf和printLn的区别
package main

import "fmt"

func main() {
	a := 10
	// 一段一段输出，自动加换行
	fmt.Println("a=", a)
	// 格式化输出，把a的内容放在%d的位置
	// "a=10\n"这个字符串输出到屏幕 "\n"代表换行符
	fmt.Printf("a=%d\n", a)

	b := 20
	c := 30
	fmt.Println("a=", a, "b=", b, "c=", c)
	// 格式化的效果杠杠的
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
