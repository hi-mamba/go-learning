package main

import "fmt"

/*
函数

函数可以没有参数或接受多个参数。

在这个例子中，`add` 接受两个 int 类型的参数。

注意类型在变量名 _之后_。

*/

func add(x int, y int) int {
	return x + y
}

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(add_version2(3, 5))
}

/*

函数（续）
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

在这个例子中 ，

x int, y int
被缩写为

x, y int
x, y int
*/

func add_version2(x, y int) int {
	return x + y
}
