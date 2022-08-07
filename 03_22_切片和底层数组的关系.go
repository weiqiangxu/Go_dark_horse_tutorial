package main

import "fmt"

func main() {
	a:= []int{0,1,2,3,4,5}
	// 新切片
	s1 := a[2:5]
	s1[1] = 666
	fmt.Println("s1=",s1)
	fmt.Println("a=",a)

	// 另外新切片
	s2 := s1[1:]
	s2[1] = 777
	fmt.Println("s2=",s2)
	fmt.Println("a=",a) //切片s1和s2的重新赋值都更改了原数组a的值

	// 数组比较和复制

	d := [5]int{1,2,3,4,5}
	e := [5]int{1,2,3,4,5}
	c := [5]int{1,2,3}
	fmt.Println("d==e",d==e)
	fmt.Println("e==c",e==c)

	// 同类型的数组可以赋值
	var g [5]int
	g = d
	fmt.Println("g=",g)

}
