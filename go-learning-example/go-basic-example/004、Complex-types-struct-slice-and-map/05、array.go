package main

import "fmt"

/*

数组

类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int
定义变量 a 是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
*/

func main() {

	var a [10]int
	a[1] = 1
	a[9] = 9
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("---------")
	var array [10][10]int
	array[1][1] = 1111

	for i := 0; i < 10; i++ {
		fmt.Println(array[i][i])
	}
}
