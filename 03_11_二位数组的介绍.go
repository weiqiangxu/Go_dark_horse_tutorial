package main

import "fmt"

func main(){
	var a = [][]int{}
	k:=0
	for i := 0; i < 3; i++ {
		for j:=0;j<4;j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]-%d",i,j,a[i][j])
		}
		fmt.Printf("\n")
	}

	// 有3个元素每个元素有事以为数组
	b := [3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println("b=",b)

	e := [3][4]int{1:{5,6}}
	fmt.Println("e=",e)
}