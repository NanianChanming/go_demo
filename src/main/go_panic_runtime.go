package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	runtimePanic()
	fmt.Println("normally returned from main")
}

/*
运行时panic
运行时错误（如数组越界）也会导致panic，这等价于调用了内置函数panic，
其参数有接口类型runtime.Error给出。
而runtime.Error接口满足内建接口类型error
*/
// 示例
// 加入defer恢复运行时panic
func runtimePanic() {
	defer recoverRuntime()
	n := []int{1, 2, 3}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

//恢复后获得堆栈跟踪
//使用debug包中的PrintStack函数
func recoverRuntime() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
		debug.PrintStack()
	}
}
