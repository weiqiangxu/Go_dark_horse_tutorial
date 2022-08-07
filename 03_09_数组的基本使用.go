package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int

	fmt.Printf("len(a)=%d,len(b)=%d\n",len(a),len(b))

	// 注意：定义数组是，指定的数组元素个数必须是常量
	// n := 10
	// var c [n]int // err non-constant array bound n

	// 操作数组元素，从0开始到len()-1，不对称元素，这个数字叫下标
	// 这是下标可以使变量或常量
	a[0] = 1
	i:=1
	a[i]=2 //a[1]=2

	// 赋值，每个元素
	for i:=0;i<len(a); i++ {
		a[i] = i+1
	}
	fmt.Print(a)
}
