package main
/*
package main - 每一个go程序必须以 package name 开头.
包的设计主要用来做代码隔离和代码可复用.
这段程序里面的包名叫做 main
*/
import (
	"../math_example"
	"fmt"
)
// import "fmt" - 导入fmt包用来在main函数中的输出文字到标准输出设备


// this is a comment

/*
func main() - main函数是一个特殊的函数.
应用程序从main函数开始执行.
main 函数必须被放在main包中.
The { and } indicate the start and end of the main function.
*/
func main() {
	// fmt.Println("Hello World") - 使用fmt包中的Println用来输出文字到标准输出设备
	fmt.Println("Hello World")
	fmt.Println(math_example.Add(1,1))
	fmt.Println(math_example.Sub(1,1))

}