package main

import "fmt"

// 匿名函数

/*func main() {
	anon1()
}*/

func anon2() {
	// ()代表直接调用此匿名函数
	func() {
		fmt.Println("---匿名函数执行---")
	}()
}

func anon1() {
	f1 := func() {
		fmt.Println("---匿名函数执行---")
	}
	// 接收一个函数类型返回值然后进行调用
	f1()
}
