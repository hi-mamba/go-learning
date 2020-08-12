package main

import (
	"fmt"
	"sync"
)

type singleton2 struct {
	name string
}

var instance2 *singleton2
var mu sync.Mutex

func getInstance2() *singleton2 {
	mu.Lock()
	defer mu.Unlock()

	if instance2 == nil {
		fmt.Print("init..")
		instance2 = &singleton2{}     // unnecessary locking if instance already created
	}
	return instance2
}

func main() {
	s22 := getInstance2()
	s22.name = "test"
	fmt.Println(s22.name)
	s21 := getInstance2()
	fmt.Println(s21.name)
}
