package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	lock sync.Mutex
)

func main()  {

	for i := 0; i < 20; i++ {
		go incr();
	}
	time.Sleep(time.Millisecond* 10);
}
/*
互斥量序列化会锁住锁下的代码访问。因为默认的的 sync.Mutex 是未锁定状态，这儿我们就得先定义 lock sync.Mutex
*/
func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter ++
	fmt.Println(counter)
}
