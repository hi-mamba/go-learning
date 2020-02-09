package main

import "fmt"

/*
for

Go 只有一种循环结构——`for` 循环。

基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。
*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + 1
	}

	fmt.Println("sum=", sum)
}

/*

for（续）

跟 C 或者 Java 中一样，可以让前置、后置语句为空。

*/
func forExample() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
