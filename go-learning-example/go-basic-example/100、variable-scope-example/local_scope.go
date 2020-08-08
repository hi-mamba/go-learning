package main

import "fmt"

// 局部变量和全局变量
var a  = "G"

func main() {
	n()
	m()
	n()
}

func n()  {
	fmt.Print( a)
}

func m()  {
	//
	a:="O"
	fmt.Print(a)
}

