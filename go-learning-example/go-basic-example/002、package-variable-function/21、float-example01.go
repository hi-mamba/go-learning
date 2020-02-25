package main

import (
	"fmt"
	"math"
)

/*

 */

func main() {

	// 一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，
	//而 float64 则可以提供约 15 个十进制数的精度，通常应该优先使用 float64 类型，
	//因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大。
	var a float32 = 16777216 // 1 << 24  超过了6个精度
	fmt.Println(a == a+1)    // "true"!

	//浮点数在声明的时候可以只写整数部分或者小数部分，像下面这样：
	var b float64 = .1
	var c float64 = 1.

	fmt.Println(b, c)


	// 很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分：
	const Avogadro = 6.02214129e23   // 阿伏伽德罗常数
	const Planck   = 6.62606957e-34 // 普朗克常数
	fmt.Println(Avogadro)
	fmt.Println(Planck)




	// 用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数

	fmt.Printf("%f\n",math.Pi)
	fmt.Printf("%.2f\n",math.Pi)

}
