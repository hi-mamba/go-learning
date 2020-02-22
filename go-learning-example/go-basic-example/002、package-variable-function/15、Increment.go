package main

import "fmt"

/*
 Increment 自增

不再纠结于 i++ 和 ++i
*/

func main() {

	i := 0
	i++

	fmt.Println("i=", i)

	j := 0
	// 编译错误
	++j
	fmt.Println("j=", j)
}
