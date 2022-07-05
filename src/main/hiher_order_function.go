package main

import "fmt"

func main() {
	// 函数作为参数传递
	//f := func(a, b int) int {
	//	return a + b
	//}
	//simple(f)

	// 函数作为返回值
	f := returnFunc()
	fmt.Println(f(50, 5))
}

// 高阶函数
/*
wiki把高阶函数定义为满足下列条件之一的函数：
* 接收一个或多个函数作为参数
* 返回值是一个函数
*/
// 把函数作为参数
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 5))
}

// 在其他函数中返回函数
func returnFunc() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
