//有参无返回值的函数

package main //必须
import "fmt"

//有参数无返回值函数的定义，普通参数列表
func MyFunc(a int) {
	fmt.Println("a=", a)
}

func YouFunc(a, b int) {
	fmt.Println(a, b)
}

//主函数
func main() {
	//有参无返回值函数的调用: 函数名(所需函数)
	// 调用函数传递的参数叫实参
	MyFunc(777)
	YouFunc(6, 7)
}
