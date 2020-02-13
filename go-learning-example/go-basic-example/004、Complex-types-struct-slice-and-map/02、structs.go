package main

import "fmt"

/*

结构体
一个结构体（`struct`）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。）

*/

type Vertex struct {
	Y int
	X int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	/*
		结构体字段
		结构体字段使用点号来访问。
	*/
	fmt.Println(v.X)
}


