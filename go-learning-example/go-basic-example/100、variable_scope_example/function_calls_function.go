package main

// 变量范围
var a3 string

func main() {
	a3 = "G"
	print(a3)
	f1()
}

func f1()  {
	a3 := "O"
	print(a3)
	f2()
}

func f2()  {
	print(a3)
}