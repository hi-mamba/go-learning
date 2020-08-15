package _08_test

import (
	"fmt"
	"testing"
)

//测试整个文件 go test -v hello_test.go
// 测试单个函数：$ go test -v hello_test.go -test.run TestHello

func hello(t *testing.T){
	fmt.Println("hello test")
}


func TestHello(t *testing.T) {
	fmt.Println("TestHello")
}

func TestWorld(t *testing.T) {
	fmt.Println("TestWorld")
}