package main

import "fmt"

/*func main() {
	// x := 3
	// fmt.Printf("%d 的阶乘是 %d", x, factorial(x))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t", fibonacci(i))
	}

	fmt.Println()

	fmt.Println(fibonacci(3))

	arrFibonacci(10)

}*/

// 递归函数
// Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
// 递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

// 数组实现斐波那契数列
func arrFibonacci(n int) {
	// 声明一个长度为n的数组
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			arr[i] = i
			fmt.Println(arr[i])
			continue
		}
		arr[i] = arr[i-1] + arr[i-2]
		fmt.Println(arr[i])
	}
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func factorial(x int) int {
	// 计算1到 x 的阶乘
	if x > 0 {
		return x * factorial(x-1)
	}
	return 1
}
