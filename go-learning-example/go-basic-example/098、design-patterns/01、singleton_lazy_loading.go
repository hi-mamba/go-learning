package main

import "fmt"

/*
 懒汉模式（Lazy Loading）
懒汉模式最大的缺点是非线程安全的
*/

// 构建一个示例结构体
type lazyLoading struct {
	name string
}
// private  设置一个私有变量作为每次要返回的单例
var instance *lazyLoading

func getInstance() *lazyLoading  {
	// 存在线程安全问题，高并发时有可能创建多个对象
	if instance == nil {
		instance = new(lazyLoading)
	}
	return instance
}
func main() {

	s := getInstance()
	s.name = "test"
	fmt.Println(s.name)
	s2 := getInstance()
	fmt.Println(s2.name)

}
