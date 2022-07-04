package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go getX(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}

/*
当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。
这些修改共享资源的代码称为临界区。

Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
Mutex 可以在 sync 包内找到。Mutex 定义了两个方法：Lock 和 Unlock。
所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。
*/

var x = 0

func getX(wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()
}
