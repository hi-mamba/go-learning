package main

import "fmt"

/*
https://coolshell.cn/articles/21164.html

通过这个三个示例，你可能有一些明白，Map/Reduce/Filter只是一种控制逻辑，
真正的业务逻辑是在传给他们的数据和那个函数来定义的。

是的，这是一个很经典的“业务逻辑”和“控制逻辑”分离解耦的编程模式。
下面我们来看一个有业务意义的代码，来让大家强化理解一下什么叫“控制逻辑”与业务逻辑分离。

*/
func Filter(arr[] int, fn func(n int) bool) []int  {

	var newArray []int
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray,it)
		}
	}
	return newArray
}

func main() {
	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool {
		return n%2 == 1
	})
	fmt.Printf("%v\n", out)
	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}
