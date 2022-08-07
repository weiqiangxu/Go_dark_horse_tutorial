package main

import "fmt"

func main(){
	// break // break is not in loop switch 
	// continue // continue is not in a loop

	// goto 可以用在任何地方但是不能跨函数使用

	fmt.Println("1")

	goto End // goto是关键字 End是用户起的名字 叫标签

	fmt.Println("2")

	End:

	fmt.Println("3")


}