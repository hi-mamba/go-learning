package main

import (
	"fmt"
	"testing"
)
// 当前目录执行 go test -v 01、sinleton_lazy_loading.go sigletion_lazy_loading_test.go
// https://www.cnblogs.com/Detector/p/10010292.html
func TestGetInstance(t *testing.T) {
	fmt.Println("### start test")
	//开启三个线程，说明线程池中只有三个线程， 在实际情况下可以动态设置开启线程数量
	for w := 1; w <= 3; w++ {
		go getInstance()
	}
}

