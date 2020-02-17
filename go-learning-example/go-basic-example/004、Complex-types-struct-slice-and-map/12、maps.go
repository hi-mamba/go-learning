package main

import "fmt"

/*

https://go-tour-zh.appspot.com/moretypes/15


map

map 映射键到值。

map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。

*/

type VertexMap struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]VertexMap)
	m["Bell Labs"] = VertexMap{40.68433, -74.39967,}

	fmt.Println(m["Bell Labs"])


	fmt.Println(m["Bell Labs1"])


	m2 := make(map[string]string)
	m2["a"] = "a"
	m2["a1"] = "a"
	m2["a"] = "a1"
	m2["b"] = "b"
	fmt.Println(m2)
}

/*
map 的文法

map 的文法跟结构体文法相似，不过必须有键名。

m2 := make(map[string]string)
	m2["a"] = "a"

*/