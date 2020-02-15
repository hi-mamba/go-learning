package main

import "fmt"

/*

range

for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

*/

func main() {
	var pow =  []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
