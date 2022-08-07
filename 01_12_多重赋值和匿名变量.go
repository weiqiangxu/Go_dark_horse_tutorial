// 多重赋值和匿名变量
package main

import "fmt"

func main() {
	a, b := 10, 20
	// 交换2个变量的值
	var tmp int
	tmp = 1
	a = b
	b = tmp
	fmt.Printf("a=%d,b=%d,tmp=%d", a, b, tmp)
	// go语言特色
	i := 10
	j := 20
	i, j = j, i
	fmt.Printf("i=%d,j=%d\n", i, j)
	// 匿名变量
	tmp, _ = i, j
	fmt.Println("tmp=", tmp)
	// 为什么匿名变量不能直接tmp=i呢
	// go函数可以返回多个值
	var c, d, e int
	c, d, e = test() //return 1,2,3
	fmt.Printf("c=%d,d=%d,e=%d\n", c, d, e)
	d, _, _ = test()
	fmt.Printf("d=%d\n", d)
}
func test() (a, b, c int) {
	return 1, 2, 3
}
