// iota枚举
package main

import "fmt"

func main() {
	// 1、iota常量自动生成器，每一行自动累加1
	// 2、iota给常量赋值使用
	const (
		a = iota
		b = iota
	)
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
	// 3、iota遇到const 充值为0

	// 4、可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	// 5、如果是同一行，值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i=%d.j1=%d,j2=%d,j3=%d,k=%d", i, j1, j2, j3, k)
}
