// 无参数无返回值函数的使用
package main //必须
import "fmt"

func MyFunc() {
	a := 666
	fmt.Println("a=")
}

func main() {
	// 函数名() - 调用函数必须是定义的
	MyFunc()
}
