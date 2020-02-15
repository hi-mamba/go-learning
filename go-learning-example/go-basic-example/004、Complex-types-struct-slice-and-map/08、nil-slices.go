package main

import "fmt"

/*

nil slice

slice 的零值是 `nil`。

一个 nil 的 slice 的长度和容量是 0。
*/

func main() {
	var slice []int
	fmt.Println(len(slice), cap(slice), slice)

	if slice == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println(" slice is not nil")
	}
}
