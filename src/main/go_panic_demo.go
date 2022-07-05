package main

import "fmt"

func main() {
	// 加入defer语句
	defer fmt.Println("deferred call in main")
	firstName := "yun chang"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

/*
在go语言中，程序中一般是使用错误来处理异常情况，但在有些情况下，当程序发生异常时，无法继续运行,
这种情况下，我们会使用panic来终止程序，当函数发生panic时，它会终止运行，
在执行完所有的延迟函数后，程序控制返回到该函数的调用方，这样的过程会一直持续下去，
直到当前协程的所有函数返回退出，然后程序会打印出panic信息，接着打印出堆栈跟踪，最后程序终止

需要注意的是，你应该尽可能的使用错误，而不是使用panic和recover。
只有当程序不能继续运行的时候，才应该使用panic和recover机制

panic两个合理的用例
1.发生了一个不能恢复的错误，此时程序不能继续运行。
	一个例子就是web服务器无法绑定所要求的端口，在这种情况下，就应该使用panic，因为如果不能绑定端口，啥也做不了
2.发生了一个编程上的错误。
	假如我们有一个接收指针参数的方法，而其他人使用了nil作为参数调用了它。在这种情况下，可以使用panic，
	因为这是一个编程错误：用nil参数调用了一个只能接收合法指针的方法
*/
// 示例
//下列代码在发生panic时，首先执行了延迟函数，接着控制返回到函数调用方，
//调用方法的延迟函数继续执行，直到到达定成调用函数
func fullName(firstName *string, lastName *string) {
	// 修改当前代码，使用一个延迟语句
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s \n", firstName, lastName)
	fmt.Println("returned normally from fullName")
}
