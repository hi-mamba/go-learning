package main

import "fmt"

/*

range

for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

*/

func main() {

	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}

	for i, num := range nums {
		fmt.Printf("nums[%d]= %d\n", i, num)
	}

	fmt.Println("-----")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("-----")

	rangeContinued()
}

/*

range（续）

可以通过赋值给 _ 来忽略序号和值。

如果只需要索引值，去掉“, value”的部分即可。

*/

func rangeContinued() {

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
