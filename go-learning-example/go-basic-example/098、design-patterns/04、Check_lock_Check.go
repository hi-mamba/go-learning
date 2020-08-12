package main

import "sync"

var mux sync.Mutex

var checkLockCheck_instance *checkLockCheck

// 构建一个结构体，用来实例化单例
type checkLockCheck struct {
	name string
}


func GetInstance() *checkLockCheck {
	if checkLockCheck_instance == nil { // 单例没被实例化，才会加锁
		mux.Lock()
		defer mux.Unlock()
		if checkLockCheck_instance == nil { // 单例没被实例化才会创建
			checkLockCheck_instance = &checkLockCheck{}
		}
	}
	return checkLockCheck_instance
}
