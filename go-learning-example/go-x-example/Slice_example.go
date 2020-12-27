package main

import (
	"fmt"
	"unsafe"
)

/**
 原文：https://coolshell.cn/articles/21128.html
 */
type slice struct {
	array unsafe.Pointer //指向存放数据的数组指针
	len   int            //长度有多大
	cap   int            //容量有多大
}

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	fmt.Print(foo)
	fmt.Println()
	fmt.Print(bar)
	fmt.Println("foo和bar的内存是共享的，所以，foo和bar的对数组内容的修改都会影响到对方")
	fmt.Print(foo)
}
