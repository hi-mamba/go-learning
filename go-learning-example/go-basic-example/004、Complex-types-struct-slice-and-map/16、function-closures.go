package main

import "fmt"

/*

函数的闭包

Go 语言闭包详解

https://www.sulinehk.com/post/golang-closure-details/

Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。
函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
*/

func adder() func(int) int {
	//因为闭包对外层词法域变量 pos 是引用的，所以这段代码 sum 不会重置 0 。
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println("序号：", i,
			pos(i),
			//neg(-2*i),
		)
	}
	/*

	pos := adder() 通过把这个函数变量赋值给 pos，pos 就成为了一个闭包。

	所以 pos 保存着对 sum 的引用，可以想象 sum 中有着一个指针指向 sum 或 sum 中有 sum 的地址。

	由于 sum 有着指向 sum 的指针，所以可以修改 sum，且保持着状态：

	pos(i), // 1
	pos(i), // 3
	pos(i), // 6
	也就是说，sum 逃逸了，它的生命周期没有随着它的作用域结束而结束。

	*/


	/*
	但是这段代码却不会递增：

	adder()(i), // 1
	adder()(i), // 2
	adder()(i), // 3
	这是因为这里调用了三次 adder()(i)，返回了三个闭包，这三个闭包引用着三个不同的 sum，它们的状态是各自独立的。
	 */

	for i := 0; i < 10; i++ {
		fmt.Println("序号：", i,
			adder()(i),
		)
	}
}
