package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3}
	// [low:high:max] 取下表从low开始的元素 len=hight-low,cap=max-low
	s1 := arr[:] //[0:len(array):len(array)]不指定容量和长度一样
	fmt.Println("s1=",s1)
	fmt.Printf("len=%d,cap=%d\n",len(s1),cap(s1))

	// 操作某一个元素和数组操作方式一样
	data := arr[1]
	fmt.Println("data=",data)

	s2 := arr[3:6:7]//a[3] a[4] a[5] len= 6-3=3  cap=7-3=4
	fmt.Println("s2=",s2)
	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))

	s3 := arr[:6] //从0开始去6个元素容量也是6 常用
	fmt.Println("s3=",s3)
	fmt.Printf("len=%d,cap=%d\n",len(s3),cap(s3))

	s4 := arr[3:]//从下标3开始到结尾
	fmt.Println("s4=",s4)
	fmt.Printf("len=%d,cap=%d\n",len(s4),cap(s4))
}
