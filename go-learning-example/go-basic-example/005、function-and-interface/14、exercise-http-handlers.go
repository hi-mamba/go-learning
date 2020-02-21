package main

/*
练习：HTTP 处理

实现下面的类型，并在其上定义 ServeHTTP 方法。在 web 服务器中注册它们来处理指定的路径。

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
例如，可以使用如下方式注册处理方法：

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
*/


import (
	"log"
	"net/http"
)

func main() {
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}


