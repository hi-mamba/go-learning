package main

import "net/http"

/*

【实例】HTTP 文件服务器是常见的 Web 服务之一。

 使用Go语言实现一个简单的 HTTP 服务器只需要几行代码，如下所示。
*/

func main() {

	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	http.Handle("/", http.FileServer(http.Dir(".")))

	//默认的 HTTP 服务侦听在本机 8080 端口。
	http.ListenAndServe(":8080", nil)
}
