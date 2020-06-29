package main

import "fmt"

/*
例1.这是一个简单的回调例子，调用函数test时，调用真正的实现函数add

https://www.jianshu.com/p/0a9d25b6b0ea
*/
type Callback func(x, y int) int

// 提供一个接口，让外部去实现
func test(x, y int, callback Callback) int {
	return callback(x, y)
}

// 回调函数的具体实现
func add(x, y int) int {
	return x + y
}

func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

