package main

import "fmt"

func main() {
	// 假如defer语句
	defer fmt.Println("deferred call in main")
	firstName := "yun chang"
	fullNameRecover(&firstName, nil)
	fmt.Println("returned normally from main")
}

/*
recover 是一个内建函数，用于重新获得panic协程的控制
只有在延迟函数的内部，调用recover才有用。
在延迟函数内调用recover,可以取到panic的错误信息，并且停止panic续发事件（Panicking Sequence），
程序运行回复正常，如果在延迟函数的外部调用recover，就不能停止panic续发事件
*/
// 在发生panic之后，使用recover来恢复正常运行
func fullNameRecover(firstName *string, lastName *string) {
	// 修改当前代码，使用一个延迟语句
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s \n", firstName, lastName)
	fmt.Println("returned normally from fullName")
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
