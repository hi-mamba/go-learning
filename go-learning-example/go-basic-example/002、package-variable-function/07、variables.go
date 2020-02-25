package main

import (
	"fmt"
	"net"
)

/*

变量

var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。

就像在这个例子中看到的一样，`var` 语句可以定义在包或函数级别。

*/

var c, python, golang,java bool

func main() {
	fmt.Println(c, java, python, golang)



	//在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
	// 但是 err 是 第二个的
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	conn2, err := net.Dial("tcp", "127.0.0.1:8081")
	conn2, err2 := net.Dial("tcp", "127.0.0.1:8080")

	fmt.Println(conn,"err:",err)
	fmt.Println(conn2,"err:",err)
	fmt.Println(conn2,"err2:",err2)
}
