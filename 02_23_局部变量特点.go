package main
import "fmt"

func test(){
	a := 10
	fmt.Println("a=",a)
}

func main(){
	// 变量定义在{}里面就是局部变量
	// 执行到定义变量那一句,才开始分配空间,离开作用域就自动释放
	// 作用域,变量其作用的范围

	// a=11
	{
		i:= 10
		fmt.Println("i=",i)
	}

	if flag:=3;flag==3{
		fmt.Println("flag=",flag)
	}
	// flag := 4
}