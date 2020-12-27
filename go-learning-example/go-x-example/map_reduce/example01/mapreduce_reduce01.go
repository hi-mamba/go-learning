package main

import "fmt"

/*
 原文：https://coolshell.cn/articles/21164.html

	 Reduce 示例
*/


func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _,it := range arr {
		sum += fn(it)
	}
	return sum
}
func main() {

	var list = []string{"kobe","lakers"}

	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
}
