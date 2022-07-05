package main

import "fmt"

func main() {
	//funcToVariable()
	//funcRun()
	//funcAndParam()
	funcFuncType()
}

/*
头等函数（第一类函数）
支持头等函数的编程语言，可以把函数赋值给变量，也可以把函数作为其他函数的参数或者返回值
go语言支持头等函数机制
*/
//把函数赋值给一个变量
func funcToVariable() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
}

//调用匿名函数，可以不用赋值给变量
func funcRun() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

//对匿名函数传参
func funcAndParam() {
	func(s string) {
		fmt.Println("welcome", s)
	}("guan yun chang")
}

//用户自定义函数类型
type add func(a, b int) int

func funcFuncType() {
	var a add = func(a, b int) int {
		return a + b
	}
	i := a(5, 6)
	fmt.Println("sum =", i)
}
