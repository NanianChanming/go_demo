package main

import "fmt"

func main() {
	a := appendString()
	b := appendString()
	fmt.Println(a("world"))
	fmt.Println(b("everyone"))

	fmt.Println(a("go"))
	fmt.Println(b("!"))
}

// 闭包
// 闭包是匿名函数的一个特例，当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包
func closureParam() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

// 每一个闭包都会绑定一个它自己的外围变量
func appendString() func(str string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
