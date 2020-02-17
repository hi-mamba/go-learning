package main

import (
	"fmt"
	"math"
)

/*

函数值

函数也是值

*/

func main() {

	hypot := func(x, y float64) float64 {
		return math.Pow(x, y)
	}

	fmt.Println(hypot(2, 2))
}
