package main
import "fmt"

func test(x int){
	result := 100/x
	fmt.Println("result = ",result)
}

func main(){
	// 多个defer先进先出
	defer fmt.Println("a")
	defer fmt.Println("b")
	//调用一个函数导致内存出问题
	defer test(0)
	defer fmt.Print("c")
}