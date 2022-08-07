package main

func main(){
	 type Person struct {
		 name string
		 sex byte
		 age int
	 }
	 type Student struct {
		 Person // 只有类型 没有名字 匿名字段 继承了Person的成员
		 id int
		 addr string
	 }
}