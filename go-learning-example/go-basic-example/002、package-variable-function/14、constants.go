package main

import "fmt"

/*
常量
常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。
*/

const PI = 3.14

func main() {

	const World = "世界"
	fmt.Println("Hello", World)

	fmt.Println(PI)


	fmt.Println("####### ")
	print()
}

/*

数值常量

数值常量是高精度的 _值_。

一个未指定类型的常量由上下文来决定其类型。

也尝试一下输出 needInt(Big) 吧。



*/

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func print() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
