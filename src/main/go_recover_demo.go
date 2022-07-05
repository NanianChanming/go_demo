package main

import (
	"fmt"
	"time"
)

func main() {
	// 假如defer语句
	/*defer fmt.Println("deferred call in main")
	firstName := "yun chang"
	fullNameRecover(&firstName, nil)
	fmt.Println("returned normally from main")*/

	// panic recover 和 go 协程
	testA()
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

//只有在相同的go协程中调用recover才管用，recover不能恢复一个不同协程的panic。
//在下面的程序中，函数B发生了panic， 函数A调用了一个延迟函数recovery（）用于恢复panic。
//但是函数B作为一个不同的协程来调用，下一行sleep只是保证A和B运行结束之后才退出
//这个时候，panic是不能被恢复的，因为调用recovery的协程和B方法中发生panic的协程并不相同
//如果把go testB() 修改为 testB() 就可以恢复panic了
func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered: ", r)
	}
}

func testA() {
	defer recovery()
	fmt.Println("inside A")
	//go testB()
	testB()
	time.Sleep(1 * time.Second)
}

func testB() {
	fmt.Println("inside B")
	panic("B method panicked")
}
