package main
import "fmt"

func main(){
	var p *int
	p = nil //指针变量是可以直接赋值为null的
	fmt.Println("p = ",p)

	// *p = 666 // err 因为p没有合法指向
	var a int
	p = &a //p指向a
	*p = 666 //注意指针变量的赋值前面需要加一个*

	fmt.Printf("p =%v,&a=%v\n", *p, a)
}