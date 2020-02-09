

[原文](https://segmentfault.com/a/1190000012843377)

## hello world 程序的简单解释

下面是我们刚才写的程序代码
```go
package main //1

import "fmt" //2

func main() { //3
    fmt.Println("Hello World") //4
}
```
- package main - 每一个go程序必须以 package name 开头. 包的设计主要用来做代码隔离和代码可复用.
  这段程序里面的包名叫做 main

- import "fmt" - 导入fmt包用来在main函数中的输出文字到标准输出设备

- func main() - main函数是一个特殊的函数. 应用程序从main函数开始执行.
main 函数必须被放在main包中. The { and } indicate the start and end of the main function.

- fmt.Println("Hello World") - 使用fmt包中的Println用来输出文字到标准输出设备