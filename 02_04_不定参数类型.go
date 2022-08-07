//不定参数类型
package main //必须
import "fmt"

// ...int类型这样的类型 ...type不定参数类型
func YouFunc(args ...int) { //传递的实参可以是0或者多个
	fmt.Println("len(args)=", len(args)) //获取用户传递参数的个数
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d\n", i, args[i])
	}
	//返回2个值 第一个是下标 第二个是下标所对应的数
	for i, data := range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}

// 固定参数一定要传参，不定参数按需求传递，不定参数一定是最后放置 也就是final
func heFunc(a int, args ...int) {

}

func main() {
	YouFunc(1, 3)
}
