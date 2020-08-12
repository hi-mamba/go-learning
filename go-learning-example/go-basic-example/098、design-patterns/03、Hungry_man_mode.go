package main

import "fmt"

// 构建一个结构体，用来实例化单例
type hungryExample struct {
	name string
}

// 声明一个私有变量，作为单例
var hungry_instance *hungryExample

// init函数将在包初始化时执行，实例化单例
func init() {
	hungry_instance = new(hungryExample)
	hungry_instance.name = "初始化单例模式"
}

func getHungryInstance() *hungryExample {
	return hungry_instance
}

func main() {
	s := getHungryInstance()
	fmt.Println(s.name)
}
