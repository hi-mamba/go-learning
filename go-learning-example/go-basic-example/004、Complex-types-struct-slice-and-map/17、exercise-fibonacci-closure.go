package main

import "fmt"

/*
练习：斐波纳契闭包

现在来通过函数做些有趣的事情。

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。

*/

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {

	count := 1
	fnLast := 1
	fnLastLast := 1
	return func() int {
		if count == 1 || count == 2 {
			count = count + 1
			return 1
		}

		value := fnLastLast + fnLast
		fnLastLast = fnLast
		fnLast = value
		return value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("fn(", i+1, ")", f())

	}
}
