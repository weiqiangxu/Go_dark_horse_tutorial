// 复数类型
package main

import "fmt"

func main() {
	var t complex128 //声明
	t = 2.1 + 3.14i  // 赋值 -- 实数 虚数
	fmt.Println("t=", t)
	// 自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Println("%T", t2)

	// 通过内建函数，取实部和虚部
	fmt.Println("real(t2)=", real(t2), ",imag(t2)=", imag(t2))
}
