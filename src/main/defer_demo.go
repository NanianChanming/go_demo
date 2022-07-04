package main

import "fmt"

func main() {
	defer getName1()
	getName()
}

// defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。

func getName() {
	fmt.Println("name is guan yun chang")
}

func getName1() {
	defer finally()
	fmt.Println("name is zhang yi de")
}

// 函数调用中也可使用defer
func finally() {
	fmt.Println("we are come from three kingdom")
}
