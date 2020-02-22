package main

import (
	"fmt"
	"runtime"
)

/*

switch

一个结构体（`struct`）就是一个字段的集合。

除非以 fallthrough 语句结束，否则分支会自动终止。

*/

func main() {

	fmt.Println(" GO RUN ON ")

	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)

	}

}
