package main

import "fmt"

/*

初始化变量
变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。


*/

var i, j int = 1, 2

func main() {

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// Go语言多个变量同时赋值
	// Go 的“多重赋值”特性，可以轻松完成变量交换的任务,不在使用 中间变量进行变量的临时保存
	i, j = j, i
	fmt.Println(i, j)

}
