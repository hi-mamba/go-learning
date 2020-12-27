package main

import (
	"fmt"
	"strings"
)

/*
 原文：https://coolshell.cn/articles/21164.html

	Map示例
*/

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}
func main() {

	// 给第一个 MapStrToStr() 传了函数做的是 转大写，于是出来的数组就成了全大写的，
	// 给MapStrToInt() 传的是算其长度，所以出来的数组是每个字符串的长度。
	var list = []string{"Hao", "Chen", "MegaEase"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v\n", y)
}
