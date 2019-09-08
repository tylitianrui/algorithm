package main

import "sync"

type Obj struct {
}

var (
	lock sync.Mutex
	obj  *Obj

	// 实现2
	once sync.Once
	obj1 *Obj
)

// go单例1：double check
func Singleton() *Obj {
	if obj != nil {
		return obj
	}
	lock.Lock()
	defer lock.Unlock()
	if obj == nil {
		return &Obj{}
	} else {
		return obj
	}

}

// go单例2：使用once，仅仅执行一次
func Singleton1() *Obj {
	if obj1 == nil {
		once.Do(func() {
			if obj1 == nil {
				obj1 = &Obj{}
			}
		})
	}

	return obj1
}

//go单例3： 在init函数
